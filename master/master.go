package master

import (
	"gui/module"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)
type MasterWindow struct{

}

func(m *MasterWindow) Master(){
	a := app.New()
	
	master_window := a.NewWindow("GUI")
	
	master_window.Resize(fyne.NewSize(500,500))
	



	tabs := container.NewAppTabs(
		container.NewTabItem("进制转换", module.ModuleDTB()),
	
		container.NewTabItem("test2",widget.NewLabel("tetete")),
	)

	tabs.SetTabLocation(container.TabLocationTop)


	master_window.SetContent(tabs)
	master_window.ShowAndRun()
}