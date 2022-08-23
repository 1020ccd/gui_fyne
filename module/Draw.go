package module

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)



var obj_slice []*canvas.Rectangle

var Parent fyne.Window

func Draw(parent fyne.Window) *fyne.Container{
	
	Parent = parent
	can_obj1,obj1 := box(100,100,"1")
	can_obj2,obj2 := box(190,100,"2")
	can_obj3,obj3 := box(280,100,"3")

	can_obj4,obj4 := box(100,190,"8")

	Sienna1 := color.NRGBA{R:255, G:130,B:71,A:0xff}
	obj5_can := canvas.NewRectangle(Sienna1)
	obj5_can.Resize(fyne.NewSize(80,80))
	obj5_can.Move(fyne.NewPos(190,190))
	
	obj5 := widget.NewButton("抽签", start_draw)
	obj5.Resize(fyne.NewSize(80,80))
	obj5.Move(fyne.NewPos(190,190))
	

	can_obj6,obj6 := box(280,190,"4")

	can_obj7,obj7 := box(100,280,"7")
	can_obj8,obj8 := box(190,280,"6")
	can_obj9,obj9 := box(280,280,"5")
	
	obj_slice = append(
		obj_slice, 
		can_obj1,
		can_obj2,
		can_obj3,
		can_obj6,
		can_obj9,
		can_obj8,
		can_obj7,
		can_obj4)

	draw_container := container.NewWithoutLayout(
		obj1,
		obj2,
		obj3,
		obj4,
		obj5_can,
		obj5,
		obj6,
		obj7,
		obj8,
		obj9,
		
	)
	return draw_container
}

//选项盒子
func box(width,height float32, str string)(*canvas.Rectangle, *fyne.Container){
	// red := color.NRGBA{R:0xff, A:0xff}
	obj := canvas.NewRectangle(color.White)
	obj.Resize(fyne.NewSize(80,80))
	obj.Move(fyne.NewPos(width,height))

	obj_str := canvas.NewText(str, color.Black)
	obj_str.Resize(fyne.NewSize(80,80))
	obj_str.Move(fyne.NewPos(width,height))
	obj_str.Alignment = fyne.TextAlignCenter

	obj_box := container.NewWithoutLayout(obj,obj_str)

	return obj,obj_box
}

//按钮响应操作
func start_draw(){
	gray := color.NRGBA{R:0,G:255,B:255,A:255}
	// red := color.NRGBA{R:0xff, A:0xff}
	for _,v:= range obj_slice{
		discoloration(v, color.White)
	} 
	for i:=0;i<3;i++{
		for _, v:= range obj_slice{
			discoloration(v, gray)
			discoloration(v, color.White)
		}
	}
	num := ran_num()
	for i , v:= range obj_slice{
		if i == num-1{
			discoloration(v,gray)
			break
		}
		discoloration(v, gray)
		discoloration(v, color.White)
	}
	str := fmt.Sprintf("抽中奖品为%d",num)
	time.Sleep(time.Second)
	dialog.ShowInformation("恭喜中奖", str, Parent)

}

// 生成随机数
func ran_num()int{
	min := 1
	max := 9
	nanotime := int64(time.Now().Nanosecond())
	rand.Seed(nanotime)
	num := rand.Intn(max-min)+min
	// s := strconv.Itoa(num)
	log.Println(num)
	return num
}


//变色操作
func discoloration(obj *canvas.Rectangle,color color.Color){
	
	obj.FillColor = color
	obj.Refresh()
	time.Sleep(time.Millisecond*50)
	
}


// func movto100(obj *canvas.Circle){
// 	if obj.Position() == fyne.NewPos(200,200){
// 		// log.Println("m100")
// 		canvas.NewPositionAnimation(
// 			obj.Position(),fyne.NewPos(100,100),
// 			time.Second*2,func(p fyne.Position) {
	
// 				obj.Move(p)
// 				canvas.Refresh(obj)
// 			},
// 		).Start()
// 	}
	
// }

// func movto0(obj *canvas.Circle){
// 	if obj.Position() == fyne.NewPos(100,100){
// 		// log.Println("m0")
// 		canvas.NewPositionAnimation(
// 			obj.Position(),fyne.NewPos(200,200),
// 			time.Second*2,func(p fyne.Position) {
	
// 				obj.Move(p)
// 				canvas.Refresh(obj)
// 			},
// 		).Start()
// 	}
// }

// func mov(obj *canvas.Circle){
// 	for{
// 		movto100(obj)
// 		movto0(obj)
// 	}
	
// }