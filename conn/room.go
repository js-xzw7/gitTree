package conn

import (
	api "dance_two_tiktok_ui_cli/proto/gen/dreamDanceTiktok/proto"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var RoomInfo *Room

type LogTmp struct {
	Id  string
	Msg string
}

type Gift struct {
	Id   string
	Name string
}

var GiftList map[string]Gift

func init() {
	GiftList = make(map[string]Gift)
	GiftList["仙女棒"] = Gift{FairyWand, "仙女棒"}
	GiftList["能力药丸"] = Gift{Pills, "能力药丸"}
	GiftList["派对话筒"] = Gift{PartyMicrophone, "派对话筒"}
	GiftList["甜甜圈"] = Gift{Doughnuts, "甜甜圈"}
	GiftList["恶魔炸弹"] = Gift{GiftBomb, "恶魔炸弹"}
	GiftList["神秘空投"] = Gift{MysteriousDrop, "神秘空投"}
	GiftList["超能碰射"] = Gift{SuperImpact, "超能碰射"}
}

type Gui struct {
	RoomInfo  *widget.Label
	RoomModel *widget.Label
	List      *widget.List
	DrawList  *widget.List
	Barrage
	CpCombo1           *widget.Select
	CpCombo2           *widget.Select
	FriendRelationship *widget.RadioGroup
}

type Barrage struct {
	UserCombo *widget.Select
}

type Room struct {
	Gui
	RoomId  string
	Players []Player
	Logs    []LogTmp
	Friends map[string]*api.FriendChange
}

func InitRoom() {
	RoomInfo = &Room{
		Gui: Gui{
			List: widget.NewList(
				func() int {
					return len(RoomInfo.Logs)
				},
				func() fyne.CanvasObject {
					return container.NewHBox(widget.NewIcon(theme.DocumentIcon()), widget.NewLabel("Template Object"))
				},
				func(id widget.ListItemID, item fyne.CanvasObject) {
					item.(*fyne.Container).Objects[1].(*widget.Label).SetText(RoomInfo.Logs[id].Id)
				},
			),
		},
		Friends: make(map[string]*api.FriendChange),
	}

}
