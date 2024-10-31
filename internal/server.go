package internal

import (
	"fmt"
	"log"
	log1 "minty/pkg/logger"
	"net"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func StartServer(config *Config) {
	// 创建 Fiber 应用
	app := fiber.New()

	// 禁用浏览器缓存
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Set("Pragma", "no-cache")
		c.Set("Expires", "0")
		return c.Next()
	})

	// 添加中间件
	app.Use(logger.New())
	app.Use(recover.New())

	// 静态文件服务
	for name, project := range config.Static {
		absRoot, err := filepath.Abs(project.Path)
		if err != nil {
			log.Fatalf("Failed to get absolute path for %s: %v", project.Path, err)
		}
		if _, err := os.Stat(absRoot); os.IsNotExist(err) {
			log.Fatalf("Directory %s for URI %s does not exist", absRoot, project.URI)
		}
		app.Static(project.URI, absRoot)
		log.Printf("Serving static %s files for URI %s from %s", name, project.URI, absRoot)
	}

	// 请求转发中间件
	app.Use(func(c *fiber.Ctx) error {
		// 解析请求头中的 Host 字段，去掉端口号
		host, _, err := net.SplitHostPort(c.Hostname())
		if err != nil {
			host = c.Hostname()
		}

		for name, rule := range config.Rules {
			// 解析规则，例如 "example.com -> URL_ADDRESS"
			log1.Info("name,rule -> %s, %s", name, rule)
			parts := strings.Split(rule, " -> ")
			if len(parts) == 2 {
				ruleHost := strings.TrimSpace(parts[0])
				target := strings.TrimSpace(parts[1])

				log1.Info("host,target -> %s, %s", ruleHost, target)
				if host == ruleHost {
					log1.Info("Forwarding to %s", target)
					return proxy.Do(c, target+c.OriginalURL())
				}
			}
		}
		log1.Info("c.Hostname() -> %s", c.Hostname())
		return c.Next()
	})

	// 添加一个简单的路由来测试
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// 启动服务器
	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
	log.Printf("Starting server on %s", addr)
	log.Fatal(app.Listen(addr))
}
