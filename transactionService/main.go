package main

import (
	"fmt"
	"net/http"
	"transactionService/handlers"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	DatabaseConnection()
	fmt.Println("Server running on:http://localhost:7000")
	router := mux.NewRouter()
	// below scenerio is for options method when request is coming from browser. Not required in case we are hitting the request from Postman.
	router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		w.WriteHeader(http.StatusNoContent)
		return
	})

	router.HandleFunc("/transactionservice/transaction/{transaction_id}", handlers.TransactionHandler)
	router.HandleFunc("/transactionservice/{type}", handlers.TransactionTypeHandler)
	router.HandleFunc("/transactionservice/sum/{transaction_id}", handlers.TransactionSumHandler)
	http.ListenAndServe(":7000", router)
}

/// DatabaseConnection ...Make dataBase Connection
func DatabaseConnection() bool {
	_, err := orm.GetDB()
	if err == nil {
		return true
	}
	//{username}:{password}@(localhost:3306)/test?interpolateParams=true
	mysqlConf := "root:root@(localhost:3306)/test?interpolateParams=true"
	if err := orm.RegisterDataBase("default", "mysql", mysqlConf); err != nil {
		fmt.Println("Mysql Connection Create DB Err:", err)
	}

	err = mysqlTest()
	if err != nil {
		fmt.Println("Mysql Connection Ping DB Err:", err)
	}

	fmt.Println("Mysql Connection Ping Is Done")
	orm.Debug = true
	return true

}
func mysqlTest() error {
	o := orm.NewOrm()
	_, err := o.Raw("SELECT 1").Exec()
	return err
}
