// Package limiter 处理限流逻辑
package limiter

import (
	"github.com/gin-gonic/gin"
	"strings"
)

// GetKeyIP 获取 Limiter 的key,IP
func GetKeyIP(c *gin.Context) string  {
	return c.ClientIP()
}

// GetKeyRouteWithIP Limitor 的 Key，路由+IP，针对单个路由做限流
func GetKeyRouteWithIP(c *gin.Context) string  {
	return routeToKeyString(c.FullPath()) + c.ClientIP()
}

// routeToKeyString 辅助方法，将 URL 中的 / 格式为 -
func routeToKeyString(routeName string) string  {
	routeName = strings.ReplaceAll(routeName,"/","-")
	routeName = strings.ReplaceAll(routeName,"/","_")
	return routeName
}