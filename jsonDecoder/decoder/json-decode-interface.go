package decoder

import (
	"encoding/json"
	"fmt"
)

const My_JSON = `{
	"stuff": {
	"onetype": [
	{"id":1,"name":"John Doe"},
	{"id":2,"name":"Don Joeh"}
	],
	"othertype": {"id":2,"company":"ACME"}
	},
	"otherstuff": {
	"thing": [[1,42],[2,2]]
	}
	}`

var my_json_decoder map[string]interface{}

//JsonDecoderToInterface take a json string as input and print the decoded json by storing it in a map interface.
//It will return the error if decoding is failed otherwise return error as nil
func JsonDecoderToInterface(my_json string) error {
	err := json.Unmarshal([]byte(my_json), &my_json_decoder)
	if err != nil {
		fmt.Printf("Error in decoding JSON: %v\n", err)
		return err
	}

	fmt.Printf("%+v\n", my_json_decoder)
	parseMap(my_json_decoder)
	return nil
}

//parseMap will parse the map interface and print the underlying concrete value and its key
func parseMap(json_decoder map[string]interface{}) {
	for key, value := range json_decoder {
		switch concreteValue := value.(type) {
		case map[string]interface{}:
			fmt.Println(key)
			parseMap(value.(map[string]interface{}))
		case []interface{}:
			fmt.Println(key)
			parseArray(value.([]interface{}))
		default:
			fmt.Println(key, ":", concreteValue)
		}
	}
}

//parseArray will parse the slice of interfaces and print the underlying concrete value and its key
func parseArray(json_array_decode []interface{}) {
	for key, value := range json_array_decode {
		switch concreteValue := value.(type) {
		case map[string]interface{}:
			fmt.Println("Index", key)
			parseMap(value.(map[string]interface{}))
		case []interface{}:
			fmt.Println("Index", key)
			parseArray(value.([]interface{}))
		default:
			fmt.Println("Index", key, ":", concreteValue)
		}
	}
}
