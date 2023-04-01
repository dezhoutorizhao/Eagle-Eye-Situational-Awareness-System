package main

import (
	"bufio"
	_ "bufio"
	_ "bytes"
	_ "database/sql"
	"encoding/json"
	"fmt"
	_ "io"
	"os"
	"os/exec"
	"path/filepath"
	_ "time"
)

func main() {
	//rtspUrl := "rtsp://localhost/stream"
	rtspUrl := "rtsp://127.0.0.1:554/stream"
	conf := "0.3"
	interval := "8"
	location := "A"
	task := "fire"
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前工作目录出错：", err)
		return
	}
	pythonFile := filepath.Join(dir, "./Files/Fire/dtest.py")
	//cmd := exec.Command("python",pythonFile, "--src" ,rtspUrl, "--conf",conf, "--interval" ,interval, "--location" ,location, "--task" ,task)

	//inStmt := fmt.Sprintf("python",pythonFile,"--src" ,rtspUrl, "-	-conf",conf, "--interval" ,interval, "--location" ,location, "--task" ,task)

	cmd := exec.Command("python", pythonFile, "--src", rtspUrl, "--conf", conf, "--interval", interval, "--location", location, "--task", task, "--webcam")
	fmt.Println(cmd)

	//stdoutFile, err := os.Create("stdout.log")
	//if err != nil {
	//	panic(err)
	//}
	//defer stdoutFile.Close()
	//cmd.Stdout = stdoutFile
	//if err := cmd.Start(); err != nil {
	//	panic(err)
	//}
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
		text := scanner.Text()
		fmt.Println("脚本输出：", text)
		// 读取新增内容并处理
		//tailCmd := exec.Command("tail", "-f", "stdout.log")
		//tailCmd := exec.Command("more", "+1", "stdout.log")
		//tailReader, err := tailCmd.StdoutPipe()
		//if err != nil {
		//	panic(err)
		//}
		//defer tailReader.Close()
		//if err := tailCmd.Start(); err != nil {
		//	panic(err)
		//}
		//scanner := bufio.NewScanner(tailReader)
		//for scanner.Scan() {
		//	text := scanner.Text()
		//	fmt.Println("脚本输出：", text)
		//}
		//if err := scanner.Err(); err != nil {
		//	if err == io.EOF {
		//		return
		//	}
		//	panic(err)
		//}

		//fmt.Println(out)
		//println(string(out[1]))
		//
		var algorithm_result map[string]interface{}

		//解析并添加到数据库中
		trans_err := json.Unmarshal([]byte(text), &algorithm_result)
		if trans_err != nil {
			fmt.Println(trans_err)
		}

		insert_photo := algorithm_result["Photo"].(string)
		fmt.Println(insert_photo)
		//temp_photo := algorithm_result["Photo"]
		//temp_location := algorithm_result["Location"]
		//temp_rate := algorithm_result["Rate"]
		//temp_task := algorithm_result["Task"]

		//println(temp_photo)
	}
}
