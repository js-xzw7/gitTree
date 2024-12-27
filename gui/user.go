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

var userInfo conn.Player
var otherUser conn.Player

var UserOption map[string]conn.Player
var userOptionKey []string

func init() {
	UserOption = make(map[string]conn.Player)
	UserOption["_000LEeV-a-kfvv6x9j3u3IH3t0yXdVtNilv"] = conn.Player{
		Id:        "_000LEeV-a-kfvv6x9j3u3IH3t0yXdVtNilv",
		Name:      "玩家1",
		AvatarUrl: "https://p3.douyinpic.com/aweme/100x100/aweme-avatar/mosaic-legacy_3793_3131589739.jpeg?from=3067671334",
	}
	UserOption["_000lZwvjgjh7gWTDk-WzWVp-ISgZ0JGgnSb"] = conn.Player{
		Id:        "_000lZwvjgjh7gWTDk-WzWVp-ISgZ0JGgnSb",
		Name:      "玩家2",
		AvatarUrl: "https://p3.douyinpic.com/aweme/100x100/aweme-avatar/mosaic-legacy_3793_3131589739.jpeg?from=3067671334",
	}
	UserOption["1000"] = conn.Player{
		Id:        "1000",
		Name:      "玩家3",
		AvatarUrl: "https://p3.douyinpic.com/aweme/100x100/aweme-avatar/mosaic-legacy_3793_3131589739.jpeg?from=3067671334",
	}
	UserOption["2000"] = conn.Player{
		Id:        "2000",
		Name:      "玩家4",
		AvatarUrl: "https://p3.douyinpic.com/aweme/100x100/aweme-avatar/mosaic-legacy_3793_3131589739.jpeg?from=3067671334",
	}

	userOptionKey = []string{
		"_000LEeV-a-kfvv6x9j3u3IH3t0yXdVtNilv",
		"_000lZwvjgjh7gWTDk-WzWVp-ISgZ0JGgnSb",
		"1000",
		"2000",
	}
}

func User() *fyne.Container {
	userModelLabel := widget.NewLabel("设置用户信息")
	// 用户id input
	userIdInput := widget.NewEntry()
	userIdInput.SetPlaceHolder("用户ID")
	userIdInput.OnChanged = func(s string) {
		userInfo.Id = s
	}

	userNameInput := widget.NewEntry()
	userNameInput.SetPlaceHolder("用户name")
	userNameInput.OnChanged = func(s string) {
		userInfo.Name = s
	}

	userAvatarUrlInput := widget.NewEntry()
	userAvatarUrlInput.SetPlaceHolder("用户头像")
	userAvatarUrlInput.OnChanged = func(s string) {
		userInfo.AvatarUrl = s
	}

	userCombo := widget.NewSelect(userOptionKey, func(value string) {
		player, _ := UserOption[value]
		userIdInput.SetText(player.Id)
		userNameInput.SetText(player.Name)
		userAvatarUrlInput.SetText(player.AvatarUrl)
		userInfo = player
	})

	if userInfo.Id != "" {
		conn.RoomInfo.Players = append(conn.RoomInfo.Players, userInfo)
	}

	// 用户id input
	otherUserIdInput := widget.NewEntry()
	otherUserIdInput.SetPlaceHolder("用户ID")
	otherUserIdInput.OnChanged = func(s string) {
		otherUser.Id = s
	}

	otherUserNameInput := widget.NewEntry()
	otherUserNameInput.SetPlaceHolder("用户name")
	otherUserNameInput.OnChanged = func(s string) {
		otherUser.Name = s
	}

	otherUserAvatarUrlInput := widget.NewEntry()
	otherUserAvatarUrlInput.SetPlaceHolder("用户头像")
	otherUserAvatarUrlInput.OnChanged = func(s string) {
		otherUser.AvatarUrl = s
	}

	otherUserCombo := widget.NewSelect(userOptionKey, func(value string) {
		player, _ := UserOption[value]
		otherUserIdInput.SetText(player.Id)
		otherUserNameInput.SetText(player.Name)
		otherUserAvatarUrlInput.SetText(player.AvatarUrl)
		otherUser = player
	})

	if otherUser.Id != "" {
		conn.RoomInfo.Players = append(conn.RoomInfo.Players, otherUser)
	}

	return container.NewGridWithColumns(1, container.NewVBox(userModelLabel, userIdInput, userNameInput, userAvatarUrlInput, userCombo, otherUserIdInput, otherUserNameInput, otherUserAvatarUrlInput, otherUserCombo))
}

func Enter() *fyne.Container {
	//enterModelLabel := widget.NewLabel("用户进场")
	button := widget.NewButton("用户1进场", func() {})
	button.OnTapped = func() {
		var newBarrage api.BarrageSend

		newBarrage.MsgType = dreamtwodancetiktokenums.BarrageLiveComment
		var listLiveComment []dreamtwodancetiktokstruct.LiveCommentItem
		listLiveComment = append(listLiveComment, dreamtwodancetiktokstruct.LiveCommentItem{
			MsgId:     uuid.New().String(),
			SecOpenid: userInfo.Id,
			Content:   "1",
			AvatarUrl: userInfo.AvatarUrl,
			Nickname:  userInfo.Name,
			Timestamp: time.Now().Unix(),
		})

		respInfoByte, _ := json.Marshal(listLiveComment)
		newBarrage.BodyStr = string(respInfoByte)
		newBarrage.Timestamp = helper.InterfaceHelperObject.ToString(time.Now().UnixMilli())
		newBarrage.RoomId = conn.RoomInfo.RoomId
		data, _ := proto.Marshal(&newBarrage)
		err := conn.Send(conn.Client.Conn(), uint32(api.MsgIds_TEST_BARRAGE_SEND_RESPONSE), data)
		if err != nil {
			fmt.Println("下场失败")
		}

		button.SetText("用户1已经进场")
	}

	button2 := widget.NewButton("用户2进场", func() {})
	button2.OnTapped = func() {
		var newBarrage api.BarrageSend

		newBarrage.MsgType = dreamtwodancetiktokenums.BarrageLiveComment
		var listLiveComment []dreamtwodancetiktokstruct.LiveCommentItem
		listLiveComment = append(listLiveComment, dreamtwodancetiktokstruct.LiveCommentItem{
			MsgId:     uuid.New().String(),
			SecOpenid: otherUser.Id,
			Content:   "1",
			AvatarUrl: otherUser.AvatarUrl,
			Nickname:  otherUser.Name,
			Timestamp: time.Now().Unix(),
		})

		respInfoByte, _ := json.Marshal(listLiveComment)
		newBarrage.BodyStr = string(respInfoByte)
		newBarrage.Timestamp = helper.InterfaceHelperObject.ToString(time.Now().UnixMilli())
		newBarrage.RoomId = conn.RoomInfo.RoomId
		data, _ := proto.Marshal(&newBarrage)
		err := conn.Send(conn.Client.Conn(), uint32(api.MsgIds_TEST_BARRAGE_SEND_RESPONSE), data)
		if err != nil {
			fmt.Println("下场失败")
		}

		button2.SetText("用户2已经进场")
	}

	return container.NewGridWithColumns(2, button, button2)
}
