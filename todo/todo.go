package main

import (
  "fmt"
  "log"
  "os"
  "time"
  "encoding/json"
  "bufio"
  "strconv"

  "github.com/urfave/cli/v2"
)
// --- 1. записать в файл ---  Done
// --- 2. выводить на экран красивенько все задачи --- Done
// --- 2.1. форматировать дату нормально --- Done
// 3. отметить как выполненное или проваленное: прочитать файл, записать всё в список структур, найти нужную, изменить, перезаписать всё скопом в файл 
// 4. добавить сохранение даты выполнения/провала
// 4. добавить причину выполнения/провала (тэги?) и как можно улучшить выполнение в следующий раз
// 5. Id ???

type Task struct {
  Id int
  Text string
  Status string
  Start_date string
  End_date string
  WhatHelped string
  WhatMessedUp string
}

func ReadTaskFile() []Task {
  /*
  ** Reading all tasks to a slice of Task structures
  */
  file, _ := os.Open("todo/tasks.json")
  scanner := bufio.NewScanner(file)
  taskList := []Task{}
  
  for scanner.Scan() {
    textFromFile := scanner.Text()
    task := Task{}
    err := json.Unmarshal([]byte(textFromFile), &task)
    if err != nil {
      fmt.Println(err)
    }
    taskList = append(taskList, task)
    fmt.Println(task.Status)
  }
  return taskList
}

func UpdateStatus(id int, err error)  {
  
  taskList := ReadTaskFile()
  fmt.Println("Count tasks:")
  fmt.Println(len(taskList))

  /*
  ** First step: Find exact struct by given id
  */

  /*
  ** Second step: rewrite status in any structure
  */
  taskList[id].Status = "Success"

  /*
  ** Third step: rewrite the whole file
  */
  err = os.Remove("todo/tasks.json")
  if err != nil {
    return
  }
  for _, task := range taskList {
    b, err := json.Marshal(task)
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
    fmt.Println(task.Id, "  ", task.Start_date, "  ", task.Text, "  ", task.Status)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "shouldn't see an error scanning a string")
	}

}

func SaveTask(taskText string) {
  var new_task Task
  new_task.Id = CountId()
  new_task.Text = taskText
  new_task.Status = "in process"
  t := time.Now().Format("02-01-2006")
  fmt.Println(t)
  //new_task.Start_date = t.String()
  new_task.Start_date = t
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

func CountId() int {
  taskList := ReadTaskFile()
  lastId := taskList[len(taskList) - 1].Id
  return (lastId + 1)
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
      } else if name == "ls" {
        ShowAllTasks()
      } else if name == "us" {
        UpdateStatus(strconv.Atoi(c.Args().Get(2)))
      }else {
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