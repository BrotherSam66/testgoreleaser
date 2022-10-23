# GoReleaser 工作笔记

http://www.idmiss.com/701

1.   ##  GoReleaser 介绍

        goreleaser  方便的go 二进制包分发工具。可以帮助我们进行go 二进制包的快速，简单分发，我们可以用来创建一个github release 以及发布到homwbrew formula 。

1.   ## 安装 
        go install github.com/goreleaser/goreleaser@latest
        sudo snap install --classic goreleaser
2.   ## 编辑
     1. ### 添加git 支持同时添加tag
        git init
        git add --all
        git commit -m  "init project"
        git remote add origin https://github.com/rongfengliang/gorelease-app.git
        git tag v1.0.0
        git push -u origin master
        git push -u origin v1.0.0
     2. ### 添加goreleaser 支持 创建 .goreleaser.yaml
        goreleaser init
     3. ###  测试 .goreleaser.yaml
        goreleaser check
     4. ###  本地测试构建
        goreleaser --snapshot --skip-publish --rm-dist
     5. ###  准备发布，设置 ENV github token 
        export GITHUB_TOKEN="？？？？YOUR_GH_TOKEN"
     7. ###  创建 tag 推送 github
         git tag -a v0.1.0 -m "First release"
         git push origin v0.1.0      
3.   ## 编译  测试构建
         goreleaser --snapshot --skip-publish --rm-dist     
         goreleaser release --snapshot --rm-dist
4.   ## 发布到 github
      export GITHUB_TOKEN="YOUR_GH_TOKEN"
      or
      export GITLAB_TOKEN="YOUR_GL_TOKEN"
      git tag -a v0.0.1 -m "打包v0.1.2"
      git push origin v0.1.2
      goreleaser release --rm-dist  
5.   ## 发布到 gitlab  

     5. ### 效果
         ```
         before:
         hooks:
            # You may remove this if you don't use go modules.
            - go mod tidy
            # you may remove this if you don't need go generate
            - go generate ./...
         builds:
         - env:
               - CGO_ENABLED=0
            goos:
               - linux
               - windows
               - darwin
         archives:
         - replacements:
               darwin: Darwin
               linux: Linux
               windows: Windows
               386: i386
               amd64: x86_64
         checksum:
         name_template: 'checksums.txt'
         snapshot:
         name_template: "{{ incpatch .Version }}-next"
         changelog:
         sort: asc
         filters:
            exclude:
               - '^docs:'
               - '^test:'

         ```
      


安装
mac 系统，实际根据自己的系统选择，我使用brew 安装

brew install goreleaser
使用
创建go mod
go mod init github.com/dalongrong/gorelease-app
添加简单代码
main.go
package main

import (
 "fmt"
)

func main() {
 fmt.Println("demoapp")
}
添加git 支持同时添加tag
git init
git add --all
git commit -m  "init project"
git remote add origin https://github.com/rongfengliang/gorelease-app.git
git tag v1.0.0
git push -u origin master
git push -u origin v1.0.0
添加goreleaser 支持
goreleaser init
测试构建
goreleaser --snapshot --skip-publish --rm-dist
效果

 goreleaser release --skip-publish --rm-dist

   • releasing using goreleaser 0.101.0...
   • loading config file file=.goreleaser.yml
   • RUNNING BEFORE HOOKS
      • running go mod download
      • running go generate ./...
   • GETTING AND VALIDATING GIT STATE
      • releasing v1.0.1, commit 101af58d04813fabb32567c4b4aab7925783a6fc
   • PARSING TAG      
   • SETTING DEFAULTS 
      • LOADING ENVIRONMENT VARIABLES
      • SNAPSHOTING      
      • GITHUB RELEASES  
      • PROJECT NAME     
      • ARCHIVES         
      • BUILDING BINARIES
      • LINUX PACKAGES WITH NFPM
      • SNAPCRAFT PACKAGES
      • CALCULATING CHECKSUMS
      • SIGNING ARTIFACTS
      • DOCKER IMAGES    
      • ARTIFACTORY      
      • S3               
      • HOMEBREW TAP FORMULA
      • SCOOP MANIFEST   
   • SNAPSHOTING      
      • pipe skipped error=not a snapshot
   • CHECKING ./DIST  
      • --rm-dist is set, cleaning it up
   • WRITING EFFECTIVE CONFIG FILE
      • writing config=dist/config.yaml
   • GENERATING CHANGELOG
      • writing changelog=dist/CHANGELOG.md
   • LOADING ENVIRONMENT VARIABLES
      • pipe skipped error=publishing is disabled
   • BUILDING BINARIES
      • building binary=dist/windows_386/gorelease-app.exe
      • building binary=dist/linux_386/gorelease-app
      • building binary=dist/linux_amd64/gorelease-app
      • building binary=dist/darwin_amd64/gorelease-app
      • building binary=dist/darwin_386/gorelease-app
      • building binary=dist/windows_amd64/gorelease-app.exe
   • ARCHIVES         
      • creating archive=dist/gorelease-app_1.0.1_Windows_x86_64.tar.gz
      • creating archive=dist/gorelease-app_1.0.1_Linux_i386.tar.gz
      • creating archive=dist/gorelease-app_1.0.1_Darwin_i386.tar.gz
      • creating archive=dist/gorelease-app_1.0.1_Darwin_x86_64.tar.gz
      • creating archive=dist/gorelease-app_1.0.1_Linux_x86_64.tar.gz
      • creating archive=dist/gorelease-app_1.0.1_Windows_i386.tar.gz
   • LINUX PACKAGES WITH NFPM
      • pipe skipped error=no output formats configured
   • SNAPCRAFT PACKAGES
      • pipe skipped error=no summary nor description were provided
   • CALCULATING CHECKSUMS
      • checksumming file=gorelease-app_1.0.1_Darwin_x86_64.tar.gz
      • checksumming file=gorelease-app_1.0.1_Windows_i386.tar.gz
      • checksumming file=gorelease-app_1.0.1_Darwin_i386.tar.gz
      • checksumming file=gorelease-app_1.0.1_Linux_x86_64.tar.gz
      • checksumming file=gorelease-app_1.0.1_Linux_i386.tar.gz
      • checksumming file=gorelease-app_1.0.1_Windows_x86_64.tar.gz
   • SIGNING ARTIFACTS
      • pipe skipped error=artifact signing is disabled
   • DOCKER IMAGES    
      • pipe skipped error=docker section is not configured
   • PUBLISHING       
      • pipe skipped error=publishing is disabled
   • release succeeded after 5.73s 
