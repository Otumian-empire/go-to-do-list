package logic

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// option constants
var (
	ADD      = 1
	READ     = 2
	EDIT     = 3
	DELETE   = 4
	COMPLETE = 5
	QUIT     = 12
)

var (
	WAIT_TIME = 1 // sec
)

// error messages
var (
	GET_OPTION_ERROR_MESSAGE = "Enter a valid option: 1, 2, 3, 4, 5 or 12"
	GET_INDEX_ERROR_MESSAGE  = "Enter a valid (numeric) index from the TODO List"
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
