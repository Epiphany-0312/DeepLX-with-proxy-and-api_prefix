
# DeepLX with Proxy and API Prefix

本项目为开源项目，在 [DeepLX](https://github.com/OwO-Network/DeepLX) 的基础上进行二次开发。本项目为 DeepLX 增加了代理功能和 API 前缀功能，方便用户在使用过程中通过代理服务器访问以及通过自定义的前缀路径进行 API 调用。

## 功能特性

- **代理功能**：支持通过 SOCKS5 代理服务器访问 DeepL API，增强访问的隐私性和安全性。
- **API 前缀功能**：允许用户自定义 API 路径前缀，有效防止服务被网络测绘工具扫描到，提高服务的安全性。

## 部署方法

推荐使用 Docker Compose 方法进行部署。以下是具体的部署步骤：

1. 克隆项目代码

   ```bash
   git clone https://github.com/Epiphany-0312/DeepLX-with-proxy-and-api_prefix.git
   ```

2. 进入项目文件夹
 ```
cd DeepLX-with-proxy-and-api_prefix
   ```

3. 修改 config.yaml 文件，根据需要设置相关配置参数。

   ```yaml
   port: 1188
   token: ""
   auth_key: ""
   dl_session: ""
   #前三个变量含义见deeplx原项目
   api_prefix: "/your_api_prefix"
   #api前缀，设置后需访问http://your_ip:1188/your_api_prefix/translate
   use_proxy: false
   #是否开启代理
   proxy_address: ""
   #代理地址和端口

   ```

4. 使用 Docker Compose 启动服务

   ```bash
   docker-compose up --build
   ```




## 项目结构

- **main.go**：主程序入口，定义了路由和中间件。
- **translate.go**：包含与 DeepL API 的交互逻辑。
- **types.go**：定义了各种数据类型。
- **utils.go**：包含一些工具函数。
- **config.go**：负责加载和解析配置文件。
- **Dockerfile**：用于构建 Docker 镜像。
- **docker-compose.yml**：用于定义 Docker Compose 服务。
- **config.yaml**：配置文件，包含所有配置参数。

## 贡献

欢迎提交问题 (issues) 和拉取请求 (pull requests) 来贡献代码。

## 许可

本项目基于 MIT License 许可进行开源。

如果您在使用过程中遇到任何问题，欢迎在 GitHub 项目页面提交问题 (issues)，我们将尽快回复并解决。
