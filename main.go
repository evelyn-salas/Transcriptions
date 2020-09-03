package main

import (
	"fmt"
	"os"
)

//not sure here
type originalText struct {
	originalText []originalText `json:"originalText"`
}

func main() {
	//open json files
	jsonFile, err1 := os.Open("intentall-off.json")
	jsonFile, err2 := os.Open("intent.out.json")

	//in case file does not open
	if err1 != nil {
		fmt.Println(err1)
	}
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("Sucessfully Opened files")

	//defer the closing of the file
	defer jsonFile.Close()

	//read open file as byte array
	//byteValue, _ := ioutil.ReadAll(jsonFile)

	//initialize array
	//var -- --
}
