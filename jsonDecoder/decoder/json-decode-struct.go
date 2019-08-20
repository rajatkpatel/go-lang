package decoder

import (
	"encoding/json"
	"fmt"
)

type My_JSON_decoder struct {
	Stuff struct {
		Onetype []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"onetype"`
		Othertype struct {
			ID      int    `json:"id"`
			Company string `json:"company"`
		} `json:"othertype"`
	} `json:"stuff"`
	Otherstuff struct {
		Thing [][]int `json:"thing"`
	} `json:"otherstuff"`
}

//JsonDecoderToStruct take a json string as input and print the decoded json by storing it into a struct.
//It will return the error if decoding is failed otherwise return error as nil
func JsonDecoderToStruct(my_json string) error {
	var my_json_decoder My_JSON_decoder
	err := json.Unmarshal([]byte(my_json), &my_json_decoder)
	if err != nil {
		fmt.Printf("Error in decoding JSON: %v\n", err)
		return err
	}

	fmt.Printf("%+v\n", my_json_decoder)
	return nil
}
