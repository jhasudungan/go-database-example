package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

/**
Notes from ChatGPT About Using Pointer:

In general, if you don't have any specific requirements or constraints,
returning a list of structs ([]Transaction) is a safe and straightforward choice.
But if you need to optimize for performance or have specific use cases that require
pointer access to the original struct,
returning a list of pointers ([]*Transaction) could be a better fit.

*/
type TransactionsRepository interface {
	FindAll() ([]Transaction, error)
	FindAllWithPage(page int64) ([]Transaction, error)
	Insert(newTransaction Transaction) (Transaction, error)
	Update(newTransaction Transaction) (Transaction, error)
	Delete(transactionId string) error
}

type TransactionsRepositoryImpl struct {
	DbTools DbTools
}

func (t *TransactionsRepositoryImpl) FindAll() ([]Transaction, error) {

	result := make([]Transaction, 0)

	con, err := t.DbTools.CreateConnection()
	defer con.Close()

	if err != nil {
		return result, err
	}

	sql := fmt.Sprint("select transaction_id, transaction_total, ",
		"transaction_person_in_charge, transaction_date, ",
		"created_date, last_modified ",
		"from public.transactions")

	rows, err := con.Query(sql)
	defer rows.Close()

	if err != nil {
		return result, err
	}

	for rows.Next() {

		transaction := Transaction{}

		err := rows.Scan(&transaction.TransactionId,
			&transaction.TransactionTotal,
			&transaction.TransactionPIC,
			&transaction.TransactionDate,
			&transaction.CreatedDate,
			&transaction.LastModified)

		if err != nil {
			return result, err
		}

		result = append(result, transaction)

	}

	return result, nil

}

func (t *TransactionsRepositoryImpl) FindAllWithPage(page int64) ([]Transaction, error) {

	result := make([]Transaction, 0)

	con, err := t.DbTools.CreateConnection()
	defer con.Close()

	if err != nil {
		return result, err
	}

	sql := fmt.Sprint("select transaction_id, transaction_total, ",
		"transaction_person_in_charge, transaction_date, ",
		"created_date, last_modified ",
		"from public.transactions offset $1 limit $2")

	if page <= 0 {
		page = 1
	}

	temp := os.Getenv("LIMIT_PAGINATION_PER_PAGE")

	limit, err := strconv.Atoi(temp)

	if err != nil {
		return result, err
	}

	offset := (page - 1) * int64(limit)

	rows, err := con.Query(sql, offset, limit)
	defer rows.Close()

	if err != nil {
		log.Print("Error " + err.Error())
		return result, err
	}

	for rows.Next() {

		transaction := Transaction{}

		err := rows.Scan(&transaction.TransactionId,
			&transaction.TransactionTotal,
			&transaction.TransactionPIC,
			&transaction.TransactionDate,
			&transaction.CreatedDate,
			&transaction.LastModified)

		if err != nil {
			return result, err
		}

		result = append(result, transaction)

	}

	return result, nil

}

func (t *TransactionsRepositoryImpl) Insert(newTransaction Transaction) (Transaction, error) {

	con, err := t.DbTools.CreateConnection()
	defer con.Close()

	if err != nil {
		return newTransaction, err
	}

	sql := fmt.Sprint("insert into public.transactions ",
		"(transaction_id, transaction_total, ",
		"transaction_person_in_charge, transaction_date, ",
		"created_date, last_modified) ",
		"VALUES($1, $2, $3, now(),",
		"now(), now())")

	// Get The Value From the pointer
	_, err = con.Exec(sql, newTransaction.TransactionId,
		newTransaction.TransactionTotal,
		newTransaction.TransactionPIC)

	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}

func (t *TransactionsRepositoryImpl) Update(newTransaction Transaction) (Transaction, error) {

	con, err := t.DbTools.CreateConnection()
	defer con.Close()

	if err != nil {
		return newTransaction, err
	}

	sql := fmt.Sprint("update public.transactions ",
		"set transaction_total = $1,",
		"transaction_person_in_charge = $2, ",
		"last_modified = now() ",
		"where transaction_id = $3")

	// Get The Value From the pointer
	_, err = con.Exec(sql, newTransaction.TransactionTotal,
		newTransaction.TransactionPIC,
		newTransaction.TransactionId)

	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}

func (t *TransactionsRepositoryImpl) Delete(transactionId string) error {

	con, err := t.DbTools.CreateConnection()
	defer con.Close()

	if err != nil {
		return err
	}

	sql := fmt.Sprint("delete from public.transactions where transaction_id = $1")

	// Get The Value From the pointer
	_, err = con.Exec(sql, transactionId)

	if err != nil {
		return err
	}

	return nil
}
