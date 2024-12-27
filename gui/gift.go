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
	"strconv"
	"time"
)

func Gift() *fyne.Container {
	var giftNumber int
	var barrage api.BarrageSend

	giftNumberInput := widget.NewEntry()
	giftNumberInput.SetPlaceHolder("赠送礼物数量")
	giftNumberInput.SetText("1")

	giftNumberInput.OnChanged = func(s string) {
		giftNumber, _ = strconv.Atoi(s)
	}
	var options []string

	for s, _ := range conn.GiftList {
		options = append(options, s)
	}

	combo := widget.NewSelect(options, func(value string) {
		gift, _ := conn.GiftList[value]
		var listLiveComment1 []dreamtwodancetiktokstruct.LiveGiftItem
		listLiveComment1 = append(listLiveComment1, dreamtwodancetiktokstruct.LiveGiftItem{
			MsgId:     uuid.New().String(),
			SecOpenid: userInfo.Id,
			SecGiftId: gift.Id,
			GiftNum:   giftNumber,
			GiftValue: 10,
			Test:      false,
			AvatarUrl: userInfo.AvatarUrl,
			Nickname:  userInfo.Name,
			Timestamp: time.Now().Unix(),
		})
		respInfoByte, _ := json.Marshal(listLiveComment1)
		barrage.BodyStr = string(respInfoByte)
	})

	button := widget.NewButton("送礼物", func() {
		if conn.Client == nil {
			return
		}
		barrage.MsgType = dreamtwodancetiktokenums.BarrageLiveGift
		var listLiveComment []dreamtwodancetiktokstruct.LiveGiftItem
		_ = json.Unmarshal([]byte(barrage.BodyStr), &listLiveComment)
		for i, item := range listLiveComment {
			item.MsgId = uuid.New().String()
			item.SecOpenid = userInfo.Id
			item.GiftNum = giftNumber
			item.Timestamp = time.Now().Unix()
			listLiveComment[i] = item
		}
		respInfoByte, _ := json.Marshal(listLiveComment)
		barrage.BodyStr = string(respInfoByte)

		barrage.Timestamp = helper.InterfaceHelperObject.ToString(time.Now().UnixMilli())
		barrage.RoomId = conn.RoomInfo.RoomId
		data, _ := proto.Marshal(&barrage)
		err := conn.Send(conn.Client.Conn(), uint32(api.MsgIds_TEST_BARRAGE_SEND), data)
		if err != nil {
			fmt.Println("发送失败", err)
		}
	})

	return container.NewGridWithColumns(1, container.NewVBox(giftNumberInput, combo, button))
}
