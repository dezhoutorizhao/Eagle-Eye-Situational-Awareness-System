package detect_result

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "strongwill.com/utils"
)

var Result_to_front_logs []Add_to_database

func Result_to_front(to_front *gin.Context) {
	////查询
	//var resultsList []Results
	////将查询结果append为一个数组，返回给前端
	////var to_front []Results
	//DB.Find(&resultsList)
	////for _, results := range resultsList {
	////	//fmt.Println(results)
	////	//to_front.append()
	////	//单条结果为result
	////	result := Results{
	////		Id: Id,
	////	}
	////}
	//jsonResult, err := json.MarshalIndent(resultsList, "", "    ")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//result_to_front.JSON(200, jsonResult)

	rows, err := Db_sql.Query("SELECT * FROM detection.results")

	results_to_front := make([]Add_to_database, 0)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for rows.Next() {
		var Id int
		var Photo sql.NullString
		var Rate sql.NullFloat64
		var Task sql.NullString
		var Location sql.NullString
		var Time sql.NullString
		var Review int

		err = rows.Scan(&Id, &Photo, &Rate, &Task, &Location, &Time, &Review)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result_to_front := Add_to_database{
			Id:       Id,
			Photo:    Photo,
			Rate:     Rate,
			Task:     Task,
			Location: Location,
			Time:     Time,
			Review:   Review,
		}
		results_to_front = append(results_to_front, result_to_front)
	}
	Result_to_front_logs = results_to_front

	to_front.JSON(200, results_to_front)
}
