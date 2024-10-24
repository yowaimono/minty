package internal

import (
	"log"


	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/logger"
	// "github.com/gofiber/fiber/v2/middleware/recover"
)

func StartServer(projects []Project,app *fiber.App) {

	// 创建 Fiber 应用


	// 添加中间件
	// app.Use(logger.New())
	// app.Use(recover.New())

	// 为每个项目创建路由组
	for _, project := range projects {
		// 获取绝对路径
		//absRoot, err := filepath.Abs(project.Root)
		// if err != nil {
		// 	log.Fatalf("Failed to get absolute path for %s: %v", project.Root, err)
		// }

		// // 检查目录是否存在
		// if _, err := os.Stat(project.Root); os.IsNotExist(err) {
		// 	log.Fatalf("Directory %s for project %s does not exist", project.Root, project.Prefix)
		// }

		// 配置静态文件服务
		app.Static(project.Prefix, project.Root)
		log.Printf("Serving project %s from %s", project.Prefix, project.Root)

	}
	// 启动服务器
	log.Printf("Starting server on http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
