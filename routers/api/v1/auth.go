package v1

import (
	"github.com/gin-gonic/gin"
	"go-web-test-gin/pkg/app"
	"go-web-test-gin/pkg/e"
	"go-web-test-gin/service/auth_service"
	"net/http"
)

type AddAuth struct {
	UserName string `form:"username" valid:"Required"`
	PassWord string `form:"password" valid:"Required"`
}

func Register(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddAuth
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	authService := auth_service.Member{UserName: form.UserName}
	exists, err := authService.ExistByUserName()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_MEMBER, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_EXIST_MEMBER, nil)
		return
	}

	auth := auth_service.Member{
		UserName: form.UserName,
		PassWord: form.PassWord,
	}

	err = auth.AddMember()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_REGISTER_MEMBER, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}
