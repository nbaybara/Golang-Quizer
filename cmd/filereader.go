/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func readJsonFile(w http.ResponseWriter, r *http.Request) {
	readJson()
}

func readJson() {

	jsonFile, err := os.Open("assets/problems.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened problems.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &questions)

}
