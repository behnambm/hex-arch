package utils

import "runtime"

func GetCallerInfo() string {
	// Get the PC (program counter) for the caller
	pc, _, _, _ := runtime.Caller(1)
	// Get the function from the PC
	function := runtime.FuncForPC(pc)
	// Retrieve the full function name
	fullName := function.Name()
	// Split the full function name into package name and function name
	dotIndex := len(fullName)
	for i := len(fullName) - 1; i >= 0; i-- {
		if fullName[i] == '.' {
			dotIndex = i
			break
		}
	}
	packageName := fullName[:dotIndex]
	funcName := fullName[dotIndex+1:]
	return packageName + "." + funcName
}
