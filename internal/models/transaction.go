package models

type Transaction struct {
	Id      int64
	Receipt *Receipt
	Points  int
}
