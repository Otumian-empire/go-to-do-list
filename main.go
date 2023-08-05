package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/otumian-empire/go-to-list/data"
	"github.com/otumian-empire/go-to-list/logic"
)

func clearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("Unsupported operating system")
	}
}

func DisplayInitialMessage() {
	clearScreen()
	fmt.Println("Todo List")
}

func main() {
	defer logic.CatchError()

	// Initialize an empty todo list
	list := data.NewTodoList()

	DisplayInitialMessage()

	for {
		logic.DisplayOptionsPrompt()
		option, err := logic.GetOption()
		clearScreen()

		if err != nil {
			fmt.Println(err)
		} else {
			switch option {
			case logic.ADD:
				logic.DisplayEnterTaskPrompt()
				task := logic.GetTask()

				list.Add(*data.NewTodo(task))
				clearScreen()

				logic.DisplaySuccessCue("added")
				logic.WaitForASec()
				clearScreen()
			case logic.READ:
				logic.DisplayTodoList(*list)
			case logic.EDIT:
				logic.DisplayEnterIndexPrompt()

				indexValue, indexErr := logic.GetIndex(len(list.Read()))

				if indexErr != nil {
					fmt.Println(indexErr)
				} else {
					logic.DisplayEnterTaskPrompt()
					_task := logic.GetTask()
					clearScreen()

					// get todo at index and update the task
					todo := list.ReadOne(indexValue)
					todo.Edit(_task)

					// update list
					list.Update(indexValue, todo)
					clearScreen()

					logic.DisplaySuccessCue("edited")
					logic.WaitForASec()
					clearScreen()
				}

			case logic.COMPLETE:
				logic.DisplayEnterIndexPrompt()
				index, err := logic.GetIndex(len(list.Read()))

				if err != nil {
					fmt.Println(err)
				} else {
					// get todo at index and update the isCompleted
					todo := list.ReadOne(index)
					todo.Complete()

					// update list
					list.Update(index, todo)
					clearScreen()

					logic.DisplaySuccessCue("completed")
					logic.WaitForASec()
					clearScreen()
				}

			case logic.DELETE:
				logic.DisplayEnterIndexPrompt()
				index, err := logic.GetIndex(len(list.Read()))

				if err != nil {
					fmt.Println(err)
				} else {
					list.Delete(index)
					clearScreen()

					logic.DisplaySuccessCue("deleted")
					logic.WaitForASec()
					clearScreen()
				}

			case logic.QUIT:
				clearScreen()
				fmt.Println("Program ended")
				os.Exit(1)
			default:
				fmt.Println("You have to select a valid option")
			}
		}
	}
}
