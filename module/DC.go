package module

import (

	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var base string  //进制类型
var calc_container,calc_container1,calc_container2,calc_container3,calc_container4  *fyne.Container //计算结果container
var container_ret *fyne.Container
var input *widget.Entry
func ModuleDTB() *fyne.Container{
	
	//输入框
	input = widget.NewEntry()
	
	box := container.NewVBox(input)
	box.Resize(fyne.NewSize(170,30))
	box.Move(fyne.NewPos(150,100))
	
	
	//单选框
	
	radio := widget.NewRadioGroup([]string{"DEC","HEX", "OCT","BIN"}, 
	func(value string ) {
		// log.Println("Radio set to", PlaceHolder)
		base = value
		PlaceHolder := "请输入"+value+"数据"
		input.SetPlaceHolder(PlaceHolder)//动态修改输入框内的文本
		input.SetText("")
		
	})
	
	radio.Horizontal = true
	radio.Required = true
	radio.SetSelected("DEC")
	box_radio := container.NewHBox(radio)
	box_radio.Move(fyne.NewPos(100,40))



	//RESULT
	RESULT := widget.NewLabel("RESULT:")
	RESULT.Alignment = fyne.TextAlignTrailing
	RESULT.Move(fyne.NewPos(130,150))

	//HEX
	HEX := widget.NewLabel("HEX:")
	HEX.Alignment = fyne.TextAlignTrailing
	HEX.Move(fyne.NewPos(130,200))

	//OCT
	OCT := widget.NewLabel("OCT:")
	OCT.Alignment = fyne.TextAlignTrailing
	OCT.Move(fyne.NewPos(130,250))

	//BIN
	BIN := widget.NewLabel("BIN:")
	BIN.Alignment = fyne.TextAlignTrailing
	BIN.Move(fyne.NewPos(130,300))

	//DEC
	DEC := widget.NewLabel("DEC:")
	DEC.Alignment = fyne.TextAlignTrailing
	DEC.Move(fyne.NewPos(130,350))

	//页面显示的各个容器
	container_ret = container.NewWithoutLayout(
		box_radio,
		box,
		RESULT,
		DEC,
		HEX,
		OCT,
		BIN,
	)

	//通过按钮计算结果
	// calcBTN := container.NewVBox(widget.NewButton("calc", func() {
	// 	// log.Println("Content was:", input.Text)
	// 	container_ret.Remove(calc_container)
	// 	calc_container = calc_ret("ok",input.Text, base, 160,350)
	// 	container_ret.Add(calc_container)
		
	// }))
	// calcBTN.Move(fyne.NewPos(400,100))
	

	// 计算结果生成container到页面
	input.OnChanged = func(t string) {

		clearRet(container_ret,calc_container,calc_container1,calc_container2,calc_container3,calc_container4)
		calc_container,calc_container1,calc_container2,calc_container3,calc_container4 = calc_ret(t, base)
		container_ret.Add(calc_container)
		container_ret.Add(calc_container1)
		container_ret.Add(calc_container2)
		container_ret.Add(calc_container3)
		container_ret.Add(calc_container4)
	}

	//添加按钮到容器
	// container_ret.Add(calcBTN)
	
	return container_ret
}

func calc_ret(val string, base string) (*fyne.Container,*fyne.Container,*fyne.Container,*fyne.Container,*fyne.Container){
	//clearRet(container_ret,calc_container,calc_container1,calc_container2,calc_container3,calc_container4)
	switch base{
		
		case "DEC":
			// input.Text = ""
			
			num, err := strconv.Atoi(val)
			if err != nil{
				str := "输入有误"
				if val == ""{
					str = ""
				}
				
				return widget_of_label(str,"","","","")
			}
			val_ret := "succeed"
			val_2 := dto(num,"2")
			val_8 := dto(num,"8")
			val_10 := dto(num,"10")
			val_16 := dto(num,"16")
			return widget_of_label(val_ret,val_16,val_10,val_2,val_8)
				
		case "HEX":
			// input.Text = ""
			
			val_10,err := otherTo(val,"16")
			if err != nil{
				str := "输入有误"
				if val == ""{
					str = ""
				}
				
				return widget_of_label(str,"","","","")
			}
			val_ret := "succeed"
			val_2 := dto(val_10,"2")
			val_8 := dto(val_10,"8")
			val_16 := dto(val_10,"16")
			val_10_tem := strconv.Itoa(val_10)
			return widget_of_label(val_ret,val_16,val_10_tem,val_2,val_8)

		case "BIN":
			val_10,err := otherTo(val,"2")
			if err != nil{
				str := "输入有误"
				if val == ""{
					str = ""
				}
				
				return widget_of_label(str,"","","","")
			}
			val_ret := "succeed"
			val_2 := dto(val_10,"2")
			val_8 := dto(val_10,"8")
			val_16 := dto(val_10,"16")
			val_10_tem := strconv.Itoa(val_10)
			return widget_of_label(val_ret,val_16,val_10_tem,val_2,val_8)
		case "OCT":
			val_10,err := otherTo(val,"8")
			if err != nil{
				str := "输入有误"
				if val == ""{
					str = ""
				}
				
				return widget_of_label(str,"","","","")
			}
			val_ret := "succeed"
			val_2 := dto(val_10,"2")
			val_8 := dto(val_10,"8")
			val_16 := dto(val_10,"16")
			val_10_tem := strconv.Itoa(val_10)
			return widget_of_label(val_ret,val_16,val_10_tem,val_2,val_8)
	}

	return nil,nil,nil,nil,nil
}

//十进制转换
func dto(num int, base string)string{
	switch base{
		case "2":
			return fmt.Sprintf("%b",num)
			
		case "8":
			return fmt.Sprintf("%o",num)
		
		case "10":
			return fmt.Sprintf("%d",num)

		case "16":
			return fmt.Sprintf("%x",num)
	}
	return ""
}

//其他进制转换成十进制
func otherTo(val,base string)(int,error){
	switch base{
		case "2":
			s,err := strconv.ParseInt(val,2,64)
			if err != nil{
				return 0, err
			}
			return int(s),nil
		case "8":
			s,err := strconv.ParseInt(val,8,64)
			if err != nil{
				return 0, err
			}
			return int(s),nil
		case "16":
			s,err := strconv.ParseInt(val,16,64)
			if err != nil{
				return 0, err
			}
			return int(s),nil
	}
	return 0,nil
}


func widget_of_label(ret,hex,dec,bin,otc string)(*fyne.Container, *fyne.Container, *fyne.Container, *fyne.Container, *fyne.Container){
	retLabel := widget.NewLabel(ret)
	hexLabel := widget.NewLabel(hex) 
	decLabel := widget.NewLabel(dec)
	binLabel := widget.NewLabel(bin)
	otcLabel := widget.NewLabel(otc)
	// ret.Refresh()
	box_ret := container.NewVBox(retLabel)
	box_hex := container.NewVBox(hexLabel)
	box_bin := container.NewVBox(binLabel)
	box_otc := container.NewVBox(otcLabel)
	box_dec := container.NewVBox(decLabel)
	box_ret.Move(fyne.NewPos(160,150))
	box_bin.Move(fyne.NewPos(160,300))
	box_hex.Move(fyne.NewPos(160,200))
	box_otc.Move(fyne.NewPos(160,250))
	box_dec.Move(fyne.NewPos(160,350))
	return box_ret, box_bin, box_dec, box_hex, box_otc
}

//清除计算结果
func clearRet(container_ret,calc_container,calc_container1,calc_container2,calc_container3,calc_container4 *fyne.Container){
	container_ret.Remove(calc_container)
	container_ret.Remove(calc_container1)
	container_ret.Remove(calc_container2)
	container_ret.Remove(calc_container3)
	container_ret.Remove(calc_container4)
}