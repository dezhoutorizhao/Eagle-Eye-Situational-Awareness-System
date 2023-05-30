package detect_result

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 要推送给mqtt服务器的string数据
var byteData string

type Integration_Struct struct {
	log_list        []Add_to_database `json:"log_list"`
	Fire_times      int               `json:"fire_times"`
	Fire_percent    float32           `json:"fire_percent"`
	Smoke_times     int               `json:"smoke_times"`
	Smoke_percent   float32           `json:"smoke_percent"`
	Railing_times   int               `json:"railing_times"`
	Railing_percent float32           `json:"'railing_percent'"`
	Wave_times      int               `json:"wave_times"`
	Wave_percent    float32           `json:"wave_percent"`
	Fall_times      int               `json:"fall_times"`
	Fall_percent    float32           `json:"fall_percent"`
	Drawn_times     int               `json:"drawn_times"`
	Drawn_percent   float32           `json:"drawn_percent"`
	Water_times     int               `json:"water_times"`
	Water_percent   float32           `json:"water_percent"`
	Maximum_type    string            `json:"maximum_type"`
	Detect_Photo    string            `json:"'detect_photo'"`
	Detect_Rate     float64           `json:"detect_rate"`
	Detect_Location string            `json:"detect_location"`
	Detect_type     string            `json:"detect_type"`
	Detect_time     sql.NullString    `json:"detect_time"`
	Location_times  map[string]int    `json:"location_times"`
}

func Integration_Function(c *gin.Context) {
	Result_to_front(c)
	var integration_value Integration_Struct
	integration_value.log_list = Result_to_front_logs
	integration_value.Fire_times = Fire_times
	integration_value.Fire_percent = Fire_percent
	integration_value.Smoke_times = Smoke_times
	integration_value.Smoke_percent = Smoke_percent
	integration_value.Railing_times = Railing_times
	integration_value.Railing_percent = Railing_percent
	integration_value.Wave_times = Wave_times
	integration_value.Wave_percent = Wave_percent
	integration_value.Fall_times = Fall_times
	integration_value.Fall_percent = Fall_percent
	integration_value.Drawn_times = Drawn_times
	integration_value.Drawn_percent = Drawn_percent
	integration_value.Water_times = Water_times
	integration_value.Water_percent = Water_percent
	integration_value.Maximum_type = Maximum_type
	integration_value.Detect_Photo = Detect_Photo
	integration_value.Detect_Rate = Detect_Rate
	integration_value.Detect_Location = Detect_Location
	integration_value.Detect_type = Detect_type
	integration_value.Detect_time = Detect_time
	integration_value.Location_times = Location_times

	//if Result_to_front_logs == nil {
	//	fmt.Println("Result_to_front_logs是空的")
	//}
	//
	//if integration_value.log_list == nil {
	//	fmt.Println("integration_value.log_list是空的")
	//}
	//
	//json_list, _ := json.Marshal(integration_value.log_list)
	//if json_list == nil {
	//	fmt.Println("json_list是空的")
	//}
	//fmt.Println(json_list)

	// 日志推送mqtt的base64编码
	//fmt.Println("这是integration_value.log_list", integration_value.log_list)
	//json_log_list, _ := json.Marshal(integration_value.log_list)
	//if err != nil {
	//	fmt.Println(json_log_list)
	//}
	//json_log_list_data := base64.StdEncoding.EncodeToString(json_log_list)
	//fmt.Println("这是单个的编码：", json_log_list_data)

	jsonData, err := json.Marshal(integration_value)
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println("这是json_data", jsonData)

	// byteData变量是要传给前端的经过base64编码后的json变量
	byteData = base64.StdEncoding.EncodeToString(jsonData)
	//byteData += json_log_list_data

	fmt.Println("这是编码后的结果", byteData)
	c.String(200, byteData)
}
