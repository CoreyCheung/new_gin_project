package controller

import (
	"fmt"
	"net/http"
	"new_gin_project/dto"
	"new_gin_project/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

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

	code, err := service.Login(req)
	if err != nil {
		logrus.Errorf("Login  err: %v", err.Error())
		resp.Code = code
		resp.Message = fmt.Sprintf("Login err: %v", err.Error())
		c.JSON(http.StatusOK, resp)
		return
	}

	c.JSON(http.StatusOK, dto.Response{Code: code, Message: "success"})
}
