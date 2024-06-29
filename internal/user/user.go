package user

import (
	"database/sql"
	"fmt"
	"game-server/internal/db"
	"log"
	"strings"

	"github.com/valyala/fasthttp"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int
	Username     string
	Password     string
	CreatedAt    string
	AvatarBase64 string
}

func Register(ctx *fasthttp.RequestCtx) {
	if !ctx.IsPost() {
		ctx.Error("只支持 POST 方法", fasthttp.StatusMethodNotAllowed)
		return
	}
	username := string(ctx.FormValue("username"))
	password := string(ctx.FormValue("password"))
	avatarBase64 := string(ctx.FormValue("avatar"))

	if username == "" || password == "" {
		ctx.Error("用户名和密码不能为空", fasthttp.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.Error("服务器错误", fasthttp.StatusInternalServerError)
		return
	}

	_, err = db.DB.Exec("INSERT INTO users (username, password, avatar_base64) VALUES (?, ?, ?)", username, string(hashedPassword), avatarBase64)
	if err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), "Duplicate entry") {
			ctx.Error("用户已存在", fasthttp.StatusConflict)
		} else {
			ctx.Error("数据库操作失败", fasthttp.StatusInternalServerError)
		}

		return
	}

	fmt.Fprintf(ctx, "用户注册成功")
}

func Login(ctx *fasthttp.RequestCtx) {
	// 只允许 POST 方法
	if !ctx.IsPost() {
		ctx.Error("只支持 POST 方法", fasthttp.StatusMethodNotAllowed)
		return
	}

	// 从请求中获取用户信息
	username := string(ctx.FormValue("username"))
	password := string(ctx.FormValue("password"))

	if username == "" || password == "" {
		ctx.Error("用户名和密码不能为空", fasthttp.StatusBadRequest)
		return
	}

	// 从数据库中获取用户信息
	var storedPassword string
	var avatarBase64 string
	err := db.DB.QueryRow("SELECT password, avatar_base64 FROM users WHERE username = ?", username).Scan(&storedPassword, &avatarBase64)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.Error("用户名或密码错误", fasthttp.StatusUnauthorized)
		} else {
			ctx.Error("服务器错误", fasthttp.StatusUnauthorized)
		}

		return
	}

	// 比较密码
	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	if err != nil {
		ctx.Error("用户名或密码错误", fasthttp.StatusUnauthorized)
		return
	}

	fmt.Fprintf(ctx, "用户登录成功。头像数据（Base64）：%s", avatarBase64)
}
