package aws

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func TestEqualsIpPerms(t *testing.T) {
	tests := []struct {
		name     string
		a        []types.IpPermission
		b        []types.IpPermission
		expected bool
	}{
		{
			name:     "Nil slices",
			a:        nil,
			b:        nil,
			expected: true,
		},
		{
			name:     "One nil and one empty slice",
			a:        nil,
			b:        []types.IpPermission{},
			expected: true,
		},
		{
			name:     "One nil slice",
			a:        nil,
			b:        []types.IpPermission{{IpProtocol: stringPointer("tcp")}},
			expected: false,
		},
		{
			name:     "Different lengths",
			a:        []types.IpPermission{{IpProtocol: stringPointer("tcp")}},
			b:        []types.IpPermission{},
			expected: false,
		},
		{
			name:     "Equal slices",
			a:        []types.IpPermission{{IpProtocol: stringPointer("tcp")}},
			b:        []types.IpPermission{{IpProtocol: stringPointer("tcp")}},
			expected: true,
		},
		{
			name:     "Different protocols",
			a:        []types.IpPermission{{IpProtocol: stringPointer("tcp")}},
			b:        []types.IpPermission{{IpProtocol: stringPointer("udp")}},
			expected: false,
		},
		{
			name: "Different IP ranges",
			a: []types.IpPermission{
				{IpProtocol: stringPointer("tcp"), IpRanges: []types.IpRange{{CidrIp: stringPointer("192.168.1.0/24")}}},
			},
			b: []types.IpPermission{
				{IpProtocol: stringPointer("tcp"), IpRanges: []types.IpRange{{CidrIp: stringPointer("10.0.0.0/24")}}},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EqualsIpPerms(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("EqualsIpPerms() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestEqualsIpRange(t *testing.T) {
	tests := []struct {
		name     string
		a        []types.IpRange
		b        []types.IpRange
		expected bool
	}{
		{
			name:     "Nil slices",
			a:        nil,
			b:        nil,
			expected: true,
		},
		{
			name:     "One nil and one empty slice",
			a:        nil,
			b:        []types.IpRange{},
			expected: true,
		},
		{
			name:     "One nil slice",
			a:        nil,
			b:        []types.IpRange{{CidrIp: stringPointer("192.168.1.0/24")}},
			expected: false,
		},
		{
			name:     "Different lengths",
			a:        []types.IpRange{{CidrIp: stringPointer("192.168.1.0/24")}},
			b:        []types.IpRange{},
			expected: false,
		},
		{
			name:     "Equal slices",
			a:        []types.IpRange{{CidrIp: stringPointer("192.168.1.0/24")}},
			b:        []types.IpRange{{CidrIp: stringPointer("192.168.1.0/24")}},
			expected: true,
		},
		{
			name:     "Different CIDR IPs",
			a:        []types.IpRange{{CidrIp: stringPointer("192.168.1.0/24")}},
			b:        []types.IpRange{{CidrIp: stringPointer("10.0.0.0/24")}},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EqualsIpRange(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("EqualsIpRange() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestEqualsIpv6Range(t *testing.T) {
	tests := []struct {
		name     string
		a        []types.Ipv6Range
		b        []types.Ipv6Range
		expected bool
	}{
		{
			name:     "Nil slices",
			a:        nil,
			b:        nil,
			expected: true,
		},
		{
			name:     "One nil and one empty slice",
			a:        nil,
			b:        []types.Ipv6Range{},
			expected: true,
		},
		{
			name:     "One nil slice",
			a:        nil,
			b:        []types.Ipv6Range{{CidrIpv6: stringPointer("2001:db8::/32")}},
			expected: false,
		},
		{
			name:     "Different lengths",
			a:        []types.Ipv6Range{{CidrIpv6: stringPointer("2001:db8::/32")}},
			b:        []types.Ipv6Range{},
			expected: false,
		},
		{
			name:     "Equal slices",
			a:        []types.Ipv6Range{{CidrIpv6: stringPointer("2001:db8::/32")}},
			b:        []types.Ipv6Range{{CidrIpv6: stringPointer("2001:db8::/32")}},
			expected: true,
		},
		{
			name:     "Different CIDR IPv6s",
			a:        []types.Ipv6Range{{CidrIpv6: stringPointer("2001:db8::/32")}},
			b:        []types.Ipv6Range{{CidrIpv6: stringPointer("2001:0db8::/32")}},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EqualsIpv6Range(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("EqualsIpv6Range() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestEqualsString(t *testing.T) {
	tests := []struct {
		name     string
		a        *string
		b        *string
		expected bool
	}{
		{
			name:     "Nil strings",
			a:        nil,
			b:        nil,
			expected: true,
		},
		{
			name:     "One nil and one empty string",
			a:        nil,
			b:        stringPointer(""),
			expected: true,
		},
		{
			name:     "One nil string",
			a:        nil,
			b:        stringPointer("test"),
			expected: false,
		},
		{
			name:     "Different strings",
			a:        stringPointer("abc"),
			b:        stringPointer("ABC"),
			expected: true,
		},
		{
			name:     "Same strings",
			a:        stringPointer("abc"),
			b:        stringPointer("def"),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EqualsString(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("EqualsString() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// Helper function to create string pointers for testing
func stringPointer(s string) *string {
	return &s
}
