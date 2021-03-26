package http

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/wesley-vinicius/planning_finance/adapter/http/actuator"
	"github.com/wesley-vinicius/planning_finance/adapter/http/transaction"
	"net/http"
)

func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateTransactions)

	http.HandleFunc("/health", actuator.Health)

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)
}
