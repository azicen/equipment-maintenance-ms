package middleware

import (
	"os"
	"regexp"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//InitCors 跨域配置
func InitCors() gin.HandlerFunc {
	//读取配置文件
	s := strings.Split(os.Getenv("ALLOW_ORIGINS"), ";")

	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie"}
	if gin.Mode() == gin.ReleaseMode {
		//生产环境需要配置跨域域名，否则403
		//config.AllowOrigins = []string{"http://www.example.com"}
		config.AllowOrigins = s
	} else {
		//测试环境下模糊匹配本地开头的请求
		config.AllowOriginFunc = func(origin string) bool {
			if regexp.MustCompile(`^http://127\.0\.0\.1:\d+$`).MatchString(origin) {
				return true
			}
			if regexp.MustCompile(`^http://localhost:\d+$`).MatchString(origin) {
				return true
			}
			return false
		}
	}
	config.AllowCredentials = true
	return cors.New(config)
}
