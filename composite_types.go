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
	hasString := false
	for _, item := range strings {
		if item != removeItem {
			out = append(out, item)
		} else {
			hasString = true
		}
	}
	fmt.Println("Underlaying array capacity :: ", cap(out))

	if !hasString {
		return removeItem + " does not exist in slice", strings
	}

	return removeItem, out
}

func PopItemOnStack(stack []string) (string, []string) {
	if len(stack) == 0 {
		return "", stack
	}
	stackSlice := stack[:len(stack)-1]
	fmt.Println("Underlaying array capacity :: ", cap(stackSlice))
	return stack[len(stack)-1], stackSlice
}

func RemoveItemFromStack(stack []string, itemIndex int) (string, []string) {
	copy(stack[itemIndex:], stack[itemIndex+1:])
	updatedSlice := stack[:len(stack)-1]
	return stack[itemIndex], updatedSlice
}
