package detect_result

import (
	"encoding/json"
	"fmt"
	"github.com/Jeffail/gabs"
	"github.com/esap/wechat"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

var AccessToken string
var Input_Accesstoken string
var NextOpenId string
var Input_NextOpenId []string

type Data struct {
	Openid []string `json:"openid"`
}

type Weixin_json struct {
	Total      int    `json:"total"`
	Count      int    `json:"count"`
	Data       Data   `json:"data"`
	NextOpenid string `json:"next_openid"`
}

type Weixin_Results struct {
	Id       int    `json:"id"`
	Photo    string `json:"photo"`
	Rate     string `json:"rate"`
	Task     string `json:"task"`
	Location string `json:"location"`
}

func To_weixin(to_weixin *gin.Context) {
	Get_token()
	Get_openid()
	body, _ := to_weixin.GetRawData()
	var result Weixin_Results
	err := json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println(err.Error())
	}

	cfg := &wechat.WxConfig{
		AppId: "wx0942daad1454b2fe",
		//AppId:  NextOpenId,
		Secret: "d1be5a7ac246c706a389dbf45656ea2c",
	}
	print("这是主函数中调用的结果 ", Input_NextOpenId)
	temp_rate := fmt.Sprintf("%s", result.Rate)
	temp_location := fmt.Sprintf("%s", result.Location)
	temp_task := fmt.Sprintf("%s", result.Task)

	wechat_msg := "警告!!!\n" + "地点:" + temp_location + "\n" + "发生了" + temp_task + "的紧急事件\n" + "检测概率为:" + temp_rate

	wechat.New(cfg).SendText(NextOpenId, wechat_msg)

	for _, value := range Input_NextOpenId {
		wechat.New(cfg).SendText(value, wechat_msg)
		println("已发送")
	}

	Input_NextOpenId = []string{}
}

// 获取token
func Get_token() {
	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wx0942daad1454b2fe&secret=d1be5a7ac246c706a389dbf45656ea2c"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	println(body)
	// 解析JSON
	jsonParsed, err := gabs.ParseJSON(body)
	if err != nil {
		fmt.Println(err)
		return
	}
	AccessToken, _ := jsonParsed.Path("access_token").Data().(string)
	println(AccessToken + " this is accesstoken,token获取成功")
	Input_Accesstoken = AccessToken
	Get_openid()
}

// 获取用户openid
//func Get_openid() {
//	url := "https://api.weixin.qq.com/cgi-bin/user/get?access_token=" + Input_Accesstoken + "&next_openid=" + "os1tV6JjmH5D1fAADh8GF7j5FRgs"
//	//println("这是get_openid中的accesstoken: ", Input_Accesstoken)
//	//url := "https://api.weixin.qq.com/cgi-bin/user/get?access_token=" + Input_Accesstoken
//	resp, err := http.Get(url)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer resp.Body.Close()
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(string(body))
//	// 解析JSON
//	jsonParsed, err := gabs.ParseJSON(body)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	NextOpenId, _ := jsonParsed.Path("next_openid").Data().(string)
//	Input_NextOpenId = NextOpenId
//
//	print(NextOpenId + " this is next next_openid,openid获取成功")
//}

type OpenidList struct {
	TotalCount int `json:"total"`
	Count      int `json:"count"`
	Data       struct {
		Openid []string `json:"openid"`
	} `json:"data"`
	NextOpenid string `json:"next_openid"`
}

func Get_openid() {
	//url := "https://api.weixin.qq.com/cgi-bin/user/get?access_token=" + Input_Accesstoken + "&next_openid=" + "os1tV6JjmH5D1fAADh8GF7j5FRgs"
	url := "https://api.weixin.qq.com/cgi-bin/user/get?access_token=" + Input_Accesstoken
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body), "this is body")
	// 解析JSON
	//jsonParsed, err := gabs.ParseJSON(body)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//NextOpenId, ok := jsonParsed.Path("next_openid").Data().(string)
	//if !ok || NextOpenId == "" {
	//	fmt.Println("已经获取完所有用户的openid")
	//	return
	//}
	////Input_NextOpenId = NextOpenId
	//var weixin_openid Weixin_json
	//err = json.Unmarshal(body, &weixin_openid)
	//for _, openid := range weixin_openid.Data.Openid {
	//	Input_NextOpenId = append(Input_NextOpenId, openid)
	//}

	var openidList OpenidList
	err = json.Unmarshal(body, &openidList)
	if err != nil {
		fmt.Println(err)
	}
	for _, openid := range openidList.Data.Openid {
		fmt.Println(openid)
		Input_NextOpenId = append(Input_NextOpenId, openid)
	}

	fmt.Println(NextOpenId + " this is next next_openid,openid获取成功")

	fmt.Println(NextOpenId + " this is next next_openid,openid获取成功")
}
