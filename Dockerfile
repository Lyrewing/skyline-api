#源镜像
FROM golang:1.10
#作者
MAINTAINER FengZhanyuan "iven.feng@outlook.com"
#将服务器的go工程代码加入到docker容器中
ADD ./skyline $/skyline-api
#设置环境变量
ENV MYSQL_DB_CONNECTION root:1qaz2wsx@tcp(192.168.31.214:13306)/skyline_db?charset=utf8&parseTime=True
ENV STATSD_URL 192.168.31.214:18125
#设置工作目录
WORKDIR ./skyline-api
#暴露端口
EXPOSE 8080
#最终运行docker的命令
ENTRYPOINT  ["./skyline"]
