package main

import "fmt"

func main() {
	user := map[string]string{
		"name":    "John Cena",
		"address": "Grove Street",
		"email":   "some@email.com",
	}

	fmt.Println(user)
	fmt.Println(user["address"])

	mapWithMap := map[string]map[string]string{
		"something": {
			"second":  "value",
			"another": "thing",
		},
		"valueToDelete": {
			"value": "toDelete",
		},
	}

	fmt.Println(mapWithMap)

	delete(mapWithMap, "valueToDelete")

	fmt.Println(mapWithMap)

	mapWithMap["addedKey"] = map[string]string{
		"the": "green",
	}

	fmt.Println(mapWithMap)
}
