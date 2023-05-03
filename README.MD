# GoFrame Template For SingleRepo

Project Makefile Commands: 
- `make cli`: Install or Update to the latest GoFrame CLI tool.
- `make dao`: Generate go files for `Entity/DAO/DO` according to the configuration file from `hack` folder.
- `make service`: Parse `logic` folder to generate interface go files into `service` folder.
- `make image TAG=xxx`: Run `docker build` to build image according `manifest/docker`.
- `make image.push TAG=xxx`: Run `docker build` and `docker push` to build and push image according `manifest/docker`.
- `make deploy TAG=xxx`: Run `kustomize build` to build and deploy deployment to kubernetes server group according `manifest/deploy`.


# FAQ
## 1. 初始化项目失败

错误信息
```
go run main.go
# github.com/gogf/gf/contrib/drivers/sqlite/v2
../../../pkg/mod/github.com/gogf/gf/contrib/drivers/sqlite/v2@v2.3.3/sqlite.go:81:12: config.Extra undefined (type *gdb.ConfigNode has no field or method Extra)
../../../pkg/mod/github.com/gogf/gf/contrib/drivers/sqlite/v2@v2.3.3/sqlite.go:86:40: config.Extra undefined (type *gdb.ConfigNode has no field or method Extra)
../../../pkg/mod/github.com/gogf/gf/contrib/drivers/sqlite/v2@v2.3.3/sqlite.go:150:22: undefined: gutil.GetOrDefaultStr
```

原因：gdb 版本太老，无法识别 Extra 字段。您可以尝试升级到最新版本的 gdb 包，或者尝试使用其他合适的配置来连接 SQLite 数据库，以解决此问题。

解决方案，在 go.sum 中将 github.com/gogf/gf/contrib/drivers/sqlite/v2 升级到 v2.4.0，重新 `go mod tidy`


## 2. 生成 dao 失败

错误信息
```
gf gen dao -r user           
error format "internal/dao/.go" go files: internal/dao/.go:22:3: expected 'IDENT', found '=' (and 1 more errors)
generated: internal/dao/.go
generated: internal/dao/internal/.go
error format "internal/model/do/.go" go files: internal/model/do/.go:13:7: expected 'IDENT', found 'struct' (and 1 more errors)
generated: internal/model/do/.go
error format "internal/model/entity/.go" go files: internal/model/entity/.go:12:7: expected 'IDENT', found 'struct' (and 1 more errors)
generated: internal/model/entity/.go
```

原因，命令错误，下面是正确命令，-r 是 removePrefix 的意思，把 user remove，表名就为空了

```
gf gen dao
```

## 3. go test 单元测试失败

将 _test.go 放在 controller 下，进入目录执行 go test 显示数据库驱动 sqlite 找不到。单元测试中并没有数据库连接，怀疑编译了 controller 文件。