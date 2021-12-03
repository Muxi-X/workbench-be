# Muxi Workbench


Backend mono repo for  Muxi Workbench.

Built with [go-micro](https://github.com/micro/go-micro), [gin](https://github.com/gin-gonic/gin), [gorm](https://github.com/jinzhu/gorm), [zap](https://github.com/uber-go/zap), [jager](https://github.com/jaegertracing/jaeger).

### Services

+ [Gateway](https://github.com/Muxi-X/workbench-be/tree/master/service/gateway)
+ [Status](https://github.com/Muxi-X/workbench-be/tree/master/service/status)
+ [User](https://github.com/Muxi-X/workbench-be/tree/master/service/user)
+ [Team](https://github.com/Muxi-X/workbench-be/tree/master/service/team)
+ [Project](https://github.com/Muxi-X/workbench-be/tree/master/service/project)
+ [Feed](https://github.com/Muxi-X/workbench-be/tree/master/service/feed)


### Trigger Build

```shell
// add muxi origin
git remote add muxi http://gitea.muxixyz.com/root/workbench_be.git

// commit code(git commit ... git add ...) and
// sync code to muxi repo
git push muxi master

// create tag
git tag release-${service_name}-${version}
git push --tags muxi
// open ci.muxixyz.com to check build progress
```

### deploy subscribe

```shell
go run main.go -sub # 增加命令行参数 -sub 来运行
# 实际部署的时候需要改 Dockerfile （暂定）
```