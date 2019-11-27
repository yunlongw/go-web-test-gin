package api

import (
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-web-test-gin/models"
	"go-web-test-gin/pkg/e"
	"go-web-test-gin/pkg/logging"
	"go-web-test-gin/pkg/queue"
	"go-web-test-gin/pkg/util"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	valid := validation.Validation{}
	a := auth{
		Username: username,
		Password: password,
	}

	ok, _ := valid.Valid(&a)
	data := make(map[string]interface{})
	code := e.INVALID_PARAMS

	if ok {
		auth, isExist := models.CheckAuth(username, password) // 账号密码检查
		if isExist {
			token, err := util.GenerateToken(auth) // token 生成
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}

	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	loginTask()

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func loginTask() {
	var (
		uid = uuid.New().String()
	)

	signature := &tasks.Signature{
		UUID: uid,
		Name: "login",
		Args: []tasks.Arg{
			{
				Type:  "int64",
				Value: 1,
			},
			{
				Type:  "int64",
				Value: 1,
			},
		},
		RetryCount:   2,
		RetryTimeout: 3,
	}

	asyncResult, err := queue.MServer.SendTask(signature)
	if err != nil {
		panic(err.Error())
	}
	logging.Info(asyncResult)
}
