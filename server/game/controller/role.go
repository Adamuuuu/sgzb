package controller

import (
	"mssgserver/net"
)

var DefaultRoleController = &RoleController{}

type RoleController struct {
}

func (r *RoleController) Router(router *net.Router) {
	g := router.Group("role")
	g.AddRouter("enterServer", r.enterServer)
}

func (r *RoleController) enterServer(req *net.WsMsgReq, rsp *net.WsMsgRsp) {
	//进入游戏的逻辑
	//session 需要验证是否合法
	//根据角色id  查询角色拥有的资源  如果有就返回  如果没有就初始化资源
	//log.Println(base.Base)

}
