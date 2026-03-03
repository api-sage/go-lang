package main

import (
	"fmt"

	composite_types "github.com/api-sage/go-lang"
)

func main() {
	fmt.Println(composite_types.PopItemOnStack([]string{"Hello", "World", "Go", "Lang"}))
	fmt.Println(composite_types.PopString([]string{"Hello", "World", "Go", "Lang"}, "name"))
	//composite_types.ReverseString([]string{"Hello", "World", "", "Go", "Lang"})
}
