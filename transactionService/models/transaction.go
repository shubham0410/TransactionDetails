package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

// structure for transation table
type Transaction struct {
	Id       int     `orm:"column(id);pk" json:"id"`
	Amount   float32 `orm:"column(amount);size(45);null" json:"amount"`
	Type     string  `orm:"column(type);size(90);null" json:"type"`
	ParentId *int    `orm:"column(parentId);" json:"parent_id"`
}

func (t *Transaction) TableName() string {
	return "transaction"
}

func init() {
	orm.RegisterModel(new(Transaction))
}

// AddTransaction insert a new Transaction into database and returns
// last inserted Id on success.
func AddTransaction(m *Transaction) (id int64, err error) {
	fmt.Println("Adding transaction")
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTransactionById retrieves Transaction by Id. Returns error if
// Id doesn't exist
func GetTransactionById(id int) (v *Transaction, err error) {
	o := orm.NewOrm()
	v = &Transaction{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

//GetTransactionByType ... give all the transaction for the particular type
func GetTransactionByType(txnType string) ([]orm.Params, error) {
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("select id from transaction where type= ?", txnType).Values(&maps)
	if err == nil && num > 0 {
		fmt.Println("number of rows are", num)
		return maps, nil
	}
	return nil, err
}

//GetAmountSumById ... give sum of all the transaction for particular parentId.
// tried to use procedure but not able to get the sum of all the parents
func GetAmountSumById(txnId int) ([]orm.Params, error) {
	o := orm.NewOrm()
	var maps []orm.Params
	_, err := o.Raw("CALL calctotal(?, @total)", txnId).Exec()
	if err != nil {
		fmt.Println("GetAmountSumById:Error is", err)
		return nil, err
	}
	_, err = o.Raw("SELECT @total as sum").Values(&maps)
	fmt.Println("ERRor in Db", err)
	if err == nil {
		return maps, nil
	}
	return nil, err
}

var Sum float32 = 0
var Flag bool

//GetAmountSumByIdTemp ... give sum of all the transaction for particular parentId. so many network calls
func GetAmountSumByIdTemp(txnId int) (float32, error) {
	o := orm.NewOrm()
	var psum float32
	if !Flag {
		err := o.Raw("select amount from transaction where id= ?", txnId).QueryRow(&psum)
		if err != nil {
			fmt.Println("number of rows are")
			return 0, nil
		}
		Sum += psum
	}
	Flag = true
	var txn []Transaction
	_, err := o.Raw("select * from transaction where parentId= ?", txnId).QueryRows(&txn)
	for _, value := range txn {
		Sum += value.Amount
		fmt.Println("SumIs:", Sum)
		//recursively calling to get the sum of child transactionIds.
		GetAmountSumByIdTemp(value.Id)
	}
	return Sum, err
}

//GetAmountSumByIdSingleQuery ... give sum of all the transaction for particular parentId in single query(by this we can save network calls).
func GetAmountSumByIdSingleQuery(txnId int) (int, error) {
	o := orm.NewOrm()
	var totSum int
	err := o.Raw(`select sum(totalsum) from (select sum(amount) as totalsum from transaction where find_in_set(id, 
		(select GROUP_CONCAT(y SEPARATOR ',') FROM (SELECT @pId:=(SELECT GROUP_CONCAT(id SEPARATOR ',') 
		FROM transaction WHERE FIND_IN_SET(parentId, @pId)) AS y FROM transaction JOIN(SELECT @pId:=?) tmp)a)) 
		union all select amount as totalsum  from transaction where id=?) as x`, txnId, txnId).QueryRow(&totSum)
	if err == nil {
		fmt.Println("total sum is:", totSum)
		return totSum, nil
	}
	return 0, err
}
