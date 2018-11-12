#源镜像
FROM golang:1.10
#作者
MAINTAINER FengZhanyuan "iven.feng@outlook.com"
#将服务器的go工程代码加入到docker容器中
ADD ./skyline $/skyline-app
#设置工作目录
WORKDIR $/skyline-app
#暴露端口
EXPOSE 8080
#最终运行docker的命令
ENTRYPOINT  ["./skyline"]
