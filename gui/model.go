package gui

import (
	"dance_two_tiktok_ui_cli/conn"
	api "dance_two_tiktok_ui_cli/proto/gen/dreamDanceTiktok/proto"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/gogo/protobuf/proto"
)

var modelOption map[string]int64
var modelOptionKey []string

var modelType int64

func init() {
	modelOption = make(map[string]int64, 0)
	modelOption["普通"] = 1
	modelOption["结婚"] = 3

	modelOptionKey = []string{"普通", "结婚"}
}

func ModelCard() *fyne.Container {
	roomModelLabel := widget.NewLabel("模式信息")

	modelSelect := widget.NewSelect(modelOptionKey, func(value string) {
		if conn.Client == nil {
			roomModelLabel.SetText("模式信息：未建立连接")
			return
		}
		modelType = modelOption[value]
		roomModelLabel.SetText(fmt.Sprintf("已选择：%+v", modelType))
		if conn.RoomInfo != nil {
			conn.RoomInfo.RoomModel = roomModelLabel
			req := &api.CreateModelRequest{
				ModelType: int32(modelType),
			}

			data, _ := proto.Marshal(req)
			err := conn.Send(conn.Client.Conn(), uint32(api.MsgIds_CREATE_MODEL_REQUEST), data)
			if err != nil {
				fmt.Println("创建模式请求失败：", err)
			}
		}
	})

	return container.NewGridWithColumns(1, roomModelLabel, modelSelect)
}
