package main

import (
	"log"
	"os"

	"brendanthompson.dev/terrago"
)

func main() {
	workingDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	options := &terrago.Options{
		TerraformDir: workingDirectory,
	}

	_, err = terrago.InitAndPlan(options)
	if err != nil {
		log.Fatal(err)
	}
}
