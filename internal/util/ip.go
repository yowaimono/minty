package util

import (
	"fmt"
	"net"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// GetClientIP 从请求头中获取客户端IP地址
func GetClientIP(c *fiber.Ctx) (string, error) {
	ip := c.Get("X-Forwarded-For")
	ip = strings.TrimSpace(strings.Split(ip, ",")[0])
	if ip == "" {
		ip = strings.TrimSpace(c.Get("X-Real-IP"))
	}
	if ip == "" {
		var err error
		ip, _, err = net.SplitHostPort(strings.TrimSpace(c.IP()))
		if err != nil {
			return "", fmt.Errorf("failed to parse RemoteAddr: %v", err)
		}
	}

	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return "", fmt.Errorf("invalid IP address: %s", ip)
	}

	return ip, nil
}

// GetRequestDomain 从请求头中获取请求的域名
func GetRequestDomain(c *fiber.Ctx) (string, error) {
	host := c.Hostname()
	if host == "" {
		return "", fmt.Errorf("Host header is missing")
	}

	// 如果Host头包含端口号，去掉端口号
	domain, _, err := net.SplitHostPort(host)
	if err != nil {
		// 如果没有端口号，直接使用Host头的值
		domain = host
	}

	return domain, nil
}