// 这里主要是各个组件，火灾占比等
package detect_result

import (
	"fmt"
	"math"
)

// 总次数
var Total_times int

// 各个类型次数
var Fire_times int
var Fire_percent float32
var Smoke_times int
var Smoke_percent float32
var Railing_times int
var Railing_percent float32
var Wave_times int
var Wave_percent float32
var Fall_times int
var Fall_percent float32
var Drawn_times int
var Drawn_percent float32
var Water_times int
var Water_percent float32
var Maximum_type string

// 审核次数
var Reviewed_times int
var Waiting_Review_times int
var UnReviewed_percent float32

// 各个地点次数
var Location_times map[string]int

func Main_Starter() {
	Count_Total()
	Count_Fire()
	Count_Fall()
	Count_Drawn()
	Count_Smoke()
	Count_Water()
	Count_Wave()
	Count_Railing()
	Count_Review()
	Find_Max()
	Fire_percent = float32(Fire_times) / float32(Total_times)
	Smoke_percent = float32(Smoke_times) / float32(Total_times)
	Railing_percent = float32(Railing_times) / float32(Total_times)
	Water_percent = float32(Water_times) / float32(Total_times)
	Drawn_percent = float32(Drawn_times) / float32(Total_times)
	Fall_percent = float32(Fall_times) / float32(Total_times)
	Wave_percent = float32(Wave_times) / float32(Total_times)
	UnReviewed_percent = float32(Waiting_Review_times) / float32(Total_times)
	fmt.Println(Maximum_type, "这是最多事件类型")
}

func Count_Total() {
	sqlStatement := "SELECT COUNT(*) FROM results;"
	rows, err := Db_sql.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	var count int
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			fmt.Println(err)
		}
	}
	Total_times = count
	//fmt.Println("这是总次数", count)
}

func Count_Fire() {
	sqlStatement := "SELECT COUNT(*) FROM results WHERE task = '火灾';"
	rows, err := Db_sql.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(rows)
	defer rows.Close()

	var count int
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			fmt.Println(err)
		}
	}
	Fire_times = count
	//fmt.Println("这是火灾次数", count)
}

func Count_Smoke() {
	sqlStatement := "SELECT COUNT(*) FROM results WHERE task = '吸烟';"
	rows, err := Db_sql.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(rows)
	defer rows.Close()

	var count int
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			fmt.Println(err)
		}
	}
	Smoke_times = count
	//fmt.Println("这是吸烟次数", count)
}

func Count_Railing() {
	sqlStatement := "SELECT COUNT(*) FROM results WHERE task = '栏杆';"
	rows, err := Db_sql.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(rows)
	defer rows.Close()

	var count int
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			fmt.Println(err)
		}
	}
	Railing_times = count
	//fmt.Println("这是倚靠栏杆次数", count)
}

func Count_Water() {
	sqlStatement := "SELECT COUNT(*) FROM results WHERE task = '积水';"
	rows, err := Db_sql.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(rows)
	defer rows.Close()

	var count int
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			fmt.Println(err)
		}
	}
	Water_times = count
	//fmt.Println("这是积水次数", count)
}

func Count_Fall() {
	sqlStatement := "SELECT COUNT(*) FROM results WHERE task = '摔倒';"
	rows, err := Db_sql.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(rows)
	defer rows.Close()

	var count int
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			fmt.Println(err)
		}
	}
	Fall_times = count
	//fmt.Println("这是摔倒次数", count)
}

func Count_Drawn() {
	sqlStatement := "SELECT COUNT(*) FROM results WHERE task = '溺水';"
	rows, err := Db_sql.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(rows)
	defer rows.Close()

	var count int
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			fmt.Println(err)
		}
	}
	Drawn_times = count
	//fmt.Println("这是倚靠栏杆次数", count)
}

func Count_Wave() {
	sqlStatement := "SELECT COUNT(*) FROM results WHERE task = '挥手';"
	rows, err := Db_sql.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(rows)
	defer rows.Close()

	var count int
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			fmt.Println(err)
		}
	}
	Wave_times = count
	//fmt.Println("这是倚靠栏杆次数", count)
}

// 计算审核次数
func Count_Review() {
	sqlStatement := "SELECT COUNT(*) FROM results WHERE Review = '1';"
	rows, err := Db_sql.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(rows)
	defer rows.Close()

	var count int
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			fmt.Println(err)
		}
	}
	Reviewed_times = count
	Waiting_Review_times = Total_times - count
	//fmt.Println("这是审核次数次数", count)
}

func Count_Location() {
	sqlStatement := "SELECT location,COUNT(*) as count FROM results GROUP BY location;"
	// 执行查询
	rows, err := Db_sql.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	Location_times = make(map[string]int)
	// 遍历结果集，解析查询结果
	for rows.Next() {
		var location string
		var count int
		if err := rows.Scan(&location, &count); err != nil {
			panic(err)
		}
		Location_times[location] += count
		//fmt.Printf("%s 发生了 %d 次\n", location, Location_times[location])
	}

	// 检查遍历过程中是否发生错误
	if err := rows.Err(); err != nil {
		panic(err)
	}
}

func Find_Max() {
	nums := []float64{float64(Fire_times), float64(Smoke_times), float64(Drawn_times), float64(Wave_times), float64(Water_times), float64(Railing_times), float64(Fall_times)}
	maxNum := nums[0]
	for _, num := range nums {
		maxNum = math.Max(maxNum, num)
	}
	if maxNum == float64(Fire_times) {
		Maximum_type = "火灾"
	}
	if maxNum == float64(Smoke_times) {
		Maximum_type = "吸烟"
	}
	if maxNum == float64(Drawn_times) {
		Maximum_type = "溺水"
	}
	if maxNum == float64(Wave_times) {
		Maximum_type = "挥手"
	}
	if maxNum == float64(Railing_times) {
		Maximum_type = "栏杆"
	}
	if maxNum == float64(Fall_times) {
		Maximum_type = "摔倒"
	}
	if maxNum == float64(Water_times) {
		Maximum_type = "积水"
	}
}
