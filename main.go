package main

import (
	"os"
	// "fmt"
	"strings"
	"gui/master"
	findfont "github.com/flopp/go-findfont"
	
)

func init(){
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		// fmt.Println(path)
		//楷体:simkai.ttf
		//黑体:simhei.ttf
		if strings.Contains(path, "STHeiti Medium.ttc") {
			os.Setenv("FYNE_FONT", path)
			break
		}
	}
}

func main(){
	var master_window master.MasterWindow
	master_window.Master()
}