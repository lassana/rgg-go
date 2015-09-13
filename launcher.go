package main

import (
	"fmt"
	"./platform"
)

func main() {
	fmt.Println(platform.GetFilenameFor(platform.DOS))
}
