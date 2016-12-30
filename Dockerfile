FROM golang:1.7.4

MAINTAINER Bin Du <northtyphoon@gmail.com>

# Install beego & bee
RUN go get github.com/astaxie/beego
RUN go get github.com/beego/bee