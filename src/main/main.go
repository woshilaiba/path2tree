package main

import (
	"common"
	"flag"
)

func main() {
	s := flag.String("fn", "", "path of the file wanted")
	flag.Parse()

	common.Readfile(*s)
	common.Print_pathmeta()
}
