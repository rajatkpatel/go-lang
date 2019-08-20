package main

import (
	"fmt"
)

func main() {
	var num1 interface{} = 1
	var num2 interface{} = 2
	fmt.Println(Addition(num1, num2))
}

func Addition(num1 interface{}, num2 interface{}) (result interface{}) {
	num1_type := typeof(num1)
	num2_type := typeof(num2)
	if num1_type == "int" && num2_type == "int" {
		return num1.(int) + num2.(int)
	} else if num1_type == "float64" && num2_type == "float64" {
		return num1.(float64) + num2.(float64)
	} else if num1_type == "string" && num2_type == "string" {
		return num1.(string) + num2.(string)
	} else {
		fmt.Println("Error in Addition")
		return "Error in Addition"
	}
}

func typeof(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case float64:
		return "float64"
	case string:
		return "string"
	default:
		return "unknown"
	}
}
