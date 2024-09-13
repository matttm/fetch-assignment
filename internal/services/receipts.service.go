package services

import (
	"math"
	"strconv"
	"strings"
	"unicode"

	"fetch-assignment/internal/models"
)

func ProcessReceipts(receipt *models.Receipt) bool {
	points := 0
	// One point for every alphanumeric character in the retailer name.
	for _, r := range receipt.Retailer {
		if isAlphanumeric(r) {
			points++
		}
	}
	total, err := strconv.ParseFloat(receipt.Total)
	if err != nil {
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
			price = math.Ceil(receipt.Price * 0.2)
		}

	}
}

func GetPoints(id int) (int, error) {
	return 0, nil
}

func isAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsNumber(r)
}
