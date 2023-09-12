package utils

var test string

func RenamedVariableKeep(renamed string) {
	test = renamed
}

func RenamedVariableTurn() string {
	return test
}
