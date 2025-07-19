package main

import (
	"fmt"
	"os"
	"path/filepath"
	"stereo-mono-converter-go/conversion"
	"strings"
)
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Invalid argument")
		return
	}

	switch os.Args[1] {
	case "convert":
		if len(os.Args) < 3 {
			fmt.Println("Missing file name")
			return
		}
		
		infilePath := filepath.Join(".", os.Args[2])
		if strings.ToLower(filepath.Ext(infilePath)) != ".wav" {
			fmt.Println("only .wav files are supported")
			return
		}

		result := conversion.ConvertToMono(infilePath)
		fmt.Println(result)

	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}
