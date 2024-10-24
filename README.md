Minty is an HTTP proxy server implemented in Golang, designed to replicate some functionalities of Nginx, including request forwarding based on the Host header and static file proxying. Its development aims to provide a solution that is easy to use, requires minimal or no configuration to deploy front-end projects, and addresses the issue of request forwarding for multiple projects on a single host.

### Static File Proxying

Currently, Minty supports static file proxying, which can be used to deploy projects like Vue. You can quickly deploy a project using the following command:

```sh
minty -p "uri_prefix:/path/to/dist"
```

You can also deploy multiple Vue projects on the same host by setting different URI prefixes for each project. Minty will proxy the `dist` static files to the corresponding URI prefixes.

### Request Forwarding

The request forwarding feature is currently under development.

### Configuration File

Minty supports starting without a configuration file, using semantic commands to launch the service. Of course, you can also choose to use a configuration file for more detailed settings. Here is an example of a configuration file:

```toml
[static]
app1 = "/project:/home/admin/dist/"
app2 = "/project2:/home/admin/dist2/"

[rules]
minapp1 = "www.app1.com -> :8080"
minapp2 = "www.app2.com -> :8090"
```

### Conclusion

Welcome to use Minty! We look forward to your feedback and suggestions to help us further improve and enhance the features of Minty.