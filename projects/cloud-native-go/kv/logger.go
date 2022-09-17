package main

type TransactionLogger interface {
	WritePut(key, value string)
	WriteDelete(key string)
}
