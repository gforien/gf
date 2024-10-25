package aws

import (
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func EqualsIpPerms(a []types.IpPermission, b []types.IpPermission) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !EqualsString(a[i].IpProtocol, b[i].IpProtocol) {
			return false
		}
		if !EqualsIpRange(a[i].IpRanges, b[i].IpRanges) {
			log.Default().Print("IpRanges not equal")
			return false
		}
		if !EqualsIpv6Range(a[i].Ipv6Ranges, b[i].Ipv6Ranges) {
			log.Default().Print("Ipv6Ranges not equal")
			return false
		}
	}
	return true
}

func EqualsIpRange(a []types.IpRange, b []types.IpRange) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !EqualsString(a[i].CidrIp, b[i].CidrIp) {
			return false
		}
	}
	return true
}

func EqualsIpv6Range(a []types.Ipv6Range, b []types.Ipv6Range) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !EqualsString(a[i].CidrIpv6, b[i].CidrIpv6) {
			return false
		}
	}
	return true
}

func EqualsString(a *string, b *string) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	return strings.EqualFold(*a, *b)
}
