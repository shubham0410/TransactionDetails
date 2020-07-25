package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"transactionService/models"

	"github.com/gorilla/mux"
	"github.com/spf13/cast"
)

// return response format
type Response struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Model interface{} `json:"model"`
}

//AddTransactionDetails ... add new transaction record to the database
func AddTransactionDetails(w http.ResponseWriter, r *http.Request) Response {
	params := mux.Vars(r)
	txnId := params["transaction_id"]
	fmt.Println("Transaction Id is:", txnId)
	var response Response
	response.Code = 200
	response.Model = nil
	response.Msg = "success"
	var transaction models.Transaction
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &transaction)
	if err != nil {
		fmt.Println("Error while unmarshalling body", err)
		response.Code = 400
		response.Msg = "payload is not correct"
		return response
	}
	transaction.Id = cast.ToInt(txnId)
	fmt.Println("DETAILS ARE:", transaction.Id, transaction.Type)
	_, err = models.AddTransaction(&transaction)
	if err != nil {
		fmt.Println("DB Error is", err)
		response.Code = 400
		response.Msg = "Databse Error"
		return response
	}

	return response
}

//GetTransactionDetails ... get transation details on the basis of ID
func GetTransactionDetails(w http.ResponseWriter, r *http.Request) Response {
	params := mux.Vars(r)
	txnId := params["transaction_id"]
	fmt.Println("Transaction Id is:", txnId)
	var response Response
	response.Code = 200
	response.Model = nil
	response.Msg = "success"
	txdnDetails, err := models.GetTransactionById(cast.ToInt(txnId))
	if err != nil {
		fmt.Println("DB Error is", err)
		response.Code = 400
		response.Msg = "Databse Error"
		return response
	}
	response.Model = txdnDetails
	return response
}

//GetTransactionDetailsByType ... get transation details on the basis of type
func GetTransactionDetailsByType(w http.ResponseWriter, r *http.Request) Response {
	params := mux.Vars(r)
	txntype := params["type"]
	fmt.Println("Type is:", txntype)
	var response Response
	response.Code = 200
	response.Model = nil
	response.Msg = "success"
	txnIds, err := models.GetTransactionByType(txntype)
	if err != nil || len(txnIds) == 0 {
		fmt.Println("DB Error is", err)
		response.Code = 400
		response.Msg = "Databse Error"
		return response
	}
	response.Model = txnIds
	return response
}

//GetTransactionSumByParentId ... get sum on the basis of parentId i.e  If A is parent of B and C,  and C is parent of D and E . sum(A) = B + C + D + E
func GetTransactionSumByParentId(w http.ResponseWriter, r *http.Request) Response {
	params := mux.Vars(r)
	txnId := params["transaction_id"]
	fmt.Println("Type is:", txnId)
	var response Response
	response.Code = 200
	response.Model = nil
	response.Msg = "success"
	sum, err := models.GetAmountSumByIdSingleQuery(cast.ToInt(txnId))
	if err != nil {
		fmt.Println("DB Error is", err)
		response.Code = 400
		response.Msg = "Databse Error"
		return response
	}
	response.Model = sum
	models.Sum = 0
	models.Flag = false
	return response
}
