package transaction

import (
	"encoding/json"
	"fmt"
	"github.com/wesley-vinicius/planning_finance/model/transaction"
	"io/ioutil"
	"net/http"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-type", "json")
	transactions := transaction.GetAll()
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
