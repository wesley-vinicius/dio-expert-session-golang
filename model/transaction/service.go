package transaction

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "dev:dev@/dev")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Store(transaction *Transaction) error {
	var db = connect()
	defer db.Close()

	stmt, err := db.Prepare("insert into transactions(title, amount, type, created_at) values (?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(transaction.Title, transaction.Amount, transaction.Type, transaction.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}


func GetAll() Transactions{
	var db = connect()
	defer db.Close()

	rows, _ := db.Query("SELECT id ,title, amount, type, created_at FROM transactions")
	var transactionSlice []Transaction
	for rows.Next() {
		var transaction Transaction
		_ = rows.Scan(&transaction.Id,
			&transaction.Title,
			&transaction.Amount,
			&transaction.Type,
			&transaction.CreatedAt)
		transactionSlice = append(transactionSlice, transaction)
	}
	return transactionSlice
}
