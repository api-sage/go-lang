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
