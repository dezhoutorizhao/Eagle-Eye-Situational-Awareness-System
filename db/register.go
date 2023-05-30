package db

import (
	"bytes"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var (
	re_username string
	re_password string
	re_email    string
	re_number   string
	re_role     string
)

// re代表register
type Re_User struct {
	re_username string `json:"username"`
	re_password string `json:"password"`
	re_email    string `json:"email"`
	re_number   string `json:"school_id"`
	re_role     string `json:"re_role"`
}

// 接收从前端返回过来的数据部分
func AddUser_front(add_c *gin.Context) {
	re_username = add_c.PostForm("username")
	re_password = add_c.PostForm("password")
	re_email = add_c.PostForm("email")
	re_number = add_c.PostForm("school_id")
	re_role = add_c.PostForm("role")
	fmt.Println(re_username, re_password, re_email, re_role)
	if len(re_username) == 0 || len(re_password) == 0 || len(re_email) == 0 || len(re_number) == 0 || len(re_role) == 0 {
		add_c.String(200, "error\n")
		return
	}
	if len(re_username) == 0 {
		add_c.String(200, "username is null\n")
		return
	}
	if len(re_password) == 0 {
		add_c.String(200, "password is null\n")
		return
	}
	if len(re_email) == 0 {
		add_c.String(200, "email is null\n")
		return
	}
	if len(re_number) == 0 {
		add_c.String(200, "school_id is null\n")
		return
	}
	if len(re_role) == 0 {
		add_c.String(200, "role is null\n")
		return
	}

	// re代表register
	u := &Re_User{re_username, re_password, re_email, re_number, re_role}
	u.AddUser()
}

// AddUser 添加用户的方法一
func (user *Re_User) AddUser() error {
	//写sql语句
	sqlStr := "insert into register_users(username,password,email,number,role) values(?,?,?,?,?)"
	//预编译
	fmt.Println(sqlStr)
	inStmt, err := Db.Prepare(sqlStr) //预编译得到的是inStmt,通过操作inStmt得到不同的结果
	if err != nil {
		fmt.Println("预编译出现异常", err)
		return err
	}
	//3.执行
	fmt.Println(inStmt)
	_, err2 := inStmt.Exec(re_username, re_password, re_email, re_number, re_role)
	if err2 != nil {
		fmt.Println("执行出现异常", err2)
		return err2
	}
	return nil
}

// 中间件，处理session
func Session(keyPairs string) gin.HandlerFunc {
	store := SessionConfig()
	return sessions.Sessions(keyPairs, store)
}
func SessionConfig() sessions.Store {
	sessionMaxAge := 3600
	sessionSecret := "topgoer"
	var store sessions.Store
	store = cookie.NewStore([]byte(sessionSecret))
	store.Options(sessions.Options{
		MaxAge: sessionMaxAge, //seconds
		Path:   "/",
	})
	return store
}

func Captcha(c *gin.Context, length ...int) {
	l := captcha.DefaultLen
	w, h := 107, 36
	if len(length) == 1 {
		l = length[0]
	}
	if len(length) == 2 {
		w = length[1]
	}
	if len(length) == 3 {
		h = length[2]
	}
	captchaId := captcha.NewLen(l)
	session := sessions.Default(c)
	session.Set("captcha", captchaId)
	_ = session.Save()
	_ = Serve(c.Writer, c.Request, captchaId, ".png", "zh", false, w, h)
}

func CaptchaVerify(c *gin.Context, code string) bool {
	session := sessions.Default(c)
	if captchaId := session.Get("captcha"); captchaId != nil {
		session.Delete("captcha")
		_ = session.Save()
		if captcha.VerifyString(captchaId.(string), code) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
func Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		_ = captcha.WriteImage(&content, id, width, height)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		_ = captcha.WriteAudio(&content, id, lang)
	default:
		return captcha.ErrNotFound
	}

	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}
