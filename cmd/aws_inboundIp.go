package cmd

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/gforien/gf/internal/aws"
	"github.com/gforien/gf/internal/net"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

	var profiles []string
	if err = viper.UnmarshalKey("aws.inbound_ip.profiles", &profiles); err != nil {
		log.Default().Panic("unable to load aws profiles: " + err.Error())
	}
	if len(profiles) == 0 {
		log.Default().Panic("no aws profile provided, exiting")
	}

	for _, p := range profiles {
		log.Default().Printf("Loading profile '%s'", p)

		cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(p))
		if err != nil {
			log.Default().Panicf("unable to load AWS SDK config with profile '%s': %s", p, err.Error())
		}
		aws.FindAndUpdateSg(cfg, ips)
	}
}

func init() {
	awsCmd.AddCommand(inboundIpCmd)
}
