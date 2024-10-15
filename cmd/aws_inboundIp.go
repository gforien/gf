package cmd

import (
	"log"

	"github.com/gforien/gf/internal/aws"
	"github.com/gforien/gf/internal/net"
	"github.com/spf13/cobra"
)

// inboundIpCmd represents the inboundMyip command
var inboundIpCmd = &cobra.Command{
	Use:   "inboundIp",
	Short: "Update AWS sg 'inbound-myip' with current ipv4/ipv6",
	Run:   inboundIp,
}

func inboundIp(cmd *cobra.Command, args []string) {
	// Retrieve public IPv4 and IPv6
	ipv4, ipv6, err := net.GetPublicIPs()
	if err != nil {
		log.Default().Panic("unable to get public IP: " + err.Error())
	}
	log.Default().Printf("Got public IPv4: %s\n", ipv4)
	log.Default().Printf("Got public IPv6: %s\n", ipv6)

	aws.UpdateSg(ipv4, ipv6)
}

func init() {
	awsCmd.AddCommand(inboundIpCmd)
}
