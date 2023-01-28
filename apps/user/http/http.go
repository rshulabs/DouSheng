// @Author: Ciusyan 2023/1/24
package http

import (
	"github.com/Go-To-Byte/DouSheng/apps/user"
	"github.com/Go-To-Byte/DouSheng/ioc"
	"github.com/gin-gonic/gin"
)

// 用于注入IOC中
var handler = &Handler{}

func NewUserHttpHandler() *Handler {
	return &Handler{}
}

// Handler 通过一个实体类，把内部接口用HTTP暴露出去【控制层Controller】
type Handler struct {
	service user.Service
}

// Version 当前模块API的版本
func (h *Handler) Version() string {
	return "v1"
}

// Registry 用于注册Handler所需要暴露的路由
func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/register", h.RegisterUser)
}

// Init 初始化Handler对象
func (h *Handler) Init() error {
	// 从IOC中获取UserServiceImpl实例
	h.service = ioc.GetInternalDependency(user.AppName).(user.Service)
	return nil
}

func (h *Handler) Name() string {
	return user.AppName
}

func init() {
	// 将此Gin服务注入IOC中
	ioc.GinDI(handler)
}