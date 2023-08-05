package logic

import (
	"fmt"

	"github.com/otumian-empire/go-to-list/data"
)

func DisplayOptionsPrompt() {
	fmt.Println("1.  Add")
	fmt.Println("2.  Read")
	fmt.Println("3.  Edit")
	fmt.Println("4.  Delete")
	fmt.Println("5.  Complete")
	fmt.Println("12. Quit")
	fmt.Print("Choose from the options above: ")
}

func DisplayEnterTaskPrompt() {
	fmt.Print("Enter Task: ")
}

func DisplayEnterIndexPrompt() {
	fmt.Print("Enter Index: ")
}

func DisplaySuccessCue(cue string) {
	fmt.Printf("Task %v successfully!!😍️\n", cue)
}

func DisplayTodoList(list data.TodoList) {
	fmt.Println("TODO List")
	for index, todo := range list.Read() {
		emoji := "🎁️"

		if todo.GetIsCompleted() {
			emoji = "✅️"
		}

		fmt.Printf("%v %v (%v)\n", index, todo.GetTask(), emoji)
	}

	fmt.Println("TODO List 😉️!!!")
}
