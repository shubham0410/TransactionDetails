package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"transactionService/controllers"
)

//TransactionHandler ...handle /trnasaction requests
func TransactionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Start of Function TransactionHandler")

	switch true {
	case strings.Contains(r.Method, "PUT"):
		URLReturnResponseJson(w, controllers.AddTransactionDetails(w, r))
	case strings.Contains(r.Method, "GET"):
		URLReturnResponseJson(w, controllers.GetTransactionDetails(w, r))
	default:
		fmt.Println("Handler not found", r.Method)
	}

}

//TransactionTypeHandler ...handle transactionType requests
func TransactionTypeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Start of Function TransactionTypeHandler")

	switch true {
	case strings.Contains(r.Method, "GET"):
		URLReturnResponseJson(w, controllers.GetTransactionDetailsByType(w, r))
	default:
		fmt.Println("Handler not found", r.Method)
	}

}

//TransactionSumHandler ...handle transactionSum requests
func TransactionSumHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Start of Function TransactionSumHandler")

	switch true {
	case strings.Contains(r.Method, "GET"):
		URLReturnResponseJson(w, controllers.GetTransactionSumByParentId(w, r))
	default:
		fmt.Println("Handler not found", r.Method)
	}

}

//URLReturnResponseJson ...write response to writer
func URLReturnResponseJson(w http.ResponseWriter, data interface{}) {
	returnJson, _ := json.Marshal(data)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(returnJson)
}
