package logic

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func GetOption() (int, error) {
	var option int
	_, _ = fmt.Scan(&option)

	switch option {
	case QUIT, ADD, READ, EDIT, DELETE, COMPLETE:
		return option, nil
	default:
		return 0, fmt.Errorf(GET_OPTION_ERROR_MESSAGE)
	}
}

/*
// wasn't reading line
func GetTask() string {
	var task string
	fmt.Scanln(&task)

	return task
} */

func GetTask() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()         // Read a line
	task := scanner.Text() // Retrieve the line as text

	return task
}

func GetIndex(size int) (int, error) {
	var index int
	_, err := fmt.Scan(&index)

	if err != nil || index >= size || index < 0 {
		return 0, fmt.Errorf(GET_INDEX_ERROR_MESSAGE)
	}

	return index, nil
}

func CatchError() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}

func WaitForASec() {
	time.Sleep(time.Second * time.Duration(WAIT_TIME))
}

func ClearScreen() {
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
