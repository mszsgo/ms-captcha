FROM alpine:3.10
COPY ./captcha /captcha
CMD ["/captcha"]
