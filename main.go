package main

import (
	"dance_two_tiktok_ui_cli/gui"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	myApp := app.NewWithID("myApp")
	myWin := myApp.NewWindow("My Window")

	// 按钮：建立链接
	connNewVBox := gui.ConnContainer()
	// 创建房间按钮
	// 显示房间Id
	roomCard := gui.RoomCard()
	// 创建模式
	model := gui.ModelCard()

	//选择上场用户
	user := gui.User()

	//用户下场
	enter := gui.Enter()

	left := container.NewGridWithColumns(1, container.NewVBox(connNewVBox, roomCard, model, user, enter))

	//弹幕
	barrage := gui.Barrage()

	//点赞
	like := gui.Like()

	//送礼物
	gift := gui.Gift()

	center := container.NewGridWithColumns(1, container.NewVBox(barrage, like, gift))
	containerList := container.NewGridWithColumns(3, left, center)

	myWin.SetContent(containerList)

	myWin.Resize(fyne.NewSize(800, 600))
	myWin.ShowAndRun()

}
