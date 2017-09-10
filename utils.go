package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func loadJSON(fname string, data interface{}) {
	file, err := ioutil.ReadFile(fname)
	checkErr(err)
	json.Unmarshal(file, data)
}
func saveJSON(jsonInterface interface{}, fname string) {
	file, err := json.Marshal(jsonInterface)
	checkErr(err)

	err = ioutil.WriteFile(fname, file, 0644)
	checkErr(err)
}
