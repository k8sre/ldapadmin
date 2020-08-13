FROM hub-mirror.c.163.com/library/golang:1.13.10 AS builder
WORKDIR /go/src/k8sre/ldapadmin
COPY . ./
RUN export GOPROXY=https://mirrors.aliyun.com/goproxy/ && \
    make build


FROM registry-vpc.cn-zhangjiakou.aliyuncs.com/k8sre/alpine:3.11-glibc

RUN apk update                             && \
    apk upgrade

WORKDIR /opt/ldapadmin
COPY --from=builder /go/src/k8sre/ldapadmin/output/ /opt/ldapadmin/
CMD ["/opt/ldapadmin/bin/ldapadmin"]
