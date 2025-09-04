package cors

import (
	"time"
)

// Config 定义CORS中间件的配置选项
// Config defines configuration options for CORS middleware
type Config struct {
	AllowOrigins []string // 允许的请求来源，默认值为["*"]表示允许所有来源
	// Allowed origins, default is ["*"] which allows all origins
	AllowMethods []string // 允许的HTTP方法，默认值为["*"]表示允许所有方法
	// Allowed HTTP methods, default is ["*"] which allows all methods
	AllowHeaders []string // 允许的请求头，默认值为["*"]表示允许所有请求头
	// Allowed request headers, default is ["*"] which allows all headers
	ExposeHeaders []string // 允许客户端访问的响应头，默认值为["*"]表示允许所有响应头
	// Response headers exposed to clients, default is ["*"] which exposes all headers
	AllowCredentials bool // 是否允许携带认证信息（如cookies），默认值为false
	// Whether to allow credentials (e.g. cookies), default is false
	MaxAge time.Duration // 预检请求的缓存时间，默认值为12小时
	// Cache duration for preflight requests, default is 12 hours
}

// DefaultConfig 返回默认的CORS配置（默认允许所有来源、方法和头信息）
// DefaultConfig returns default CORS configuration (allows all origins, methods and headers by default)
func DefaultConfig() Config {
	return Config{
		AllowOrigins: []string{"*"}, // 默认允许所有来源
		// Allow all origins by default
		AllowMethods: []string{"*"}, // 默认允许所有HTTP方法
		// Allow all HTTP methods by default
		AllowHeaders: []string{"*"}, // 默认允许所有请求头
		// Allow all request headers by default
		ExposeHeaders: []string{"*"}, // 默认暴露所有响应头
		// Expose all response headers by default
		AllowCredentials: false, // 默认不允许携带认证信息（出于安全考虑）
		// Do not allow credentials by default (for security)
		MaxAge: 12 * time.Hour, // 预检请求默认缓存12小时
		// Preflight requests cached for 12 hours by default
	}
}
