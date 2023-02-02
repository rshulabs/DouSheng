// Author: BeYoung
// Date: 2023/2/2 21:07
// Software: GoLand

package init

import (
	"github.com/Go-To-Byte/DouSheng/apps/video/models"
	"github.com/Go-To-Byte/DouSheng/apps/video/service"
)

func initPort() {
	port, err := service.GetFreePort()
	if err == nil {
		models.Config.Localhost.Port = port
	}
}
