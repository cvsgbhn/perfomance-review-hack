package main

import (
  "fmt"
  "log"
  "os"
  "time"
  "encoding/json"
  "bufio"

  "github.com/urfave/cli/v2"
)
// --- 1. записать в файл ---  DoNe
// --- 2. выводить на экран красивенько все задачи --- DoNe
// 2.1. форматировать дату нормально
// 3. отметить как выполненное или проваленное
// 4. добавить сохранение даты выполнения/провала
// 4. добавить причину выполнения/провала (тэги?) и как можно улучшить выполнение в следующий раз

type Task struct {
  Text string
  Status string
  Start_date string
  End_date string
}

func ShowAllTasks() {
  // Open tasks.json in read-mode.
  file, _ := os.Open("todo/tasks.json")
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    textFromFile := scanner.Text()
    task := Task{}
    err := json.Unmarshal([]byte(textFromFile), &task)
    if err != nil {
      fmt.Println(err)
    }
    fmt.Println(task.Start_date, "  ", task.Text, "  ", task.Status)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "shouldn't see an error scanning a string")
	}

}

func SaveTask(taskText string) {
  var new_task Task
  new_task.Text = taskText
  new_task.Status = "in process"
  t := time.Now()
  new_task.Start_date = t.String()
  b, err := json.Marshal(new_task)
  if err != nil {
    fmt.Println("error occured")
  }
  fmt.Println(string(b))
  // Open tasks.json in append-mode.
  f, _ := os.OpenFile("todo/tasks.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
  // Append our json to tasks.json
  if _, err = f.Write(b); err != nil {
    panic(err)
  }
  if _, err = f.Write([]byte("\n")); err != nil {
    panic(err)
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
        SaveTask(c.Args().Get(1))
      } 
      if name == "ls" {
        ShowAllTasks()
      } else {
        fmt.Println("master of this dungeon counts 10")
      }
      return nil
      },
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}