package detect_result

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

var text string

func Run_python_fire(data string, vid_stride int, threshold float32, where_loc string, fire *gin.Context) bool {

	//获取工作目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前工作目录出错：", err)
		return false
	}
	pythonFile := filepath.Join(dir, "./../Files/Fire", "/dtest.py")
	fmt.Println(pythonFile)

	//将参数转为指定格式
	temp_vid_stride := strconv.Itoa(vid_stride)
	temp_threshold := fmt.Sprintf("%f", threshold)

	//传参调用，并返回返回值
	cmd := exec.Command("python", pythonFile, "--src", data, "--conf", temp_threshold, "--interval", temp_vid_stride, "--location", where_loc, "--task", "fire", "--webcam")
	fmt.Println(cmd)
	fmt.Println("这是data", data)

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	defer stdoutPipe.Close()
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(stdoutPipe)
	for scanner.Scan() {
		text = scanner.Text()
		fmt.Println("脚本输出：", text)

		//这里放的转换
		var algorithm_result map[string]interface{}

		Push_to_front(fire)

		//解析并添加到数据库中
		trans_err := json.Unmarshal([]byte(text), &algorithm_result)
		if trans_err != nil {
			fmt.Println(trans_err)
		}

		insert_photo := algorithm_result["Photo"].(string)
		rate := algorithm_result["Rate"].(string)
		//insert_task := algorithm_result["Task"].(string)
		insert_location := algorithm_result["Location"].(string)

		var temp_Review int8
		temp_Review = 0
		this_time := getDatetime()
		insert_rate, _ := strconv.ParseFloat(rate, 64)

		sqlStr := "insert into detection.results(photo,rate,task,location,time,Review) values(?,?,?,?,?,?)"

		fmt.Println(sqlStr)
		inStmt, err := Db_sql.Prepare(sqlStr)
		if err != nil {
			fmt.Println("预编译出现异常", err)
		}
		fmt.Println(inStmt)
		_, err2 := inStmt.Exec(insert_photo, insert_rate, "火灾", insert_location, this_time, temp_Review)
		if err2 != nil {
			fmt.Println("执行出现异常", err2)
		}

		if err := scanner.Err(); err != nil {
			panic(err)
		}

	}
	return true
}

func Run_python_smoke(data string, vid_stride int, threshold float32, where_loc string, fire *gin.Context) bool {
	//获取工作目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前工作目录出错：", err)
		return false
	}
	pythonFile := filepath.Join(dir, "./../Files/Smoking", "/dtest.py")
	fmt.Println(pythonFile)

	//将参数转为指定格式
	temp_vid_stride := strconv.Itoa(vid_stride)
	temp_threshold := fmt.Sprintf("%f", threshold)

	//传参调用，并返回返回值
	cmd := exec.Command("python", pythonFile, "--src", data, "--conf", temp_threshold, "--interval", temp_vid_stride, "--location", where_loc, "--task", "smoke", "--webcam")
	fmt.Println(cmd)

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	defer stdoutPipe.Close()
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(stdoutPipe)
	for scanner.Scan() {
		text = scanner.Text()
		fmt.Println("脚本输出：", text)

		//这里放的转换
		var algorithm_result map[string]interface{}

		Push_to_front(fire)

		//解析并添加到数据库中
		trans_err := json.Unmarshal([]byte(text), &algorithm_result)
		if trans_err != nil {
			fmt.Println(trans_err)
		}

		insert_photo := algorithm_result["Photo"].(string)
		rate := algorithm_result["Rate"].(string)
		//insert_task := algorithm_result["Task"].(string)
		insert_location := algorithm_result["Location"].(string)

		var temp_Review int8
		temp_Review = 0
		this_time := getDatetime()
		insert_rate, _ := strconv.ParseFloat(rate, 64)

		sqlStr := "insert into detection.results(photo,rate,task,location,time,Review) values(?,?,?,?,?,?)"

		fmt.Println(sqlStr)
		inStmt, err := Db_sql.Prepare(sqlStr)
		if err != nil {
			fmt.Println("预编译出现异常", err)
		}
		fmt.Println(inStmt)
		_, err2 := inStmt.Exec(insert_photo, insert_rate, "吸烟", insert_location, this_time, temp_Review)
		if err2 != nil {
			fmt.Println("执行出现异常", err2)
		}

		if err := scanner.Err(); err != nil {
			panic(err)
		}

	}
	return true
}

