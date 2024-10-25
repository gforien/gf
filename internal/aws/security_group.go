package aws

import (
	"context"
	"log"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/gforien/gf/internal/net"
)

func FindAndUpdateSg(cfg aws.Config, ips []net.Ip) {
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

	var wg sync.WaitGroup
	for _, sg := range sgResult.SecurityGroups {
		wg.Add(1)
		go func() {
			defer wg.Done()
			AuthorizeInboundIps(ec2Client, sg, ips)
		}()
	}

	wg.Wait()
}

func AuthorizeInboundIps(ec2Client *ec2.Client, sg types.SecurityGroup, ips []net.Ip) {
	log.Default().Printf("Checking security group '%s'", *sg.GroupId)
	perms := []types.IpPermission{}
	for _, ip := range ips {

		if ip.ExistsInAwsSg(sg) {
			log.Default().Printf("Security group '%s' allows '%s'. Skipping.\n", *sg.GroupId, ip)
			continue
		}

		perms = append(perms, ip.ToAwsIpPerms())
		log.Default().Printf("Adding %v to group", ip)
	}
	if len(perms) == 0 {
		return
	}

	// Cleanup group rules
	if len(sg.IpPermissions) > 0 {
		revokeInput := &ec2.RevokeSecurityGroupIngressInput{
			GroupId:       sg.GroupId,
			IpPermissions: sg.IpPermissions,
		}
		_, err := ec2Client.RevokeSecurityGroupIngress(context.TODO(), revokeInput)
		if err != nil {
			log.Default().Panic("failed to revoke security group ingress: " + err.Error())
		}
	}

	autorizeInput := &ec2.AuthorizeSecurityGroupIngressInput{
		GroupId:       sg.GroupId,
		IpPermissions: perms,
	}
	log.Default().Printf("Updating SG '%s' with group.\n", *sg.GroupId)

	_, err := ec2Client.AuthorizeSecurityGroupIngress(context.TODO(), autorizeInput)
	if err != nil {
		log.Default().Panic("failed to update security group ingress: " + err.Error())
	}

	log.Default().Printf("Security group '%s' updated.\n", *sg.GroupId)
}
