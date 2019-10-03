package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"go-web-test/pkg/e"
	"net/http"
)
/**
绑定和验证参数
 */
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)

	if err != nil {
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)

	if err != nil {
		return http.StatusInternalServerError, e.ERROR
	}

	if !check {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	return http.StatusOK, e.SUCCESS
}
