package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Fruit struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	fruitJson := `{"name":"orange", "price": 1}`

	var fruit Fruit

	if err := json.Unmarshal([]byte(fruitJson), &fruit); err != nil {
		log.Fatal(err)
	}

	fmt.Println(fruit)

	fruitMap := make(map[string]any)

	if err := json.Unmarshal([]byte(fruitJson), &fruitMap); err != nil {
		log.Fatal(err)
	}

	fmt.Println(fruitMap)
}
