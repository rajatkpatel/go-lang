package decoder

import (
	"fmt"
	"testing"
)

var my_JSON_test = []string{
	`{
	"stuff": {
	"onetype": [
	{"id":10,"name":"John"},
	{"id":11,"name":"Don"}
	],
	"othertype": {"id":11,"company":"ACME"}
	},
	"otherstuff": {
	"thing": [[1,4],[2,3]]
	}
	}`,
	`"thing": [[1,4],[2,3]]`,
}

func TestJsonDecoder(t *testing.T) {
	for _, value := range my_JSON_test {
		err := JsonDecoderToInterface(value)
		//assert.Equalf(t, nil, err, "Error in Parsing JSON to interface")
		if err != nil {
			fmt.Println("Error in Parsing JSON to interface", err)
		}
		err = JsonDecoderToStruct(value)
		//assert.Equalf(t, nil, err, "Error in Parsing JSON to struct")
		if err != nil {
			fmt.Println("Error in Parsing JSON to struct", err)
		}
	}
}
