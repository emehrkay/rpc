package tests

import (
	"encoding/json"
	"fmt"

	"github.com/emehrkay/rpc/storage"
)

var (
	Cases = []struct {
		Json    string
		Points  uint64
		Receipt *storage.Receipt
	}{
		{
			Json: `{
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
			}`,
			Points: 28,
			// Total Points: 28
			// Breakdown:
			// 	 6 points - retailer name has 6 characters
			// 	10 points - 5 items (2 pairs @ 5 points each)
			// 	 3 Points - "Emils Cheese Pizza" is 18 characters (a multiple of 3)
			// 				item price of 12.25 * 0.2 = 2.45, rounded up is 3 points
			// 	 3 Points - "Klarbrunn 12-PK 12 FL OZ" is 24 characters (a multiple of 3)
			// 				item price of 12.00 * 0.2 = 2.4, rounded up is 3 points
			// 	 6 points - purchase day is odd
			//   + ---------
			//   = 28 points
		},
		{
			Json: `{
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
			Points: 109,
			// Total Points: 109
			// Breakdown:
			// 	50 points - total is a round dollar amount
			// 	25 points - total is a multiple of 0.25
			// 	14 points - retailer name (M&M Corner Market) has 14 alphanumeric characters
			// 				note: '&' is not alphanumeric
			// 	10 points - 2:33pm is between 2:00pm and 4:00pm
			// 	10 points - 4 items (2 pairs @ 5 points each)
			//   + ---------
			//   = 109 points
		},
	}
)

func init() {
	for i, c := range Cases {
		var rec storage.Receipt
		err := json.Unmarshal([]byte(c.Json), &rec)
		if err != nil {
			panic(fmt.Sprintf(`unable to marshal json at index: %d to recepit -- %v`, i, err))
		}

		c.Receipt = &rec
		Cases[i] = c
	}
}
