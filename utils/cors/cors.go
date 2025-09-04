package cors

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewCORS 创建一个新的CORS中间件，接收用户自定义配置
// 未指定的配置项将自动使用默认值（默认允许所有）
// NewCORS creates a new CORS middleware with user-defined configuration
// Unspecified configuration items will automatically use default values (allow all by default)
func NewCORS(userConfig Config) gin.HandlerFunc {
	// 从默认配置开始
	// Start with default configuration
	config := DefaultConfig()

	// 合并用户配置（用户指定的项会覆盖默认值）
	// Merge user configuration (user-specified items override defaults)
	if len(userConfig.AllowOrigins) > 0 {
		config.AllowOrigins = userConfig.AllowOrigins
	}
	if len(userConfig.AllowMethods) > 0 {
		config.AllowMethods = userConfig.AllowMethods
	}
	if len(userConfig.AllowHeaders) > 0 {
		config.AllowHeaders = userConfig.AllowHeaders
	}
	if len(userConfig.ExposeHeaders) > 0 {
		config.ExposeHeaders = userConfig.ExposeHeaders
	}
	if userConfig.MaxAge > 0 {
		config.MaxAge = userConfig.MaxAge
	}
	// 布尔值无论真假都应该由用户配置决定
	// Boolean value should be determined by user configuration regardless of true/false
	config.AllowCredentials = userConfig.AllowCredentials

	// 转换为gin-cors需要的配置格式并返回中间件
	// Convert to configuration format required by gin-cors and return middleware
	return cors.New(cors.Config{
		AllowOrigins:     config.AllowOrigins,
		AllowMethods:     config.AllowMethods,
		AllowHeaders:     config.AllowHeaders,
		ExposeHeaders:    config.ExposeHeaders,
		AllowCredentials: config.AllowCredentials,
		MaxAge:           config.MaxAge,
	})
}

// CORS 使用默认配置（允许所有）的快捷函数
// 等价于NewCORS(DefaultConfig())
// CORS is a shortcut function using default configuration (allow all)
// Equivalent to NewCORS(DefaultConfig())
func CORS() gin.HandlerFunc {
	return NewCORS(DefaultConfig())
}
