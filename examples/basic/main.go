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
  
  _, err = terrago.InitAndPlanE(options)
  if err != nil {
    log.Fatal(err)
  }
}
