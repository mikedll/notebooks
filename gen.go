package main

import (
	"pkg"
	"encoding/json"
	"io/ioutil"
	"log"
)

func graph() {
 	result := pkg.Problem3d2()

	output, err := json.Marshal(*result)
	if err != nil {
		log.Fatal(err)
	}

	/*
	  var outputPretty []byte
	  outputPretty, err = json.MarshalIndent(*result, "", "  ")
	  if err != nil {
	    log.Fatal(err)
	  }
	*/

	ioutil.WriteFile("data.json", output, 0644)
	log.Println("Wrote data.json")
	// println(string(outputPretty))	
}

func main() {
	graph()
}
