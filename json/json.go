package main

import (
	"encoding/json"
	"fmt"
)

var jsonString string

func main() {

	jsonString = `{
						"data" : {
								"name": "John",	
								"age": 30, 
								"city": "New York"
							}	
}`
	var mapData map[string]map[string]interface{}

	if err := json.Unmarshal([]byte(jsonString), &mapData); err != nil {
		panic(err)
	}

	fmt.Println(mapData)

	jsonString2, err := json.Marshal(mapData)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonString2))

}
