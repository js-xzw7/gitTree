package gui

import (
	"dance_two_tiktok_ui_cli/conn"
	api "dance_two_tiktok_ui_cli/proto/gen/dreamDanceTiktok/proto"
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/gogo/protobuf/proto"
	"github.com/google/uuid"
	"gitlab.dream288.com/backend/common-util-dream-two/dreamtwodancetiktok/enums/dreamtwodancetiktokenums"
	"gitlab.dream288.com/backend/common-util-dream-two/dreamtwodancetiktok/struct/dreamtwodancetiktokstruct"
	"gitlab.dream288.com/backend/common-util-dream-two/helper"
	"time"
)

func Like() *fyne.Container {
	button := widget.NewButton("点赞", func() {
		if conn.Client == nil {
			return
		}

		var barrageSend api.BarrageSend
		barrageSend.MsgType = dreamtwodancetiktokenums.BarrageLiveLike
		var listLiveComment []dreamtwodancetiktokstruct.LiveLikeItem

		listLiveComment = append(listLiveComment, dreamtwodancetiktokstruct.LiveLikeItem{
			MsgId:     uuid.New().String(),
			SecOpenid: userInfo.Id,
			LikeNum:   1,
			AvatarUrl: userInfo.AvatarUrl,
			Nickname:  userInfo.Name,
			Timestamp: time.Now().Unix(),
		})

		respByte, _ := json.Marshal(listLiveComment)
		barrageSend.BodyStr = string(respByte)
		barrageSend.Timestamp = helper.InterfaceHelperObject.ToString(time.Now().UnixMilli())
		barrageSend.RoomId = conn.RoomInfo.RoomId

		data, _ := proto.Marshal(&barrageSend)
		err := conn.Send(conn.Client.Conn(), uint32(api.MsgIds_TEST_BARRAGE_SEND), data)
		if err != nil {
			fmt.Println("发送失败")
		}

	})

	return container.NewGridWithColumns(1, container.NewVBox(button))
}
