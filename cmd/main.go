package main

import (
	"flag"
	"minty/internal"
	"strings"

	"github.com/gofiber/fiber/v2"
)


func parseProjects(projectsStr string) []internal.Project {
	var projects []internal.Project

	if projectsStr == "" {
		return projects
	}

	projectPairs := strings.Split(projectsStr, ",")
	for _, pair := range projectPairs {
		parts := strings.Split(pair, ":")
		if len(parts) == 2 {
			projects = append(projects, internal.Project{
				Prefix: parts[0],
				Root:   parts[1],
			})
		}
	}

	return projects
}

func main() {
	// 解析命令行参数
	projectsStr := flag.String("p", "", "Projects in the format 'prefix1:root1,prefix2:root2,...'")
	flag.Parse()

	// 解析项目配置
	projects := parseProjects(*projectsStr)
	app := fiber.New()

	// 启动服务器
	internal.StartServer(projects,app)
}
