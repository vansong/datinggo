package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goinspect.cn/app/conf"
	"goinspect.cn/app/http"
	"goinspect.cn/app/http/response"
	"goinspect.cn/app/account/service"
)

func Login(c *gin.Context) {

}

func MyProfile(c *gin.Context) {
	uid := c.GetInt("loggedin")
	fmt.Print(uid);
	if uid <= 0 {
		response.Error(c, http.CODE_NOT_LOGGEDIN, "当前未登录")
		return
	}

	s := service.New(conf.Conf)
	u := s.GetUser(uid)
	if u == nil {
		response.Error(c, http.CODE_NOT_FOUND, "用户不存在")
		return
	}

	response.Success(c, u)
}
