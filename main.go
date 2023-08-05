package main

import (
	"fmt"
	"os"

	"github.com/otumian-empire/go-to-list/data"
	"github.com/otumian-empire/go-to-list/logic"
)

func DisplayInitialMessage() {
	logic.ClearScreen()
	fmt.Println("Todo List")
}

func AddTodoItem(list *data.TodoList) {
	logic.DisplayEnterTaskPrompt()
	task := logic.GetTask()

	list.Add(*data.NewTodo(task))
	logic.ClearScreen()

	logic.DisplaySuccessCue("added")
	logic.WaitForASec()
	logic.ClearScreen()
}

func EditTodoItem(list *data.TodoList) {
	logic.DisplayEnterIndexPrompt()

	indexValue, indexErr := logic.GetIndex(len(list.Read()))

	if indexErr != nil {
		fmt.Println(indexErr)
	} else {
		logic.DisplayEnterTaskPrompt()
		_task := logic.GetTask()
		logic.ClearScreen()

		// get todo at index and update the task
		todo := list.ReadOne(indexValue)
		todo.Edit(_task)

		// update list
		list.Update(indexValue, todo)
		logic.ClearScreen()

		logic.DisplaySuccessCue("edited")
		logic.WaitForASec()
		logic.ClearScreen()
	}
}

func CompleteTodoItem(list *data.TodoList) {
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
		logic.ClearScreen()

		logic.DisplaySuccessCue("completed")
		logic.WaitForASec()
		logic.ClearScreen()
	}
}

func DeleteTodoItem(list *data.TodoList) {
	logic.DisplayEnterIndexPrompt()
	index, err := logic.GetIndex(len(list.Read()))

	if err != nil {
		fmt.Println(err)
	} else {
		list.Delete(index)
		logic.ClearScreen()

		logic.DisplaySuccessCue("deleted")
		logic.WaitForASec()
		logic.ClearScreen()
	}
}

func Quit() {
	logic.ClearScreen()
	fmt.Println("Program ended")
	os.Exit(1)
}

func main() {
	defer logic.CatchError()

	// Initialize an empty todo list
	list := data.NewTodoList()

	DisplayInitialMessage()

	for {
		logic.DisplayOptionsPrompt()
		option, err := logic.GetOption()
		logic.ClearScreen()

		if err != nil {
			fmt.Println(err)
		} else {
			switch option {
			case logic.ADD:
				AddTodoItem(list)
			case logic.READ:
				logic.DisplayTodoList(*list)
			case logic.EDIT:
				EditTodoItem(list)
			case logic.COMPLETE:
				CompleteTodoItem(list)
			case logic.DELETE:
				DeleteTodoItem(list)
			case logic.QUIT:
				Quit()
			default:
				fmt.Println("You have to select a valid option")
			}
		}
	}
}