func Run_python_railing(data string, vid_stride int, threshold float32, where_loc string, fire *gin.Context) bool {
	//获取工作目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前工作目录出错：", err)
		return false
	}
	pythonFile := filepath.Join(dir, "./../Files/Climb", "/dtest.py")
	fmt.Println(pythonFile)

	//将参数转为指定格式
	temp_vid_stride := strconv.Itoa(vid_stride)
	temp_threshold := fmt.Sprintf("%f", threshold)

	//传参调用，并返回返回值
	cmd := exec.Command("python", pythonFile, "--src", data, "--conf", temp_threshold, "--interval", temp_vid_stride, "--location", where_loc, "--task", "climb", "--webcam")
	fmt.Println(cmd)

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	defer stdoutPipe.Close()
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(stdoutPipe)
	for scanner.Scan() {
		text = scanner.Text()
		fmt.Println("脚本输出：", text)

		//这里放的转换
		var algorithm_result map[string]interface{}

		Push_to_front(fire)

		//解析并添加到数据库中
		trans_err := json.Unmarshal([]byte(text), &algorithm_result)
		if trans_err != nil {
			fmt.Println(trans_err)
		}

		insert_photo := algorithm_result["Photo"].(string)
		rate := algorithm_result["Rate"].(string)
		//insert_task := algorithm_result["Task"].(string)
		insert_location := algorithm_result["Location"].(string)

		var temp_Review int8
		temp_Review = 0
		this_time := getDatetime()
		insert_rate, _ := strconv.ParseFloat(rate, 64)

		sqlStr := "insert into detection.results(photo,rate,task,location,time,Review) values(?,?,?,?,?,?)"

		fmt.Println(sqlStr)
		inStmt, err := Db_sql.Prepare(sqlStr)
		if err != nil {
			fmt.Println("预编译出现异常", err)
		}
		fmt.Println(inStmt)
		_, err2 := inStmt.Exec(insert_photo, insert_rate, "栏杆", insert_location, this_time, temp_Review)
		if err2 != nil {
			fmt.Println("执行出现异常", err2)
		}

		if err := scanner.Err(); err != nil {
			panic(err)
		}

	}
	return true
}

func Run_python_wave(data string, vid_stride int, threshold float32, where_loc string, fire *gin.Context) bool {
	//获取工作目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前工作目录出错：", err)
		return false
	}
	pythonFile := filepath.Join(dir, "./../Files/Wave", "/dtest.py")
	fmt.Println(pythonFile)

	//将参数转为指定格式
	temp_vid_stride := strconv.Itoa(vid_stride)
	temp_threshold := fmt.Sprintf("%f", threshold)

	//传参调用，并返回返回值
	cmd := exec.Command("python", pythonFile, "--src", data, "--conf", temp_threshold, "--interval", temp_vid_stride, "--location", where_loc, "--task", "wave", "--webcam")
	fmt.Println(cmd)

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	defer stdoutPipe.Close()
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(stdoutPipe)
	for scanner.Scan() {
		text = scanner.Text()
		fmt.Println("脚本输出：", text)

		//这里放的转换
		var algorithm_result map[string]interface{}

		Push_to_front(fire)

		//解析并添加到数据库中
		trans_err := json.Unmarshal([]byte(text), &algorithm_result)
		if trans_err != nil {
			fmt.Println(trans_err)
		}

		insert_photo := algorithm_result["Photo"].(string)
		rate := algorithm_result["Rate"].(string)
		//insert_task := algorithm_result["Task"].(string)
		insert_location := algorithm_result["Location"].(string)

		var temp_Review int8
		temp_Review = 0
		this_time := getDatetime()
		insert_rate, _ := strconv.ParseFloat(rate, 64)

		sqlStr := "insert into detection.results(photo,rate,task,location,time,Review) values(?,?,?,?,?,?)"

		fmt.Println(sqlStr)
		inStmt, err := Db_sql.Prepare(sqlStr)
		if err != nil {
			fmt.Println("预编译出现异常", err)
		}
		fmt.Println(inStmt)
		_, err2 := inStmt.Exec(insert_photo, insert_rate, "挥手", insert_location, this_time, temp_Review)
		if err2 != nil {
			fmt.Println("执行出现异常", err2)
		}

		if err := scanner.Err(); err != nil {
			panic(err)
		}

	}
	return true
}

