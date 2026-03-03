package main

import (
	"fmt"

	composite_types "github.com/api-sage/go-lang"
)

func main() {
	fmt.Println(composite_types.PopItemOnStack([]string{"Hello", "World", "Go", "Lang"}))
	fmt.Println(composite_types.PopString([]string{"Hello", "World", "Go", "Lang"}, "Go"))
	fmt.Println(composite_types.RemoveItemFromStack([]string{"Hello", "World", "Go", "Lang"}, 1))
	fmt.Println(composite_types.RemoveItemFromStackV2([]string{"Hello", "World", "Go", "Lang"}, 1))
	//composite_types.ReverseString([]string{"Hello", "World", "", "Go", "Lang"})
}
