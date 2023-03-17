package detect_result

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
)

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

	results_to_front := make([]Results, 0)
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

		err = rows.Scan(&Id, &Photo, &Rate, &Task, &Location)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result_to_front := Results{
			Id:       Id,
			Photo:    Photo,
			Rate:     Rate,
			Task:     Task,
			Location: Location,
		}
		results_to_front = append(results_to_front, result_to_front)
	}
	to_front.JSON(200, results_to_front)
}
