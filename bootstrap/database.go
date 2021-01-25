package bootstrap

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aaalik/ke-jepang/helper"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct{}

var DB *sql.DB
var err error

func (db Database) Connect() {
	DATABASE_USER := os.Getenv("DATABASE_USER")
	DATABASE_PASSWORD := os.Getenv("DATABASE_PASSWORD")
	DATABASE_HOST := os.Getenv("DATABASE_HOST")
	DATABASE_NAME := os.Getenv("DATABASE_NAME")

	connectionString := fmt.Sprintf("%v:%v@tcp(%v)/%v", DATABASE_USER, DATABASE_PASSWORD, DATABASE_HOST, DATABASE_NAME)
	DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		helper.Log.Error(err)
	}
}

func (db Database) CloseConnection() {
	DB.Close()
}

func Transaction(queries []string, data [][]interface{}) bool {
	retry := 0
	_retry := ""

	result, err := ExecuteTransaction(queries, data)

	// deadlock mitigation
	for !result && strings.Contains(err, "Error 1213") {
		retry++

		_retry = strconv.Itoa(retry)
		helper.Log.Error("Retrying transaction " + _retry)
		time.Sleep(1 * time.Second)

		result, err = ExecuteTransaction(queries, data)

		if result != false || !strings.Contains(err, "Error 1213") || retry >= 10 {
			helper.Log.Info("Retrying transaction " + _retry + " success")
			break
		}
	}

	return result
}

func ExecuteTransaction(queries []string, data [][]interface{}) (bool, string) {
	tx, err := DB.Begin()
	if err != nil {
		helper.Log.Error(err)
		return false, err.Error()
	}

	defer tx.Rollback()

	for i, v := range queries {
		stmt, err := tx.Prepare(v)
		if err != nil {
			helper.Log.Error(err)
			return false, err.Error()
		}

		defer stmt.Close()

		_, err = stmt.Exec(data[i]...)
		if err != nil {
			helper.Log.Errorf("error: %v data: %v", err, data[i])
			return false, err.Error()
		}
	}

	err = tx.Commit()
	if err != nil {
		helper.Log.Error(err)
		return false, err.Error()
	}

	return true, ""
}
