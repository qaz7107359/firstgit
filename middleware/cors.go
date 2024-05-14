package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		var headerKeys []string
		for k := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, "") // 获取请求的所有头部信息
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin,access-control-allow-headers,%s", headerStr)
		} else {
			headerStr = "access-control-allow-origin,access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT,DELETE,UPDATE")
			// 服务器支持的所有跨域请求的方法，为了避免浏览请求的多次‘预检’请求
			// header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization,Content-Length,X-CSRF-Token,Token,"+
				"session,X_Requested_With,Accept,Origin,Host,Connection,Accept-Encoding,Accept-Language,DNT,"+
				"X-CustomHeader,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,"+""+
				"Content-Type,Pragma")
			// 允许跨域设置
			// 可以返回其他字段
			c.Header("Access-Control-Expose-Headers", "Content-length,Access-Control-Allow-Origin,"+
				"Access-Control-Allow-Headers,Cache-Control-language，Control-Type,Expires,Last-Modified,Pragma,FooBar")
			// 跨域关键设置  让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")
			// 浏览器缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")
			// 跨域请求是否需要带cookie信息 默认设置为true
			c.Set("Content-type", "application/json")
			// 设置返回格式为json

		}
		// 放心所有的OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		c.Next() // 处理请求
	}
}
