package library

import (
	"strings"
)

func GenerateClassCode(name string) string {

	nameSlice := strings.Split(name, " ")

	sliceLength := len(nameSlice)
	var code string

	for i, v := range nameSlice {

		if i == sliceLength-1 {

			code += "-"

		}

		code += strings.ToUpper(string((v[0])))

	}

	return code

}
