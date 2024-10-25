package net

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type IpList []Ip

func (ipl IpList) ToAwsIpPerms() []types.IpPermission {
	var ipv4 []types.IpRange
	var ipv6 []types.Ipv6Range

	for _, ip := range ipl {
		cidr := ip.GetCidr()
		switch ip.GetVersion() {
		case IPv6:
			ipv6 = append(ipv6, types.Ipv6Range{CidrIpv6: aws.String(cidr)})
		default:
			ipv4 = append(ipv4, types.IpRange{CidrIp: aws.String(cidr)})
		}
	}

	return []types.IpPermission{
		{
			IpProtocol: aws.String("-1"),
			IpRanges:   ipv4,
			Ipv6Ranges: ipv6,
		},
	}
}
