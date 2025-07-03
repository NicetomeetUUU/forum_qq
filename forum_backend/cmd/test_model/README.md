# 模型层接口测试样例

这个目录包含了一个简单的模型层接口测试示例，展示了如何测试go-zero自动生成的CRUD接口。

## 测试的接口

基于 `model/user/userModel_gen.go` 中自动生成的接口：

```go
type userModel interface {
    Insert(ctx context.Context, data *User) (sql.Result, error)
    FindOne(ctx context.Context, id int64) (*User, error)
    FindOneByEmail(ctx context.Context, email string) (*User, error)
    FindOneByUsername(ctx context.Context, username string) (*User, error)
    Update(ctx context.Context, data *User) error
    Delete(ctx context.Context, id int64) error
}
```

## 运行测试

```bash
# 确保数据库已启动
cd database && docker-compose up -d

# 运行测试
cd cmd/test_func
go run main.go
```

## 测试内容

1. **Insert** - 插入用户数据
2. **FindOne** - 根据ID查询用户
3. **FindOneByEmail** - 根据邮箱查询用户
4. **FindOneByUsername** - 根据用户名查询用户
5. **Update** - 更新用户信息
6. **Delete** - 删除用户

## 扩展测试

当你为其他模型（如Post、Comment）添加自定义接口时，可以参考这个模式：

```go
// 1. 创建模型实例
postModel := post.NewPostModel(conn, cache.CacheConf{})

// 2. 测试插入
result, err := postModel.Insert(ctx, testPost)

// 3. 测试查询
foundPost, err := postModel.FindOne(ctx, id)

// 4. 测试更新
err = postModel.Update(ctx, updatePost)

// 5. 测试删除
err = postModel.Delete(ctx, id)
```

## 注意事项

- 测试前确保数据库连接正常
- 测试数据会在数据库中创建和删除
- 可以根据需要修改测试数据
- 这个样例只测试自动生成的基础CRUD接口 