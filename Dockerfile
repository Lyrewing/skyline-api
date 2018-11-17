#源镜像
FROM golang:1.10-alpine3.8
#作者
MAINTAINER FengZhanyuan "iven.feng@outlook.com"
#将服务器的go工程代码加入到docker容器中
RUN ls
RUN  mkdir /skyline-api
COPY skyline /skyline-api/
#设置环境变量
ENV MYSQL_DB_CONNECTION root:123456@tcp(skyline-db:3306)/skyline?charset=utf8&parseTime=True
ENV STATSD_URL skyline-telegraf:8125
#设置工作目录
WORKDIR /skyline-api
RUN [ -f /skyline-api/skyline ] && echo "file fount" || echo "file not found"
#暴露端口
EXPOSE 8080
#最终运行docker的命令
ENTRYPOINT  ["/skyline-api/skyline"]
