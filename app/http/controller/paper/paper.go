package paper

import (
	"github.com/gin-gonic/gin"
	"goinspect.cn/app/conf"
	"goinspect.cn/app/http"
	"goinspect.cn/app/http/response"
	"goinspect.cn/app/questionair/service/assess"
	"strconv"
)

func List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	psize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	service := assess.New(conf.Conf)

	papers, err := service.GetPaperPaginated(page, psize)
	if err != nil {
		response.Json(c, http.CODE_INTERNAL_ERROR, err.Error(), nil)
		return
	}
	response.Json(c, http.CODE_SUCCESS, "请求成功", gin.H{
		"list": papers,
	})
}

func Detail(c *gin.Context) {
	id, _ := strconv.Atoi(c.DefaultQuery("id", "0"))

	service := assess.New(conf.Conf)
	paper, err := service.GetPaper(id)
	if err != nil {
		response.Json(c, http.CODE_INTERNAL_ERROR, err.Error(), nil)
		return
	}
	response.Json(c, http.CODE_SUCCESS, "请求成功", paper)
}

func Questions(c *gin.Context) {
	response.Json(c, http.CODE_SUCCESS, "请求成功", gin.H{
		"id": 1,
		"title": "测评试卷1",
	})
}

func Submit(c *gin.Context) {
	response.Json(c, http.CODE_SUCCESS, "请求成功", gin.H{
		"id": 1,
		"title": "测评试卷1",
	})
}