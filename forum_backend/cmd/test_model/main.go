package main

import (
	"context"
	"fmt"
	"forum_backend/model/user"
	"log"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func main() {
	// 连接数据库
	ctx := context.Background()
	conn := sqlx.NewMysql("root:root542@tcp(localhost:3307)/qq_forum?parseTime=true")
	cacheConf := cache.CacheConf{
		{
			Weight: 100,
			RedisConf: redis.RedisConf{
				Host: "localhost:6379",
				Type: "node",
				Pass: "redisqiuqiu542",
			},
		},
	}
	userModel := user.NewUserModel(conn, cacheConf)
	// userModel.PrintPrivateInfo()
	// 测试1: 查询用户
	fmt.Println("\n1. 测试用户查询")
	users, err := userModel.FindAllActiveUsers(ctx)
	if err != nil {
		fmt.Printf("❌ 查询用户失败: %v\n", err)
	} else {
		fmt.Printf("✅ 查询用户成功，用户数量: %d\n", len(users))
	}
	// 测试2: 插入用户
	fmt.Println("\n1. 测试用户插入")
	testUser := &user.User{
		Email:     "test@example.com",
		Password:  "password123",
		Username:  "testuser",
		Role:      "user",
		Status:    "active",
		IsDeleted: 0,
	}
	var id int64 = 0
	user, _ := userModel.FindOneByEmail(ctx, testUser.Email)
	if user != nil {
		fmt.Println("用户已存在")
		id = user.Id
	} else {
		result, err := userModel.Insert(ctx, testUser)
		id, _ = result.LastInsertId()
		if err != nil {
			fmt.Printf("❌ 插入用户失败: %v\n", err)
		} else {
			fmt.Printf("✅ 用户插入成功，ID: %d\n", id)
		}
	}

	// 测试2： 删除用户
	fmt.Println("\n2. 测试用户删除")
	err = userModel.Delete(ctx, id)
	if err != nil {
		log.Fatalf("❌ 删除用户失败: %v", err)
	} else {
		fmt.Printf("✅ 用户删除成功，ID: %d\n", id)
	}

	// // 测试3: 根据ID查询用户
	// fmt.Println("\n2. 测试根据ID查询用户")
	// var foundUser *user.User
	// foundUser, err = userModel.FindOne(ctx, id)
	// if err != nil {
	// 	log.Fatalf("❌ 查询用户失败: %v", err)
	// }

	// if foundUser.Email != testUser.Email {
	// 	log.Fatalf("❌ 期望邮箱 %s, 实际 %s", testUser.Email, foundUser.Email)
	// }

	// fmt.Printf("✅ 用户查询成功: %s (%s)\n", foundUser.Username, foundUser.Email)

	// // 测试3: 根据邮箱查询用户

	// // 测试4: 根据用户名查询用户
	// fmt.Println("\n4. 测试根据用户名查询用户")
	// var usernameUser user.User
	// query = "SELECT * FROM `user` WHERE username = ? LIMIT 1"
	// err = conn.QueryRowCtx(ctx, &usernameUser, query, "testuser")
	// if err != nil {
	// 	log.Fatalf("❌ 根据用户名查询用户失败: %v", err)
	// }

	// if usernameUser.Email != testUser.Email {
	// 	log.Fatalf("❌ 期望邮箱 %s, 实际 %s", testUser.Email, usernameUser.Email)
	// }

	// fmt.Printf("✅ 根据用户名查询成功: %s\n", usernameUser.Email)

	// // 测试5: 更新用户
	// fmt.Println("\n5. 测试更新用户")
	// query = "UPDATE `user` SET username = ?, role = ? WHERE id = ?"
	// _, err = conn.ExecCtx(ctx, query, "updateduser", "admin", id)
	// if err != nil {
	// 	log.Fatalf("❌ 更新用户失败: %v", err)
	// }

	// // 验证更新结果
	// var updatedUser user.User
	// query = "SELECT * FROM `user` WHERE id = ? LIMIT 1"
	// err = conn.QueryRowCtx(ctx, &updatedUser, query, id)
	// if err != nil {
	// 	log.Fatalf("❌ 查询更新后的用户失败: %v", err)
	// }

	// if updatedUser.Username != "updateduser" {
	// 	log.Fatalf("❌ 期望用户名 updateduser, 实际 %s", updatedUser.Username)
	// }

	// if updatedUser.Role != "admin" {
	// 	log.Fatalf("❌ 期望角色 admin, 实际 %s", updatedUser.Role)
	// }

	// fmt.Printf("✅ 用户更新成功: %s, 角色: %s\n", updatedUser.Username, updatedUser.Role)

	// // 测试6: 删除用户
	// fmt.Println("\n6. 测试删除用户")
	// query = "DELETE FROM `user` WHERE id = ?"
	// _, err = conn.ExecCtx(ctx, query, id)
	// if err != nil {
	// 	log.Fatalf("❌ 删除用户失败: %v", err)
	// }

	// // 验证删除结果
	// var deletedUser user.User
	// query = "SELECT * FROM `user` WHERE id = ? LIMIT 1"
	// err = conn.QueryRowCtx(ctx, &deletedUser, query, id)
	// if err == nil {
	// 	log.Fatalf("❌ 用户应该已被删除")
	// }

	// fmt.Printf("✅ 用户删除成功\n")
}
