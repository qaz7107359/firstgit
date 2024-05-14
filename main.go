package main

import (
	"TARG_revenue_report_backend/db"
	"TARG_revenue_report_backend/routes"
)

func main() {

	//	引用数据库
	db.InitDB()

	routes.InitRouter()

}
