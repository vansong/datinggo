package todolist

import (
	"github.com/gin-gonic/gin"
	"goinspect.cn/app/conf"
	"goinspect.cn/app/http"
	"goinspect.cn/app/http/response"
	"goinspect.cn/app/todolist/service"
	"strconv"
)

func Index(c *gin.Context) {
	uid := c.GetInt("loggedin")

	s := service.New(conf.Conf)
	todo, done, err := s.GetUserTasks(uid)
	if err != nil {
		response.Json(c, http.CODE_INTERNAL_ERROR, err.Error(), nil)
		return
	}

	response.Success(c, gin.H{
		"todo": todo,
		"done": done,
	})
}

func New(c *gin.Context) {
	uid := c.GetInt("loggedin")
	title := c.PostForm("title")

	// todo 参数检查

	s := service.New(conf.Conf)
	res := s.NewTask(uid, title)
	if !res {
		response.Json(c, http.CODE_INTERNAL_ERROR, "保存失败", nil)
		return
	}
	response.Success(c, nil)
}

func SetDone(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		response.Json(c, http.CODE_INVALID_REQUEST, "参数错误", nil)
		return
	}

	s := service.New(conf.Conf)

	if !s.SetTaskDone(id) {
		response.Json(c, http.CODE_INTERNAL_ERROR, "保存失败", nil)
		return
	}
	response.Success(c, nil)
}

func Remove(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		response.Json(c, http.CODE_INVALID_REQUEST, "参数错误", nil)
		return
	}

	s := service.New(conf.Conf)

	if !s.RemoveTask(id) {
		response.Json(c, http.CODE_INTERNAL_ERROR, "保存失败", nil)
		return
	}
	response.Success(c, nil)
}