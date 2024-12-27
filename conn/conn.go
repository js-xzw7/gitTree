package conn

import (
	router "dance_two_tiktok_ui_cli/router"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/zlog"
	"github.com/aceld/zinx/znet"
	"time"
)

var Client ziface.IClient

type Router struct {
	znet.BaseRouter
}

func StartClient() {
	Client = znet.NewClient("127.0.0.1", 8999)

	Client.SetOnConnStart(DoClientConnectedBegin)
	Client.SetOnConnStop(DoClientConnectedLost)

	//注册路由
	for msgId, r := range router.GetRouters() {
		Client.AddRouter(uint32(msgId), r)
	}

	Client.Start()
}

// DoClientConnectedBegin 创建连接的时候执行
func DoClientConnectedBegin(conn ziface.IConnection) {
	zlog.Debug("do connection begin is called ...")

	conn.SetProperty("createdTime", time.Now())
}

// DoClientConnectedLost 断开连接时执行
func DoClientConnectedLost(conn ziface.IConnection) {

	if roomId, err := conn.GetProperty("roomId"); err == nil {
		zlog.Debug("connection property roomId = ", roomId)
	}

	if createdTime, err := conn.GetProperty("createdTime"); err == nil {
		zlog.Debug("connection property create time = ", createdTime)
	}

	zlog.Debug("do client connected lost is called ...")
}

func Send(conn ziface.IConnection, msgId uint32, data []byte) error {
	return conn.SendMsg(msgId, data)
}
