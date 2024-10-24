Minty 是一个用 Golang 实现的 HTTP 代理服务器，旨在实现 Nginx 的部分功能，包括根据请求的 Host 进行转发和静态文件代理。它的开发初衷是为了提供一个易于使用、配置简单或无需配置即可部署前端项目的解决方案，同时解决单主机多项目请求转发的问题。

### 静态文件代理

目前，Minty 支持静态文件代理，可用于部署 Vue 等前端项目。使用以下命令即可快速部署：

```sh
minty -p "uri_prefix:/path/to/dist"
```

您还可以在同一主机上部署多个 Vue 项目，只需为每个项目设置不同的 URI 前缀。Minty 会将打包后的 `dist` 静态文件代理到相应的 URI 前缀下。

### 请求转发

请求转发功能目前仍在开发中。

### 配置文件

Minty 支持无配置文件启动，通过语义化的命令即可启动服务。当然，您也可以选择使用配置文件进行更详细的设置。以下是一个配置文件示例：

```toml
[static]
app1 = "/project:/home/admin/dist/"
app2 = "/project2:/home/admin/dist2/"

[rules]
minapp1 = "www.app1.com -> :8080"
minapp2 = "www.app2.com -> :8090"
```

### 结语

欢迎使用 Minty！我们期待您的反馈和建议，以帮助我们进一步改进和增强 Minty 的功能。