package controllers

import (
	"bytes"
	"fetch-assignment/internal/database"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestValidReceipts_Process(t *testing.T) {
	var validJsonTable []string = []string{
		`{
		  "retailer": "Target",
		  "purchaseDate": "2022-01-01",
		  "purchaseTime": "13:01",
		  "items": [
		    {
		      "shortDescription": "Mountain Dew 12PK",
		      "price": "6.49"
		    },{
		      "shortDescription": "Emils Cheese Pizza",
		      "price": "12.25"
		    },{
		      "shortDescription": "Knorr Creamy Chicken",
		      "price": "1.26"
		    },{
		      "shortDescription": "Doritos Nacho Cheese",
		      "price": "3.35"
		    },{
		      "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
		      "price": "12.00"
		    }
		  ],
		  "total": "35.35"
		}`, `{
		  "retailer": "M&M Corner Market",
		  "purchaseDate": "2022-03-20",
		  "purchaseTime": "14:33",
		  "items": [
		    {
		      "shortDescription": "Gatorade",
		      "price": "2.25"
		    },{
		      "shortDescription": "Gatorade",
		      "price": "2.25"
		    },{
		      "shortDescription": "Gatorade",
		      "price": "2.25"
		    },{
		      "shortDescription": "Gatorade",
		      "price": "2.25"
		    }
		  ],
		  "total": "9.00"
		}`,
	}
	var validPointsTable = []int{28, 109}
	for idx, v := range validJsonTable {
		r, err := http.NewRequest("GET", "/receipts/process", bytes.NewReader([]byte(v)))
		if err != nil {
			t.Error("Test data cannot be serialized")
		}
		recorder := httptest.NewRecorder()
		println(recorder.Body)
		ProcessReceipts(recorder, r)
		if recorder.Code >= 300 {
			t.Error(recorder.Body)
		}
		i := strconv.Itoa(idx + 1)
		// we're not going to mock the db for time sake
		db := database.GetInstance().TxTable
		if db[i].Points != validPointsTable[idx] {
			t.Errorf("Expected %d and found %d", validPointsTable[idx], db[i].Points)
		}
	}
}
