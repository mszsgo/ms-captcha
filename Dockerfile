FROM alpine:3.10
ENV TZ=Asia/Shanghai
COPY ./ /
CMD ["/ms"]
