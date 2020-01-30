# 图片验证码(captcha)


提供图片验证码生成与验证服务。

Docker service 
```
docker service create --name captcha --network cluster --replicas 1  -d hub.unionlive.com/captcha:latest
docker service update --force --update-parallelism 1 --update-delay 3s --image hub.unionlive.com/captcha:latest captcha
docker service update  --replicas 3  captcha

```

# Change Log

## v1.0.0 
    [Release Date 2019-12-08]
- [feature] 创建项目
