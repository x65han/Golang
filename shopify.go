package main

import (
	"fmt"
	"net/http"
	"strconv"
	"io/ioutil"
	"math"
	"encoding/json"
)

func main() {

	// URL to download the data
	url := "https://shopicruit.myshopify.com/admin/orders.json?page=1&access_token=c32313df0d0ef512ca64d5b336a0d7c6"

	// Setup HTTP Request
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "588bd78c-34b0-3657-20a0-566610abf99b")

	// Execute Request
	res, err := http.DefaultClient.Do(req)

	// Defer to close the request
	defer res.Body.Close()

	// Report if there is any error
	if err != nil{
		fmt.Println("HTTP Error: ", err)
		panic("Cannot get data from URL")
	}

	// Convert response to []byte
	body, _ := ioutil.ReadAll(res.Body)
	byte_arr := []byte(body)

	// Convert []byte to json structure
	var input map[string]interface{}
	json.Unmarshal(byte_arr, &input)

	// Get the orders in the json
	orders := input["orders"]

	// Create a float variable to store the grand total
	var total float64 = 0.0

	switch x := orders.(type) {
		case []interface{}:
			// loop through each order
		    for _, order := range x{
				switch x := order.(type) {
					case map[string]interface{}:
						// Loop through each element in this order
					    for i, e := range x{
					    	// Acess the total_price of each order
					        if i == "total_price"{
					        	// Convert string to float64
					        	price,_ := strconv.ParseFloat(e.(string), 64)
					        	// Accumulate the price
					        	total = total + price
					        }
					    }
				}
		    }
	}
	// Change the float64 to 2 decimal precision
	total = toFixed(total, 2)
	// Finally print out the answer
	fmt.Println("Grand total: ", total)
}

// Custom made toFixed func to change float64 precision
func toFixed(num float64, precision int) float64 {
    output := math.Pow(10, float64(precision))
    return float64(round(num * output)) / output
}

// Custom made round fuc to convert float64 to int with no decimals
func round(num float64) int {
    return int(num + math.Copysign(0.5, num))
}




