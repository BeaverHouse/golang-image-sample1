package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sample/test/src/utils"
)

func main() {

	jsonString := os.Args[1]

	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(jsonString), &jsonMap)

	nasPath := jsonMap["nasPath"].(string)
	inputDir := jsonMap["inputDir"].(string)
	outputDir := jsonMap["outputDir"].(string)

	fmt.Println(nasPath)
	fmt.Println(inputDir)
	fmt.Println(outputDir)

	if err := os.MkdirAll(nasPath+outputDir, os.ModePerm); err != nil {
		panic(err)
	}

	utils.ChangeImage(nasPath+inputDir, nasPath+outputDir)
}
