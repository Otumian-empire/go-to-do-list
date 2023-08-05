package logic

// option constants
const (
	ADD      = 1
	READ     = 2
	EDIT     = 3
	DELETE   = 4
	COMPLETE = 5
	QUIT     = 12
)

// time to wait
const (
	WAIT_TIME = 1 // sec
)

// error messages
const (
	GET_OPTION_ERROR_MESSAGE = "enter a valid option: 1, 2, 3, 4, 5 or 12"
	GET_INDEX_ERROR_MESSAGE  = "enter a valid (numeric) index from the TODO List"
)
