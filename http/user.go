package http

import (
	"fmt"
	"net/http"
	"new_gin_project/dto"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Summary 登录api
// @Tags 权限管理
// @version 1.0
// @Accept application/json
// @Param req  body dto.ReqLogin
// @Success 200 object dto.Response 成功后返回值
// @Failure 200 object dto.Response 查询失败
// @Router /login [post]
func Login(c *gin.Context) {
	req := dto.ReqLogin{}
	resp := dto.Response{}

	if err := c.Bind(&req); err != nil {
		logrus.Errorf("Login json unmarshal err: %v", err.Error())
		resp.Code = -1
		resp.Message = fmt.Sprintf("json unmarshal err: %v", err.Error())
		c.JSON(http.StatusOK, resp)
		return
	}

	code, err := srv.Login(req)
	if err != nil {
		logrus.Errorf("Login  err: %v", err.Error())
		resp.Code = code
		resp.Message = fmt.Sprintf("Login err: %v", err.Error())
		c.JSON(http.StatusOK, resp)
		return
	}

	c.JSON(http.StatusOK, dto.Response{Code: code, Message: "success"})
}
