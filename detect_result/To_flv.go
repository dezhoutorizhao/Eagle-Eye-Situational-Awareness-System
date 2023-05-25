package detect_result

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

type Rtsp_Flv struct {
	Rtsp_location []string `json:"rtsp_location"`
}

var rtsp_location0 string
var rtsp_location1 string
var rtsp_location2 string
var rtsp_location3 string
var rtsp_location4 string
var rtsp_location5 string
var rtsp_location6 string
var rtsp_location7 string
var rtsp_Flv Rtsp_Flv

func Total_to_flv(total_rtsp_flv *gin.Context) {
	body, _ := total_rtsp_flv.GetRawData()

	err := json.Unmarshal(body, &rtsp_Flv)
	if err != nil {
		fmt.Println(err.Error())
	}

	rtsp_location0 = rtsp_Flv.Rtsp_location[0]
	rtsp_location1 = rtsp_Flv.Rtsp_location[1]
	rtsp_location2 = rtsp_Flv.Rtsp_location[2]
	rtsp_location3 = rtsp_Flv.Rtsp_location[3]
	rtsp_location4 = rtsp_Flv.Rtsp_location[4]
	rtsp_location5 = rtsp_Flv.Rtsp_location[5]
	rtsp_location6 = rtsp_Flv.Rtsp_location[6]
	rtsp_location7 = rtsp_Flv.Rtsp_location[7]

	//go To_flv0(rtsp_Flv.Rtsp_location[0], total_rtsp_flv)
	//go To_flv1(rtsp_Flv.Rtsp_location[1], total_rtsp_flv)
	//go To_flv2(rtsp_Flv.Rtsp_location[2], total_rtsp_flv)
	//go To_flv3(rtsp_Flv.Rtsp_location[3], total_rtsp_flv)
	//go To_flv4(rtsp_Flv.Rtsp_location[4], total_rtsp_flv)
	//go To_flv5(rtsp_Flv.Rtsp_location[5], total_rtsp_flv)

}

func Show_to_front(total_rtsp_flv *gin.Context) {
	total_rtsp_flv.JSON(200, rtsp_Flv)
}

