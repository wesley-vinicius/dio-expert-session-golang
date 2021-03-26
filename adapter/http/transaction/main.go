package transaction

import (
	"encoding/json"
	"fmt"
	"github.com/wesley-vinicius/planning_finance/model/transaction"
	"github.com/wesley-vinicius/planning_finance/util"
	"io/ioutil"
	"net/http"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "json")

	transactions := transaction.Transactions{
		transaction.Transaction{
			Title:    "Salario",
			Amount:   1200.0,
			Type:     0,
			CreateAt: util.StringToTime("2021-03-24T20:45:26"),
		},
	}
	_ = json.NewEncoder(w).Encode(transactions)

}

func CreateTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	res := transaction.Transactions{}
	body, _ := ioutil.ReadAll(r.Body)

	_ = json.Unmarshal(body, &res)

	fmt.Print(res)

}
