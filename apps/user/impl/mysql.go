// @Author: Ciusyan 2023/1/23
package impl

import (
	"github.com/Go-To-Byte/DouSheng/apps/user"
	"github.com/Go-To-Byte/DouSheng/conf"
	"github.com/Go-To-Byte/DouSheng/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"gorm.io/gorm"
)

// 用于注入IOC中
var userServiceImpl = &UserServiceImpl{}

func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{
		// User模块服务的子Logger
		l:  zap.L().Named("User"),
		db: conf.C().MySQL.GetDB(),
	}
}

// UserServiceImpl 基于Mysql实现的Service
type UserServiceImpl struct {
	// 日志实例
	l  logger.Logger
	db *gorm.DB
}

func (u *UserServiceImpl) Init() error {
	u.l = zap.L().Named("User")
	u.db = conf.C().MySQL.GetDB()

	return nil
}

func (u *UserServiceImpl) Name() string {
	return user.AppName
}

func init() {
	// 将此UserService注入内部服务的IOC容器中
	ioc.InternalDI(userServiceImpl)
}