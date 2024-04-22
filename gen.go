package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"pkg"
)

func main() {
 	result := pkg.Exercise3p3d5b()

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