func Run_python_drown(data string, vid_stride int, threshold float32, where_loc string, fire *gin.Context) bool {
	//获取工作目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前工作目录出错：", err)
		return false
	}
	pythonFile := filepath.Join(dir, "./../Files/Drawn", "/dtest.py")
	fmt.Println(pythonFile)

	//将参数转为指定格式
	temp_vid_stride := strconv.Itoa(vid_stride)
	temp_threshold := fmt.Sprintf("%f", threshold)

	//传参调用，并返回返回值
	cmd := exec.Command("python", pythonFile, "--src", data, "--conf", temp_threshold, "--interval", temp_vid_stride, "--location", where_loc, "--task", "drown", "--webcam")
	fmt.Println(cmd)

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	defer stdoutPipe.Close()
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(stdoutPipe)
	for scanner.Scan() {
		text = scanner.Text()
		fmt.Println("脚本输出：", text)

		//这里放的转换
		var algorithm_result map[string]interface{}

		Push_to_front(fire)

		//解析并添加到数据库中
		trans_err := json.Unmarshal([]byte(text), &algorithm_result)
		if trans_err != nil {
			fmt.Println(trans_err)
		}

		insert_photo := algorithm_result["Photo"].(string)
		rate := algorithm_result["Rate"].(string)
		//insert_task := algorithm_result["Task"].(string)
		insert_location := algorithm_result["Location"].(string)

		var temp_Review int8
		temp_Review = 0
		this_time := getDatetime()
		insert_rate, _ := strconv.ParseFloat(rate, 64)

		sqlStr := "insert into detection.results(photo,rate,task,location,time,Review) values(?,?,?,?,?,?)"

		fmt.Println(sqlStr)
		inStmt, err := Db_sql.Prepare(sqlStr)
		if err != nil {
			fmt.Println("预编译出现异常", err)
		}
		fmt.Println(inStmt)
		_, err2 := inStmt.Exec(insert_photo, insert_rate, "溺水", insert_location, this_time, temp_Review)
		if err2 != nil {
			fmt.Println("执行出现异常", err2)
		}

		if err := scanner.Err(); err != nil {
			panic(err)
		}

	}
	return true
}

