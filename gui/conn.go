package gui

import (
	"dance_two_tiktok_ui_cli/conn"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"time"
)

func ConnContainer() *fyne.Container {
	connLabel := widget.NewLabel("建立连接")
	button := widget.NewButton("连接", func() {})
	button.OnTapped = func() {
		if conn.Client == nil {
			go conn.StartClient()

			for {
				if conn.Client != nil {
					connLabel.SetText("建立连接成功")
					button.Hide()
					break
				} else {
					time.Sleep(time.Second * 1)
				}
			}
		}
	}

	connNewVBox := container.NewVBox(connLabel, button)
	return connNewVBox
}
