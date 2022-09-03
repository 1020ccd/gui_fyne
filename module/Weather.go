package module

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"encoding/json"
)

type weather_reltime struct{
	Temperature string `json:"temperature"`
	Humidity string  `json:"humidity"`
	Info string  `json:"info"`
	Wid string  `json:"wid"`
	Direct string `json:"direct"`
	Power string `json:"power"`
	Aqi string `json:"aqi"`
}


type weather_feture_info struct{
	Date  string `json:"date"`
    Temperature string `json:"temperature"`
    Weather string `json:"weather"`
    Wid struct{
            Day string `json:"day"`
            Night string `json:"night"`
        } `json:"wid"`
    Direct string `json:"direct"`
}


type weather_result struct{
	City string  `json:"city"`
	Realtime weather_reltime  `json:"realtime"`
	Future []weather_feture_info  `json:"future"`
	
}

type weather_info struct{
	Error_code int `json:"error_code"`
	Reason string `json:"reason"`
	Result weather_result  `json:"result"`

}

var query_result *canvas.Text
var head_text *canvas.Text
var weather_field fyne.CanvasObject
var grid *fyne.Container
var btn *fyne.Container

var info weather_info
// 主操作
func WeatherQuery(a fyne.App) *fyne.Container {

	//创建输入框
    box := input_create()

	btn = btn_create(a)
	
    grid = container.NewGridWithRows(3,box,)
	
    return grid
}

//创建输入框及确定按钮
func input_create() *fyne.Container {
    input := widget.NewEntry()
    input.SetPlaceHolder("城市名称如：苏州")

	//点击确定，请求数据并展示天气信息
	confirm_button := widget.NewButton("确定", func() {
		data, _ := get_weather_info(input.Text)
		
		
		err := json.Unmarshal([]byte(data), &info)
		if err != nil{
			fmt.Println(err)
		}
		grid.Remove(weather_field)
		grid.Remove(btn)
		green := color.NRGBA{R:0,G:255,B:0,A:255}
		red := color.NRGBA{R:255,G:0,B:0,A:255}
		space := layout.NewSpacer()
		
		//城市
		cityStr := canvas.NewText("城市", green)
		city := canvas.NewText(info.Result.City, green)

		//温度
		temperatureStr := canvas.NewText("温度", green)
		temperature := canvas.NewText(info.Result.Realtime.Temperature, green)

		//湿度
		humidityStr := canvas.NewText("湿度", green)
		humidity := canvas.NewText(info.Result.Realtime.Humidity, green)

		//天气信息
		infoStr := canvas.NewText("天气信息", green)
		infomation := canvas.NewText(info.Result.Realtime.Info, green)

		//风向
		directStr := canvas.NewText("风向", green)
		direct := canvas.NewText(info.Result.Realtime.Direct, green)

		//风力
		powerStr := canvas.NewText("风力", green)
		power := canvas.NewText(info.Result.Realtime.Power, green)

		//空气质量
		aqiStr := canvas.NewText("空气质量", green)
		aqi := canvas.NewText(info.Result.Realtime.Aqi, green)


		if info.Error_code != 0{
			query_result = canvas.NewText(info.Reason,red)
			query_result.TextSize = 10
			head_text = canvas.NewText("查询失败：", red)
			head_text.TextSize = 10
			weather_field = field_weather(
				space, head_text, space,
				space,query_result,
			)
			grid.Add(weather_field)
			grid.Refresh()
		}else{
			query_result = canvas.NewText(info.Reason,green)
			head_text = canvas.NewText("查询结果：", green)
			weather_field = field_weather(
				head_text,space, query_result,
				cityStr, space, city,
				temperatureStr, space, temperature,
				humidityStr, space, humidity,
				infoStr, space, infomation,
				directStr, space, direct,
				powerStr, space, power,
				aqiStr, space, aqi,
			)
			grid.Add(weather_field)
			grid.Add(btn)
			grid.Refresh()
		}
		
	})
	
	btn_box := container.NewVBox(confirm_button)
    box := container.NewVBox(input)

    box1 := container.NewGridWithRows(3,layout.NewSpacer(),box,btn_box)
    box2 := container.NewGridWithColumns(3,layout.NewSpacer(),box1,layout.NewSpacer())
    
    return box2
}



//未来五天天气 按钮
func btn_create(a fyne.App)*fyne.Container{
	
    btn := widget.NewButton("查看未来五天天气", func(){
        fmt.Println("121212")
		feture_weather_window(a)
    })
	
    box := container.NewVBox(btn)
    box1 := container.NewGridWithRows(3,layout.NewSpacer(),layout.NewSpacer(),box)
    box2 := container.NewGridWithColumns(3,layout.NewSpacer(),box1,layout.NewSpacer())
    return box2
}


// 天气信息显示区域
func field_weather(opt ...fyne.CanvasObject) *fyne.Container{
    
    info := container.NewGridWithColumns(3,)
    for _,obj := range opt{
        info.Add(obj)
        
    }
    return  container.NewCenter(info)
}

//创建子窗口---未来五天天气
func feture_weather_window(a fyne.App){
	
    feature_window := a.NewWindow("未来五天天气")
    feature_window.Resize(fyne.NewSize(500,500))
	centent := container.NewGridWithRows(5,)
	for _,i := range info.Result.Future{
		s := get_future(i)
		center_centent := feature_window_box(s)
		centent.Add(center_centent)
	}
	
	feature_window.SetContent(centent)
    feature_window.Show()
    
}

func feature_window_box(text []*canvas.Text)*fyne.Container{
	var pos float32 = -15
	box_without_layout := container.NewWithoutLayout()
	for _,i := range text{
		
		i.Move(fyne.NewPos(0,pos))
		i.Alignment = fyne.TextAlignLeading
		
		box_without_layout.Add(i)
		pos += 15
	}
	
	
	box_Center := container.NewCenter(box_without_layout)
	return box_Center
}


func get_future(get_info weather_feture_info)([]*canvas.Text){
	var s1 = make([]*canvas.Text,0)
	green := color.NRGBA{R:0,G:255,B:0,A:255}

	dateStr := fmt.Sprintf("日期：  %s       ", get_info.Date)
	date := canvas.NewText(dateStr, green)

	temperatureStr := fmt.Sprintf("温度：  %s", get_info.Temperature)
	temperature := canvas.NewText(temperatureStr, green)

	weatherStr := fmt.Sprintf("天气信息：  %s", get_info.Weather) 
	weather := canvas.NewText(weatherStr, green)

	directStr := fmt.Sprintf("风向：  %s", get_info.Direct) 
	direct := canvas.NewText(directStr, green)

	s1 = append(s1, date, temperature, weather, direct)

	return s1
}


//天气接口请求
func get_weather_info(city string) (string,error){
	url := fmt.Sprintf("http://apis.juhe.cn/simpleWeather/query?city=%s&key=1dd81c1a970fd9dc9a5ae549211a95e4",city)
	req, err := http.Get(url)
	if err != nil{
		log.Println(err.Error())
		return "",err
	}
	data,_ := ioutil.ReadAll(req.Body)

	return string(data), nil
}