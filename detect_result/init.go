package detect_result

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var mysqlLogger logger.Interface

func init() {
	username := "root"      //账号
	password := "20030729a" //密码
	host := "127.0.0.1"     //数据库地址，可以是Ip或者域名
	port := 3306            //数据库端口
	Dbname := "detection"   //数据库名
	timeout := "10s"        //连接超时，10秒

	mysqlLogger = logger.Default.LogMode(logger.Info)

	// root:root@tcp(127.0.0.1:3306)/gorm?
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// SkipDefaultTransaction: true,
		Logger: mysqlLogger,
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	// 连接成功
	fmt.Println(db) //返回的是一个*gorm.DB,一般放到全局变量中

	DB = db
}

type Camera struct {
	Id            int    `gorm:"column:id;type:int(11);primary_key" json:"id"`
	Number        string `gorm:"column:number;type:varchar(45)" json:"number"`
	Position      string `gorm:"column:position;type:varchar(45)" json:"position"`
	Task          string `gorm:"column:task;type:varchar(45)" json:"task"`
	Rtsp          string `gorm:"column:rtsp;type:varchar(45)" json:"rtsp"`
	Remarks       string `gorm:"column:remarks;type:varchar(45)" json:"remarks"`
	Probability   string `gorm:"column:probability;type:varchar(45)" json:"probability"`
	Framerate     string `gorm:"column:framerate;type:varchar(45)" json:"framerate"`
	Frameratetest string `gorm:"column:frameratetest;type:varchar(45)" json:"frameratetest"`
}

// 解析算法端返回结果的struct
type Results struct {
	Id       int             `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Photo    sql.NullString  `gorm:"column:photo;type:longtext" json:"photo"`
	Rate     sql.NullFloat64 `gorm:"column:rate;type:float" json:"rate"`
	Task     sql.NullString  `gorm:"column:task;type:varchar(255)" json:"task"`
	Location sql.NullString  `gorm:"column:location;type:varchar(255)" json:"location"`
}

// 添加到数据库中带时间的struct
type Add_to_database struct {
	Id       int             `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Photo    sql.NullString  `gorm:"column:photo;type:longtext" json:"photo"`
	Rate     sql.NullFloat64 `gorm:"column:rate;type:float" json:"rate"`
	Task     sql.NullString  `gorm:"column:task;type:varchar(255)" json:"task"`
	Location sql.NullString  `gorm:"column:location;type:varchar(255)" json:"location"`
	Time     sql.NullString  `gorm:"column:time;type:varchar(255)" json:"time"`
	Review   int             `gorm:"column:review;type:tinyint" json:"review"`
}

type AlgorithmReturns struct {
	Photo string  `json:"photo"`
	Rate  float64 `json:"rate"`
	Task  string  `json:"task"`
}

var (
	Db_sql *sql.DB
	err    error
)

func init() {
	Db_sql, err = sql.Open("mysql", "root:20030729a@tcp(localhost:3306)/detection")
	if Db_sql != nil {
		fmt.Println("Db_sql is not nil")
	}
	if err != nil {
		fmt.Println(err)
	}
	if err := Db_sql.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	} else {
		println("yes")
	}
}
