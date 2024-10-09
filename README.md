# CoolFunAdmin

- 开发中遇到任何痛点写到[todo](todo.md)里
- 路由接口：参考[Google 开放平台](https://github.com/googleapis/googleapis)

## 依赖

- [Hertz](https://github.com/cloudwego/hertz) 作为 HTTP 框架
- `Protobuf` IDL 定义 `HTTP` 接口
- `hz` 代码生成
- `Ent` 与 `MySQL`
- `内存缓存` 与 `Redis`（Config 文件配置启用）

## 接口定义

本项目使用`Protobuf` IDL 定义`HTTP` 接口。对应的 admin 模块相关接口在 [admin.proto](api/admin/admin.proto) 文件中定义。

## 代码生成

本项目使用 `hz` 生成代码. `hz` 详细使用说明参考 [hz](https://www.cloudwego.io/docs/hertz/tutorials/toolkit/toolkit/).

- hz install.

```bash
go install github.com/cloudwego/hertz/cmd/hz@latest
```

- hz new: 新建一个 Hertz 项目

```bash
hz new -I api -idl api/admin/admin.proto -model_dir api/model -module coolfun-admin --unset_omitempty
```

- hz update: 当你的 IDL 文件更新，使用该指令进行项目代码更新
- api.proto 与 base.proto 是不需要更新与生成的，因为它们是由导入它们的 proto 文件生成的

```bash
hz update -I api -idl api/admin/admin.proto -model_dir api/model --unset_omitempty
```

## 变量绑定与校验

变量绑定与校验使用说明参考 [Binding and Validate](https://www.cloudwego.io/docs/hertz/tutorials/basic-feature/binding-and-validate/).

#### 快速开始

- 在项目根目录下依次执行以下命令, 将在目录 data/ent/schema/ 生成一个 User 的实体:

```bash
  cd data
  go run -mod=mod entgo.io/ent/cmd/ent init User
```

- 在 User schema 中添加表字段, 运行以下命令将生成 Ent 的操作代码文件

```bash
  go generate ./ent
```

- Ent 的更多使用说明，请参考 [Ent Guides](https://entgo.io/).

### 更新配置文件

- 使用你自己的参考更新 [Prod configuration file](configs/config.yaml) 与 [Dev configuration file](configs/config_dev.yaml)。
- 注意，YAML 文件的参数结构要与 config.go 里的 Config 结构体一致。
- 当运行时环境变量 "IS_PROD" 等于 true, 将会使用生产环境 Prod 配置, 否则使用测试环境 Dev 配置。

#### 初始管理员账号: admin/123456
