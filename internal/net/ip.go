package net

import (
	"fmt"
	"io"
	"net/http"
)

// Get host public IPv4 and IPv6 from ipify.org
func GetPublicIPs() (string, string, error) {
	ipv4Resp, err := http.Get("https://api.ipify.org")
	if err != nil {
		return "", "", fmt.Errorf("failed to get IPv4: %w", err)
	}
	defer ipv4Resp.Body.Close()

	ipv4, err := io.ReadAll(ipv4Resp.Body)
	if err != nil {
		return "", "", fmt.Errorf("failed to read IPv4: %w", err)
	}

	ipv6Resp, err := http.Get("https://api6.ipify.org")
	if err != nil {
		return string(ipv4), "", nil
	}
	defer ipv6Resp.Body.Close()

	ipv6, err := io.ReadAll(ipv6Resp.Body)
	if err != nil {
		return string(ipv4), "", nil
	}

	return string(ipv4), string(ipv6), nil
}
