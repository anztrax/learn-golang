package main

import (
	"flag"
	"fmt"
	"math/rand"
)

//how to execute this is using this command : (go build tryWithCommandLine.go)

func main() {
	maxp := flag.Int("max", 6, "the max value") //flag is used to get argument from command line
	//parse
	flag.Parse()
	arguments := flag.Args()
	fmt.Println(arguments)
	fmt.Println(rand.Intn(*maxp)) //generate number 0 and max

}
