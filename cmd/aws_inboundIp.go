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
	ips, err := net.GetPublicIps([]string{
		"https://api.ipify.org",
		"https://api6.ipify.org",
	})
	if err != nil {
		log.Default().Panic("unable to get public IP: " + err.Error())
	}

	aws.FindAndUpdateSg(ips)
}

func init() {
	awsCmd.AddCommand(inboundIpCmd)
}
