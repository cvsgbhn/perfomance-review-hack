package main

import (
  "fmt"
  "log"
  "os"
  "time"
  "encoding/json"
  "io/ioutil"

  "github.com/urfave/cli/v2"
)
// 1. записать в файл
// 2. отметить как выполненное или проваленное
// 3. добавить сохранение даты выполнения/провала
// 4. выводить на экран красивенько все задачи

type Task struct {
  Text string
  Status string
  Start_date string
  End_date string
}

func save_task() {
  var new_task Task
  new_task.Text = "new task in process now"
  new_task.Status = "in process"
  t := time.Now()
  new_task.Start_date = t.String()
  b, err := json.Marshal(new_task)
  if err != nil {
    fmt.Println("error occured")
  }
  fmt.Println(string(b))
  buf := []byte(b)
  err = ioutil.WriteFile("./taskline", buf, 0644)
  if err != nil {
    fmt.Println("error occured")
  }
}

func main() {
  app := &cli.App{
    Name: "todo",
    Usage: "be organized",
    Action: func(c *cli.Context) error {
      fmt.Println("Hello friend!")
      var name string
      if c.NArg() > 0 {
        name = c.Args().Get(0)
      }
      if name == "add" {
        save_task()
      }
    },
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}