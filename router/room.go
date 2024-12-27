package router

import (
	api "dance_two_tiktok_ui_cli/proto/gen/dreamDanceTiktok/proto"
	"fmt"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/zlog"
	"github.com/gogo/protobuf/proto"
)

type RoomRouter struct {
	Router
}

func (r *RoomRouter) Handle(request ziface.IRequest) {
	zlog.Debug("Call RoomRouter Handle")
	resp := api.Response{}
	createRoomResp := api.CreateRoomResponse{}

	_ = proto.Unmarshal(request.GetData(), &resp)
	_ = proto.Unmarshal(resp.Data, &createRoomResp)

	if createRoomResp.RoomId != 0 {
		request.GetConnection().SetProperty("roomId", createRoomResp.RoomId)
		fmt.Println("直播间id = ", createRoomResp.RoomId)
	}

	fmt.Print("创建直播间返回")
}
