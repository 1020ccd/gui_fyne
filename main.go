package main

import (
	"os"
	"runtime"
	"fmt"
	"strings"
	"gui/master"
	findfont "github.com/flopp/go-findfont"
	
)

func init() {
    sysType := runtime.GOOS
    fontPaths := findfont.List()
    for _, path := range fontPaths {
        // fmt.Println(path)
        //楷体:simkai.ttf
        //黑体:simhei.ttf

        switch sysType {
        case "windows":
            if strings.Contains(path, "simkai.ttf") {
                os.Setenv("FYNE_FONT", path)

            }
        case "darwin":
            if strings.Contains(path, "STHeiti Medium.ttc") {
                os.Setenv("FYNE_FONT", path)

            }
        }

    }
    fmt.Println(sysType)
}

func main(){
	var master_window master.MasterWindow
	master_window.Master()
}