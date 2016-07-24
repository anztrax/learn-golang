package main

import (
	"fmt"
	"os"
	"strings"
)

//this main2 is used for test main core package
func trystringsPackage() {
	fmt.Println(
		strings.Contains("test", "st"),
		strings.Count("test", "t"),
		strings.HasPrefix("test", "te"),
		strings.HasSuffix("test", "st"),
		strings.Index("asep", "e"),
		strings.Join([]string{"a", "b"}, "-"),
		strings.Repeat("a", 5),
		strings.Replace("aaaa", "a", "b", 2),
		strings.Split("a-b-c-d-e", "-"),
		strings.ToLower("TEST"),
		strings.ToUpper("test"),
	)

	arr1 := []byte("test")
	str1 := string([]byte{'t', 'e', 's', 't'})
	fmt.Println("the arr1 : ", arr1)
	fmt.Println("the str1 : ", str1)
}

func tryosPackage() {
	theHardWay := func() {
		file, err := os.Open("test.txt")
		if err != nil {
			//handle error here
			return
		}
		defer file.Close()

		//get the file size
		stat, err := file.Stat() //get the stat
		if err != nil {
			return
		}

		//read the file
		bs := make([]byte, stat.Size()) //make byte[] from file size which is from stat
		_, err = file.Read(bs)

		if err != nil {
			return
		}

		str := string(bs)
		fmt.Println(str)
	}
	theHardWay()
}

func main() {
	trystringsPackage()
	tryosPackage()
}
