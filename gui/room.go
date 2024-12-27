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

func RoomCard() *fyne.Container {
	roomLabel := widget.NewLabel("房间信息")
	button := widget.NewButton("创建房间", func() {})
	button.OnTapped = func() {
		if conn.Client == nil {
			roomLabel.SetText("房间信息:未建立连接")
			return
		}
		if conn.RoomInfo == nil {
			conn.InitRoom()
			conn.RoomInfo.RoomInfo = roomLabel

			req := &api.CreateRoomRequest{
				Token: "1",
			}

			data, _ := proto.Marshal(req)

			err := conn.Send(conn.Client.Conn(), uint32(api.MsgIds_CREATE_ROOM_REQUEST), data)
			if err != nil {
				fmt.Println("创建房间请求失败：", err)
			} else {
				button.Hide()
			}
		}

		roomLabel.SetText("创建房间成功")
	}
	roomInput := widget.NewEntry()
	roomInput.SetPlaceHolder("请输入房间id")
	roomInput.OnChanged = func(s string) {
		conn.RoomInfo.RoomId = s
	}

	return container.NewGridWithColumns(1, roomLabel, button, roomInput)
}
