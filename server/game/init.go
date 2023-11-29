package game

import (
	"mssgserver/net"
	"mssgserver/server/game/controller"
	"mssgserver/server/game/gameConfig"
)

var Router = &net.Router{}

func Init() {
	//加载基础配置
	gameConfig.Base.Load()
	initRouter()
}

func initRouter() {

	controller.DefaultRoleController.Router(Router)
}
