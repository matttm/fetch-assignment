package models

type Transaction struct {
	Id      string
	Receipt *Receipt
	Points  int
}
