package main

import (
	"fmt"

	"github.com/go-lang/jsonDecoder/decoder"
)

func main() {
	if err := decoder.JsonDecoderToStruct(decoder.My_JSON); err != nil {
		fmt.Println("Error in decoding from JSON to struct", err)
	}
	if err := decoder.JsonDecoderToInterface(decoder.My_JSON); err != nil {
		fmt.Println("Error in decoding from JSON to interface", err)
	}
}
