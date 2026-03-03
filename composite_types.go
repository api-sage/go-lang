package composite_types

import "fmt"

func ReverseString(strings []string) []string {
	out := strings[:0]
	fmt.Println(out)
	for _, item := range strings {
		if item != "" {
			out = append(out, item)
			fmt.Println(out)
		}
	}
	return out
}

func PopString(strings []string, removeItem string) (string, []string) {
	var out []string
	for _, item := range strings {
		if item != removeItem {
			out = append(out, item)
		}
	}

	return removeItem, out
}
