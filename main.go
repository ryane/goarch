package main

import (
	"flag"
	"fmt"
	"runtime"
)

func main() {
	all := flag.Bool("a", false, "print <os>_<arch>")
	os := flag.Bool("os", false, "print <os>")
	flag.Parse()

	goarch := runtime.GOARCH
	goos := runtime.GOOS

	if *all {
		fmt.Printf("%s_%s\n", goos, goarch)
	} else if *os {
		fmt.Println(goos)
	} else {
		fmt.Println(goarch)
	}
}
