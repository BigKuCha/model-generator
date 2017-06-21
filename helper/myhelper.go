package helper

import "strings"

func SnakeCase2CamelCase(input string, pascal bool) string {
	names := strings.Split(input, "_")
	var n string
	for k, name := range names {
		if name == "id" {
			n += "ID"
		} else {
			if k == 0 && !pascal {
				n += name
			} else {
				n += strings.Title(name)
			}
		}
	}
	return n
}
