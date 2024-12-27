package router

import (
	api "dance_two_tiktok_ui_cli/proto/gen/dreamDanceTiktok/proto"
	"encoding/json"
	"fmt"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/zlog"
	"github.com/gogo/protobuf/proto"
)

type ModelRouter struct {
	Router
}

func (r *ModelRouter) Handle(request ziface.IRequest) {
	zlog.Debug("Call ModelRouter Handle")
	resp := api.Response{}
	respInfo := api.CreateModelResponse{}
	_ = proto.Unmarshal(request.GetData(), &resp)
	_ = proto.Unmarshal(resp.Data, &respInfo)

	zlog.Debug("recv from server : msgId=", request.GetMsgID())
	fmt.Println("模式id为", respInfo.ModelId)

	request.GetConnection().SetProperty("modelId", respInfo.ModelId)
	respInfoByte, _ := json.Marshal(respInfo)
	//conn.RoomInfo.RoomModel.SetText(fmt.Sprintf("模式Id为:%+v", respInfo.ModelId))
	fmt.Println(string(respInfoByte))
}
