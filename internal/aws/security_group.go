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
				Values: []string{"*-myip"},
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

func AuthorizeInboundIps(ec2Client *ec2.Client, sg types.SecurityGroup, ips net.IpList) {
	log.Default().Printf("Checking security group '%s'", *sg.GroupId)

	if len(sg.IpPermissions) < 1 {
		log.Default().Printf("Security group '%s' is empty. Skipping.", *sg.GroupId)
		return
	}

	port := sg.IpPermissions[0].FromPort
	protocol := sg.IpPermissions[0].IpProtocol
	if port == nil || protocol == nil {
		log.Default().Printf("Security group '%s' has nil port or protocol. Skipping.", *sg.GroupId)
		return
	}
	perms := ips.ToAwsIpPerms(port, protocol)
	if EqualsIpPerms(perms, sg.IpPermissions) {
		log.Default().Printf("Security group '%s' allow group. Skipping.", *sg.GroupId)
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
