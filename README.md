# hathor
哈索尔（Hathor），亦称哈托尔，全称哈索尔·迪特拉（Hathor Ditera），古埃及女神，她是爱神、美神、富裕之神、舞蹈之神、音乐之神。

# 初始化
```shell
go mod init github.com/pinkhello/hathor
```

# 安装 entgo
```shell
go install entgo.io/ent/cmd/ent@latest 
```
[https://entgo.io/](https://entgo.io/)

```shell
# 初始化一个model
# go run entgo.io/ent/cmd/ent init [--target ent/schema] User
go run entgo.io/ent/cmd/ent init User
# 构建 生成 model的目录
mkdir -p gen/entmodels && touch gen/entmodels/package_name.go && echo "package entmodels" > gen/entmodels/package_name.go 
# 生成 models
ent generate --target gen/entmodels ./ent/schema
# 清理
rm gen/entmodels/package_name.go
```
