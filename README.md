# go-web

`go-web` 是一个基于 Go 语言的 Web 开发工具库，旨在提供常用的 Web 开发组件和工具函数，帮助简化 Go Web 应用的开发流程，尤其适用于基于 Gin 框架的项目。


## 安装

使用 `go get` 命令安装：

```bash
go get github.com/NexusXian/go-web
```


## 主要功能

### 1. CORS 中间件
提供灵活的跨域资源共享（CORS）配置，支持自定义允许的来源、方法、请求头、响应头等，默认配置允许所有跨域请求，便于快速开发。

#### 示例使用：
```go
import (
    "github.com/gin-gonic/gin"
    "github.com/NexusXian/go-web/utils/cors"
)

func main() {
    r := gin.Default()
    
    // 使用默认CORS配置（允许所有）
    r.Use(cors.CORS())
    
    // 或使用自定义配置
    customConfig := cors.Config{
        AllowOrigins:     []string{"https://example.com"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Content-Type", "Authorization"},
        AllowCredentials: true,
        MaxAge:           24 * time.Hour,
    }
    r.Use(cors.NewCORS(customConfig))
    
    // 路由定义...
    r.Run(":8080")
}
```


### 2. Bcrypt 密码处理
封装了 bcrypt 算法的密码哈希与验证功能，支持自定义工作因子（哈希强度），便于安全地存储和验证用户密码。

#### 示例使用：
```go
import (
    "fmt"
    "github.com/NexusXian/go-web/utils/bcrypt"
)

func main() {
    // 哈希密码（使用默认工作因子）
    password := "user123"
    hash, err := bcrypt.HashPassword(password)
    if err != nil {
        fmt.Println("哈希失败：", err)
        return
    }
    fmt.Println("哈希后的密码：", hash)
    
    // 验证密码
    isValid := bcrypt.CheckPassword(password, hash)
    fmt.Println("密码验证结果：", isValid) // 输出: true
}
```


## 依赖

主要依赖项包括：
- [gin-gonic/gin](https://github.com/gin-gonic/gin)：Go 语言的高性能 Web 框架
- [gin-contrib/cors](https://github.com/gin-contrib/cors)：Gin 框架的 CORS 中间件
- [golang.org/x/crypto/bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)：bcrypt 密码哈希算法实现

