package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type Color struct {
	Name        string `json:"name"`
	Hexadecimal string `json:"hexadecimal"`
}

func main() {
	color := Color{"Orange", "#FFA500"}

	colorBytes, err := json.Marshal(color)

	if err != nil {
		log.Fatal(err)
	}

	colorJson := bytes.NewBuffer(colorBytes)

	fmt.Println(colorJson)

	colorMap := map[string]string{
		"name":        "Black",
		"hexadecimal": "#000",
	}

	colorMapBytes, err := json.Marshal(colorMap)

	if err != nil {
		log.Fatal(err)
	}

	colorJsonTwo := bytes.NewBuffer(colorMapBytes)

	fmt.Println(colorJsonTwo)
}
