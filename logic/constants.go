package logic

// option constants
var (
	ADD      = 1
	READ     = 2
	EDIT     = 3
	DELETE   = 4
	COMPLETE = 5
	QUIT     = 12
)

// time to wait
var (
	WAIT_TIME = 1 // sec
)

// error messages
var (
	GET_OPTION_ERROR_MESSAGE = "Enter a valid option: 1, 2, 3, 4, 5 or 12"
	GET_INDEX_ERROR_MESSAGE  = "Enter a valid (numeric) index from the TODO List"
)
