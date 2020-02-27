package wx

import (
	"crypto/sha1"
	"fmt"
	"github.com/gin-gonic/gin"
	"goinspect.cn/app/conf"
	"io"
	"sort"
)

func Login(c *gin.Context) {

}

func LoginCallback(c *gin.Context) {

}

func ServiceVerify(c *gin.Context) {
	echostr, _ := c.GetQuery("echostr")
	c.String(200, echostr)
	return
	sig, _ := c.GetQuery("signature")
	ts, _ := c.GetQuery("timestamp")
	nonce, _ := c.GetQuery("nonce")
	//echostr, _ := c.GetQuery("echostr")

	token := conf.Conf.Weixin.VerifyToken

	sArr := []string{token, ts, nonce}
	sort.Strings(sArr)
	sStr := sArr[0] + sArr[1] + sArr[2]

	t := sha1.New()
	io.WriteString(t, sStr)
	sha1 := fmt.Sprintf("%x", t.Sum(nil))

	if sha1 == sig {
		c.String(200, echostr)
	} else {
		c.String(200, "access denied")
	}
}