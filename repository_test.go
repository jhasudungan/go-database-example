package main

import (
	"fmt"
	"log"
	"testing"
)

func IdForTest() string {
	return "TRX_TEST"
}

func TestFindAll(t *testing.T) {

	dbtools := &DbToolsImpl{}

	transactionsRepository := &TransactionsRepositoryImpl{
		DbTools: dbtools}

	result, err := transactionsRepository.FindAll()

	if err != nil {
		log.Fatal("Error : " + err.Error())
	}

	fmt.Printf("result: %v\n", result)
}

func TestFindWithPage(t *testing.T) {

	dbtools := &DbToolsImpl{}

	transactionsRepository := &TransactionsRepositoryImpl{
		DbTools: dbtools}

	result, err := transactionsRepository.FindAllWithPage(int64(2))

	if err != nil {
		log.Fatal("Error : " + err.Error())
	}

	fmt.Printf("result: %v\n", result)
}

func TestInsert(t *testing.T) {

	dbtools := &DbToolsImpl{}

	transactionsRepository := &TransactionsRepositoryImpl{
		DbTools: dbtools}

	newTransaction := Transaction{}
	newTransaction.TransactionId = IdForTest()
	newTransaction.TransactionTotal = 150.0
	newTransaction.TransactionPIC = "JRX"

	result, err := transactionsRepository.Insert(newTransaction)

	if err != nil {
		log.Fatal("Error : " + err.Error())
	}

	fmt.Printf("result: %v\n", result)

}

func TestUpdate(t *testing.T) {

	dbtools := &DbToolsImpl{}

	transactionsRepository := &TransactionsRepositoryImpl{
		DbTools: dbtools}

	newTransaction := Transaction{}
	newTransaction.TransactionId = IdForTest()
	newTransaction.TransactionTotal = 170.0
	newTransaction.TransactionPIC = "JRX"

	result, err := transactionsRepository.Update(newTransaction)

	if err != nil {
		log.Fatal("Error : " + err.Error())
	}

	fmt.Printf("result: %v\n", result)
}

func TestDelete(t *testing.T) {

	dbtools := &DbToolsImpl{}

	transactionsRepository := &TransactionsRepositoryImpl{
		DbTools: dbtools}

	err := transactionsRepository.Delete(IdForTest())

	if err != nil {
		log.Fatal("Error : " + err.Error())
	}

}