func Run_python_fall(data string, vid_stride int, threshold float32, where_loc string, fire *gin.Context) bool {
	//获取工作目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前工作目录出错：", err)
		return false
	}
	pythonFile := filepath.Join(dir, "./../Files/Fall", "/dtest.py")
	fmt.Println(pythonFile)

	//将参数转为指定格式
	temp_vid_stride := strconv.Itoa(vid_stride)
	temp_threshold := fmt.Sprintf("%f", threshold)

	//传参调用，并返回返回值
	cmd := exec.Command("python", pythonFile, "--src", data, "--conf", temp_threshold, "--interval", temp_vid_stride, "--location", where_loc, "--task", "fall", "--webcam")
	fmt.Println(cmd)

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	defer stdoutPipe.Close()
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(stdoutPipe)
	for scanner.Scan() {
		text = scanner.Text()
		fmt.Println("脚本输出：", text)

		//这里放的转换
		var algorithm_result map[string]interface{}

		Push_to_front(fire)

		//解析并添加到数据库中
		trans_err := json.Unmarshal([]byte(text), &algorithm_result)
		if trans_err != nil {
			fmt.Println(trans_err)
		}

		insert_photo := algorithm_result["Photo"].(string)
		rate := algorithm_result["Rate"].(string)
		//insert_task := algorithm_result["Task"].(string)
		insert_location := algorithm_result["Location"].(string)

		var temp_Review int8
		temp_Review = 0
		this_time := getDatetime()
		insert_rate, _ := strconv.ParseFloat(rate, 64)

		sqlStr := "insert into detection.results(photo,rate,task,location,time,Review) values(?,?,?,?,?,?)"

		fmt.Println(sqlStr)
		inStmt, err := Db_sql.Prepare(sqlStr)
		if err != nil {
			fmt.Println("预编译出现异常", err)
		}
		fmt.Println(inStmt)
		_, err2 := inStmt.Exec(insert_photo, insert_rate, "摔倒", insert_location, this_time, temp_Review)
		if err2 != nil {
			fmt.Println("执行出现异常", err2)
		}

		if err := scanner.Err(); err != nil {
			panic(err)
		}

	}
	return true
}

func Run_python_water(data string, vid_stride int, threshold float32, where_loc string, fire *gin.Context) bool {
	//获取工作目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前工作目录出错：", err)
		return false
	}
	pythonFile := filepath.Join(dir, "./../Files/Water", "/dtest.py")
	fmt.Println(pythonFile)

	//将参数转为指定格式
	temp_vid_stride := strconv.Itoa(vid_stride)
	temp_threshold := fmt.Sprintf("%f", threshold)

	//传参调用，并返回返回值
	cmd := exec.Command("python", pythonFile, "--src", data, "--conf", temp_threshold, "--interval", temp_vid_stride, "--location", where_loc, "--task", "water", "--webcam")
	fmt.Println(cmd)

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	defer stdoutPipe.Close()
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(stdoutPipe)
	for scanner.Scan() {
		text = scanner.Text()
		fmt.Println("脚本输出：", text)

		//这里放的转换
		var algorithm_result map[string]interface{}

		Push_to_front(fire)

		//解析并添加到数据库中
		trans_err := json.Unmarshal([]byte(text), &algorithm_result)
		if trans_err != nil {
			fmt.Println(trans_err)
		}

		insert_photo := algorithm_result["Photo"].(string)
		rate := algorithm_result["Rate"].(string)
		//insert_task := algorithm_result["Task"].(string)
		insert_location := algorithm_result["Location"].(string)

		var temp_Review int8
		temp_Review = 0
		this_time := getDatetime()
		insert_rate, _ := strconv.ParseFloat(rate, 64)

		sqlStr := "insert into detection.results(photo,rate,task,location,time,Review) values(?,?,?,?,?,?)"

		fmt.Println(sqlStr)
		inStmt, err := Db_sql.Prepare(sqlStr)
		if err != nil {
			fmt.Println("预编译出现异常", err)
		}
		fmt.Println(inStmt)
		_, err2 := inStmt.Exec(insert_photo, insert_rate, "积水", insert_location, this_time, temp_Review)
		if err2 != nil {
			fmt.Println("执行出现异常", err2)
		}

		if err := scanner.Err(); err != nil {
			panic(err)
		}

	}
	return true
}

func Push_to_front(push_to_front *gin.Context) {
	//length := len(text)
	//// 设置Content-Length头
	//header := push_to_front.Writer.Header()
	//header.Set("Content-Length", strconv.Itoa(length))
	//go func() {
	//	defer func() {
	//		// 在协程结束时销毁text变量
	//		text = ""
	//	}()
	push_to_front.String(200, text)
	//}()
}
