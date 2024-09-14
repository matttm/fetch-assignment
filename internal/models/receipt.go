package models

type Receipt struct {
	Id           int
	Retailer     string
	PurchaseDate string
	PurchaseTime string
	Items        []Item
	Total        string
}
