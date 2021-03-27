package transaction

import (
	"encoding/json"
	"github.com/wesley-vinicius/planning_finance/model/transaction"
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
	var t transaction.Transaction
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		w.WriteHeader(http.StatusMethodNotAllowed)
		print(err)
		return
	}
	err = transaction.Store(&t)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	return
}
