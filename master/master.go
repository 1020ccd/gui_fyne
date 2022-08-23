package master

import (
	"gui/module"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)
type MasterWindow struct{

}

func(m *MasterWindow) Master(){
	a := app.New()
	a.Settings().SetTheme(theme.DarkTheme())

	master_window := a.NewWindow("GUI")
	
	master_window.Resize(fyne.NewSize(500,500))
	master_window.SetMaster()

	tabs := container.NewAppTabs(
		container.NewTabItem("进制转换", module.ModuleDTB()),
	
		container.NewTabItem("抽签工具", module.Draw(master_window)),
	)

	tabs.SetTabLocation(container.TabLocationTop)


	master_window.SetContent(tabs)
	master_window.ShowAndRun()
}