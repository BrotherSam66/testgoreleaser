# GoReleaser 使用

##  GoReleaser 介绍
   goreleaser 是一个方便的 go 二进制包分发工具。可以帮助我们进行 go 二进制包的快速，简单分发。
   
   我们可以用来创建一个项目，打包成不同格式，一键发布到 github、gitlib、gitea 。

   我们可以把项目一键打包成 docker 镜像，发布到 `docker.io`、`gcr.io` 。

## 安装 
   `go install github.com/goreleaser/goreleaser@latest`

   或

   `sudo snap install --classic goreleaser`

## 创建 go 项目
   略

## 项目发布
### 添加goreleaser 支持 创建 .goreleaser.yaml
   goreleaser init

### 检测 .goreleaser.yaml
   goreleaser check

### 本地测试构建
   goreleaser --snapshot --skip-publish --rm-dist

### 准备发布，设置 ENV github token 
   export GITHUB_TOKEN="YOUR_GITHUB_TOKEN"
### 创建 tag
   git tag -a v0.1.0 -m "First release"

### Docker 相关
####  创建 Makefile
   ```
      FROM scratch
      COPY YOUR_BINARY_NAME /app
      ENTRYPOINT ["/app"]  
   ```
#### 文件 .goreleaser.yaml 加上
   ```
   dockers:
     - image_templates:
       - "docker.io/YOUR_DOCKER_IO_NAME/YOUR_IMG_NAME:{{ .Version }}-amd64"
       # - "docker.io/YOUR_DOCKER_IO_NAME/YOUR_IMG_NAME:{{ .Version }}-arm64v8" 
   ```
#### 可能用到的 DOCKER 指令
   ```
   显示: docker images
   清空: docker system prune -a
   执行: docker run YOUR_IMG_NAME:VERSION
   推送：docker push YOUR_IMG_NAME:VERSION
   登录：docker login --username=YOUR_NAME
   ```  
## 发布到 `github` && `docker.io`
   goreleaser release --rm-dist
