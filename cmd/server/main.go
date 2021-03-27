package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/wesley-vinicius/planning_finance/adapter/http"
)

func main() {
	http.Init()
}
