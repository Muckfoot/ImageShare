package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"runtime"
)

func checkErr(err error) {
	_, file, line, _ := runtime.Caller(1)
	if err != nil {
		log.Fatalf("file: %s, line: %d, error: %s", file, line, err)
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
