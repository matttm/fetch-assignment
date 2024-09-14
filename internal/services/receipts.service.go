package services

import (
	"errors"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"

	"fetch-assignment/internal/models"
)

var DATE_FORMAT = "2022-01-01"
var TIME_FORMAT = "13:01"
var BEFORE, _ = time.Parse(TIME_FORMAT, "16:00")
var AFTER, _ = time.Parse(TIME_FORMAT, "14:00")

// for sake of time, we'll use an in-memory db
var db = make(map[int64]*models.Transaction)

func ProcessReceipts(receipt *models.Receipt) (int64, error) {
	points := 0
	// One point for every alphanumeric character in the retailer name.
	for _, r := range receipt.Retailer {
		if isAlphanumeric(r) {
			points++
		}
	}
	total, err := strconv.ParseFloat(receipt.Total, 32)
	if err != nil {
		return -1, errors.New("Total cannot be parsed")
	}
	dollars := int(total)
	decimal := total - float64(dollars)
	if decimal == 0 {
		points += 50
	}
	// 25 points if the total is a multiple of 0.25.
	cents := int(total * 100) // if we have $1.25, we have 125 cents
	if cents%25 == 0 {
		points += 25
	}
	// 5 points for every two items on the receipt.
	pairsOfitems := len(receipt.Items) / 2
	points = pairsOfitems * 5
	// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
	for _, item := range receipt.Items {
		description := strings.TrimSpace(item.ShortDesc)
		if len([]rune(description))%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 32)
			if err != nil {
				return -1, errors.New("Price cannot be parsed")
			}
			points += int(math.Ceil(price * 0.2))
		}

	}
	// 6 points if the day in the purchase date is odd.
	date, err := time.Parse(DATE_FORMAT, receipt.PurchaseDate)
	if err != nil {
		return -1, errors.New("Purchase date cannot be parsed")
	}
	day := date.Weekday()
	if int(day)%2 == 1 {
		points += 6
	}
	// 10 points if the time of purchase is after 2:00pm and before 4:00pm
	purchaseTime, err := time.Parse(TIME_FORMAT, receipt.PurchaseTime)
	if err != nil {
		return -1, errors.New("Purchase date cannot be parsed")
	}
	if purchaseTime.After(AFTER) && purchaseTime.Before(BEFORE) {
		points += 10
	}
	// store in db
	var id int64 = int64(len(db) + 1)
	db[id] = &models.Transaction{
		Id:      id,
		Receipt: receipt,
		Points:  points,
	}
	return id, nil
}

func GetPoints(id int64) (int, error) {
	if db[id] == nil {
		return db[id].Points, nil
	} else {
		return 0, errors.New("No such receipt")
	}
}

func isAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsNumber(r)
}
