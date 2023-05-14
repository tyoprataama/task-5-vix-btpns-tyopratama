package main

import (
	"github.com/tyoprataama/task-5-vix-btpns-tyopratama/database"
	"github.com/tyoprataama/task-5-vix-btpns-tyopratama/router"
)

func main() {
	database.InitDB()
	database.MigrateDB()
	r := router.RouteInit()
	r.Run(":8080")
}
