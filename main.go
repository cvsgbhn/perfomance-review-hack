package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"github.com/urfave/cli/v2"
)

type Task struct {
	Text 	string
	Day		string
	Number	int
}

func main() {
	app := &cli.App{
		Name: "boom",
		Usage: "make an explosive entrance",
		Action: func(c *cli.Context) error {
		  fmt.Println("boom! I say!")
		  return nil
		},
	  }
	  
	var testtask Task

	testtask.Text = "do something"
	testtask.Day = "12 may 2020"
	testtask.Number = 1
	taskJson, err := json.Marshal(testtask)
	if err != nil {
        fmt.Println(err)
        return
	}
	err = ioutil.WriteFile("output.json", taskJson, 0644)
  }