func To_flv0(rtsp_flv *gin.Context) {

	cmd := exec.Command("gst-launch-1.0", "-e", "rtspsrc", "location="+rtsp_location0, "!", "rtph264depay", "!", "flvmux", "!", "filesink", "location=../flv/"+"output0.flv")

	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

	// 返回给前端阶段
	f, err := os.Open("../flv/output0.flv")
	if err != nil {
		rtsp_flv.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer f.Close()
	//设置HTTP响应头为flv格式
	rtsp_flv.Header("Content-Type", "video/x-flv")
	io.Copy(rtsp_flv.Writer, f)

}

//使用Gin框架，我们可以通过gin.Default()函数创建一个默认配置的Gin引擎，然后使用r.GET()函数来处理HTTP GET请求。在处理函数中，我们使用os.Open()函数打开生成的FLV文件，然后使用gin.Context的Writer成员将FLV文件的内容直接写入HTTP响应中，最后使用c.Header()函数设置HTTP响应头为FLV格式。最后使用r.Run(":8080")函数启动HTTP服务器，监听端口号为8080。
//用户可以通过浏览器访问http://localhost:8080/来获取生成的FLV文件。

//这段代码使用了Go语言的os/exec包来执行一个外部命令，命令为gst-launch-1.0，用于从一个RTSP流中提取H.264视频数据并将其保存为FLV格式的文件。
//具体来说，代码中使用了exec.Command函数创建了一个*exec.Cmd类型的对象cmd，该对象表示要执行的命令及其参数。其中，命令为gst-launch-1.0，参数为-e、rtspsrc、location=rtsp://your-rtsp-link、!、rtph264depay、!、flvmux、!、filesink、location=output.flv，这些参数指定了从指定的RTSP流中提取H.264视频数据，并将其经由FLV格式进行封装，最终保存为名为output.flv的文件。
//然后，代码中使用了cmd.Output()函数来运行cmd对象代表的命令，并返回命令的输出结果。这里的输出结果是命令运行结果的标准输出，也就是直接输出到控制台的内容。如果命令在执行过程中出现错误，cmd.Output()函数会返回一个error类型的值，表示错误信息。在这里，代码使用了if语句来判断err是否为nil，如果不为nil，则表示命令执行失败，将错误信息打印输出。
//最后，代码使用fmt.Println函数将命令的输出结果打印到控制台。由于cmd.Output()函数返回的结果是一个[]byte类型的字节数组，因此需要使用string函数将其转换为一个字符串。

//上述示例代码中，使用http.HandleFunc函数创建了一个HTTP请求处理函数，该函数会打开生成的FLV文件output.flv，设置HTTP响应头为FLV格式，然后将FLV文件作为响应内容返回给客户端。
//最后，使用http.ListenAndServe函数启动HTTP服务器，监听端口号为8080。启动后，用户可以通过浏览器访问http://localhost:8080/来获取生成的FLV文件。

func To_flv1(rtsp_flv *gin.Context) {

	//// 转码阶段
	//body, _ := rtsp_flv.GetRawData()
	//type Rtsp_Flv struct {
	//	Rtsp_location string `json:"rtsp_location"`
	//}
	//var Rtsp_location Rtsp_Flv
	//err := json.Unmarshal(body, &Rtsp_location)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//
	//cmd := exec.Command("gst-launch-1.0", "-e", "rtspsrc", "location="+Rtsp_location.Rtsp_location, "!", "rtph264depay", "!", "flvmux", "!", "filesink", "location=,"+"output2.flv")
	//
	//out, err := cmd.Output()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(string(out))
	//
	//// 返回给前端阶段
	//f, err := os.Open("output2.flv")
	//if err != nil {
	//	rtsp_flv.String(http.StatusInternalServerError, err.Error())
	//	return
	//}
	//defer f.Close()
	////设置HTTP响应头为flv格式
	//rtsp_flv.Header("Content-Type", "video/x-flv")
	//io.Copy(rtsp_flv.Writer, f)

	cmd := exec.Command("gst-launch-1.0", "-e", "rtspsrc", "location="+rtsp_location1, "!", "rtph264depay", "!", "flvmux", "!", "filesink", "location=../flv/"+"output1.flv")

	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

	// 返回给前端阶段
	f, err := os.Open("../flv/output1.flv")
	if err != nil {
		rtsp_flv.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer f.Close()
	//设置HTTP响应头为flv格式
	rtsp_flv.Header("Content-Type", "video/x-flv")
	io.Copy(rtsp_flv.Writer, f)

}

func To_flv2(rtsp_flv *gin.Context) {

	cmd := exec.Command("gst-launch-1.0", "-e", "rtspsrc", "location="+rtsp_location2, "!", "rtph264depay", "!", "flvmux", "!", "filesink", "location=../flv/"+"output2.flv")

	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

	// 返回给前端阶段
	f, err := os.Open("../flv/output2.flv")
	if err != nil {
		rtsp_flv.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer f.Close()
	//设置HTTP响应头为flv格式
	rtsp_flv.Header("Content-Type", "video/x-flv")
	io.Copy(rtsp_flv.Writer, f)

}

func To_flv3(rtsp_flv *gin.Context) {

	cmd := exec.Command("gst-launch-1.0", "-e", "rtspsrc", "location="+rtsp_location3, "!", "rtph264depay", "!", "flvmux", "!", "filesink", "location=../flv/"+"output3.flv")

	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

	// 返回给前端阶段
	f, err := os.Open("../flv/output3.flv")
	if err != nil {
		rtsp_flv.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer f.Close()
	//设置HTTP响应头为flv格式
	rtsp_flv.Header("Content-Type", "video/x-flv")
	io.Copy(rtsp_flv.Writer, f)

}

func To_flv4(rtsp_flv *gin.Context) {

	cmd := exec.Command("gst-launch-1.0", "-e", "rtspsrc", "location="+rtsp_location4, "!", "rtph264depay", "!", "flvmux", "!", "filesink", "location=../flv/"+"output4.flv")

	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

	// 返回给前端阶段
	f, err := os.Open("../flv/output4.flv")
	if err != nil {
		rtsp_flv.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer f.Close()
	//设置HTTP响应头为flv格式
	rtsp_flv.Header("Content-Type", "video/x-flv")
	io.Copy(rtsp_flv.Writer, f)

}

func To_flv5(rtsp_flv *gin.Context) {

	cmd := exec.Command("gst-launch-1.0", "-e", "rtspsrc", "location="+rtsp_location5, "!", "rtph264depay", "!", "flvmux", "!", "filesink", "location=../flv/"+"output5.flv")

	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

	// 返回给前端阶段
	dir, _ := os.Getwd()
	fmt.Println(filepath.Join(dir, "./"))
	f, err := os.Open("../flv/output5.flv")
	if err != nil {
		rtsp_flv.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer f.Close()
	//设置HTTP响应头为flv格式
	rtsp_flv.Header("Content-Type", "video/x-flv")
	io.Copy(rtsp_flv.Writer, f)

}
func To_flv6(rtsp_flv *gin.Context) {

	cmd := exec.Command("gst-launch-1.0", "-e", "rtspsrc", "location="+rtsp_location6, "!", "rtph264depay", "!", "flvmux", "!", "filesink", "location=../flv/"+"output6.flv")

	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

	// 返回给前端阶段
	f, err := os.Open("../flv/output6.flv")
	if err != nil {
		rtsp_flv.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer f.Close()
	//设置HTTP响应头为flv格式
	rtsp_flv.Header("Content-Type", "video/x-flv")
	io.Copy(rtsp_flv.Writer, f)

}
func To_flv7(rtsp_flv *gin.Context) {

	cmd := exec.Command("gst-launch-1.0", "-e", "rtspsrc", "location="+rtsp_location7, "!", "rtph264depay", "!", "flvmux", "!", "filesink", "location=../flv/"+"output7.flv")

	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

	// 返回给前端阶段
	f, err := os.Open("../flv/output7.flv")
	if err != nil {
		rtsp_flv.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer f.Close()
	//设置HTTP响应头为flv格式
	rtsp_flv.Header("Content-Type", "video/x-flv")
	io.Copy(rtsp_flv.Writer, f)

}

func test1() {
	for i := 0; i < 10; i++ {
		fmt.Println(rtsp_location1, "这是test1")
	}

}

func test2() {
	for i := 0; i < 10; i++ {
		fmt.Println(rtsp_location2, "这是test2")
	}

}
