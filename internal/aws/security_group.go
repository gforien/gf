package aws

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func UpdateSg(ipv4 string, ipv6 string) {
	// Load AWS config with the specified profile "gforien-prod"
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithSharedConfigProfile("gforien-prod"),
	)
	if err != nil {
		log.Default().Panic("unable to load AWS SDK config: " + err.Error())
	}

	ec2Client := ec2.NewFromConfig(cfg)

	// Filter security group by tag "Name" == "inbound-myip"
	describeSGInput := &ec2.DescribeSecurityGroupsInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("tag:Name"),
				Values: []string{"inbound-myip"},
			},
		},
	}

	sgResult, err := ec2Client.DescribeSecurityGroups(context.TODO(), describeSGInput)
	if err != nil {
		log.Default().Panic("failed to describe security groups: " + err.Error())
	}
	log.Default().Printf("Got %d security groups", len(sgResult.SecurityGroups))

	for _, sg := range sgResult.SecurityGroups {
		allowIp(ec2Client, sg, ipv4, ipv6)
	}
}

func allowIp(ec2Client *ec2.Client, sg types.SecurityGroup, ipv4 string, ipv6 string) {
	if ruleExists(sg, ipv4, ipv6) {
		log.Default().Printf("Security group '%s' allows current IPv4/v6. Skipping.\n", *sg.GroupId)
		return
	}

	// Revoke all ingress rules
	revokeInput := &ec2.RevokeSecurityGroupIngressInput{
		GroupId:       sg.GroupId,
		IpPermissions: sg.IpPermissions,
	}
	log.Default().Printf("Security group '%s' emptied.\n", *sg.GroupId)

	_, err := ec2Client.RevokeSecurityGroupIngress(context.TODO(), revokeInput)
	if err != nil {
		log.Default().Panic("failed to update security group ingress rules: " + err.Error())
	}

	// Update the security group ingress rules to allow IPv4 and IPv6
	authorizeIngressInput := &ec2.AuthorizeSecurityGroupIngressInput{
		GroupId: sg.GroupId,
		IpPermissions: []types.IpPermission{
			{
				FromPort:   aws.Int32(0),
				ToPort:     aws.Int32(65535),
				IpProtocol: aws.String("-1"),
				IpRanges:   []types.IpRange{{CidrIp: aws.String(ipv4 + "/32")}},
				Ipv6Ranges: []types.Ipv6Range{{CidrIpv6: aws.String(ipv6 + "/128")}},
			},
		},
	}

	_, err = ec2Client.AuthorizeSecurityGroupIngress(context.TODO(), authorizeIngressInput)
	if err != nil {
		log.Default().Panic("failed to update security group ingress rules: " + err.Error())
	}

	log.Default().Printf("Security group '%s' updated.\n", *sg.GroupId)
}

func ruleExists(sg types.SecurityGroup, ipv4, ipv6 string) bool {
	for _, permission := range sg.IpPermissions {
		for _, ipRange := range permission.IpRanges {
			if *ipRange.CidrIp == ipv4+"/32" {
				return true
			}
		}

		for _, ipv6Range := range permission.Ipv6Ranges {
			if *ipv6Range.CidrIpv6 == ipv6+"/128" {
				return true
			}
		}
	}
	return false
}
