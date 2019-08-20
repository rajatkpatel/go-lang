package main

import (
	"fmt"

	"github.com/owner/jsonDecoder/decoder"
)

func main() {
	if err := decoder.JsonDecoderToStruct(decoder.My_JSON); err != nil {
		fmt.Println("Error in decoding from JSON to struct", err)
	}
	if err := decoder.JsonDecoderToInterface(decoder.My_JSON); err != nil {
		fmt.Println("Error in decoding from JSON to interface", err)
	}
}
