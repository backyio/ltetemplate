package utils

import "fmt"

var (
	// CErrorNoVariableParameter is common error message when variable parameter not set
	CErrorNoVariableParameter = fmt.Errorf("Empty variable parameter")
	// CErrorNoVariableParameter is common error message when internal instance pointer not valid
	CErrorEmptyInstanceVariable = fmt.Errorf("Empty instance variable")
	// CErrorIDAlreadyExists is common error message when added id already exists in the list
	CErrorIDAlreadyExists = fmt.Errorf("Id already exists")
	// CErrorIDNotExists is common error message when added id not exists in the list
	CErrorIDNotExists = fmt.Errorf("Id not exists")
	// // CErrorInvalidParameterType is common error message when parameter type is not valid
	CErrorInvalidParameterType = fmt.Errorf("Parameter type is not valid")
)
