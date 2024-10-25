package net

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
)

type (
	Ip        string
	IpVersion string
)

const (
	IPv4 = "IPv4"
	IPv6 = "IPv6"
)

func FromString(ip string) (Ip, error) {
	parsed := net.ParseIP(ip)
	if parsed == nil {
		return "", fmt.Errorf("failed to parse IP address '%s'", ip)
	}
	return Ip(ip), nil
}

func (ip Ip) GetVersion() IpVersion {
	if strings.Contains(string(ip), ":") {
		return IPv6
	}
	return IPv4
}

func (ip Ip) GetCidr() string {
	if ip.GetVersion() == IPv6 {
		return fmt.Sprintf("%s/128", ip)
	}
	return fmt.Sprintf("%s/32", ip)
}

// Get host public IP from an external API
//
// Exemple: GetPublicIp("https://api.ipify.org/")
func GetPublicIp(url string) (Ip, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to request '%s': %w", url, err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read body of '%s': %w", url, err)
	}

	ip, err := FromString(string(body))
	if err != nil {
		return "", fmt.Errorf("failed to parse IP address '%s': %w", string(body), err)
	}

	log.Default().Printf("Got public %s: '%s'\n", ip.GetVersion(), ip)

	return ip, nil
}

// Get host public IPs from an array of external APIs
func GetPublicIps(urls []string) ([]Ip, error) {
	var ips []Ip

	for _, url := range urls {
		ip, err := GetPublicIp(url)
		if err != nil {
			log.Default().Printf("Failed to request '%s', skipping", url)
			continue
		}

		ips = append(ips, ip)
	}

	return ips, nil
}
