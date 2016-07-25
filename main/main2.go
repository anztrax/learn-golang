package main

import (
	"container/list"
	"errors"
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
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
	theHardWayToReadFile := func() {
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

	theMoreSimpleWayToReadFile := func() {
		bs, err := ioutil.ReadFile("test.txt")
		if err != nil {
			return
		}

		str := string(bs)
		fmt.Println(str)
	}

	theMoreSimpleWayToWriteFile := func() {
		file, err := os.Create("test2.txt")
		if err != nil {
			return
		}
		defer file.Close()
		file.WriteString("test")
	}

	readDirectory := func() {
		dir, err := os.Open(".")
		if err != nil {
			return
		}
		defer dir.Close()

		fileInfos, err := dir.Readdir(-1)
		if err != nil {
			return
		}

		for _, fi := range fileInfos {
			fmt.Println("file name : ", fi.Name())
		}
	}

	walkTroughTheDirectory := func() {
		filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
			fmt.Println(path)
			return nil
		})
	}

	theHardWayToReadFile()
	theMoreSimpleWayToReadFile()
	theMoreSimpleWayToWriteFile()
	readDirectory()
	walkTroughTheDirectory()
}

func tryError() {
	err := errors.New("Error Message")
	fmt.Println("error : ", err)
}

type Person struct {
	Name string
	Age  int
}

//type of ByName is array of Person ?
type ByName []Person

//sort is an interface, so they need to implement 3 func to use sort capabilities
func (this ByName) Len() int {
	return len(this)
}
func (this ByName) Less(i, j int) bool {
	return this[i].Name < this[j].Name
}
func (this ByName) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func tryContainerAndSort() {
	//list here is linked list
	var x list.List

	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int)) //what is is mean ?
	}

	kids := []Person{
		{"Amebo", 9},
		{"Pico", 7},
	}

	sort.Sort(ByName(kids))
	fmt.Println(kids)
}

func hashAndCryptography() {
	getHash := func(filename string) (uint32, error) {
		bs, err := ioutil.ReadFile(filename) //bs = bytearray
		if err != nil {
			return 0, err
		}

		//crc32 is used to compare 2 file  it's highly likely (though not 100% certain) that the files are the same.
		h := crc32.NewIEEE()
		h.Write(bs)
		return h.Sum32(), nil //v value is uint32
	}

	h1, err := getHash("test3.txt")
	if err != nil {
		return
	}
	h2, err := getHash("test.txt")
	if err != nil {
		return
	}
	fmt.Println(h1, h2, h1 == h2)
}

func main() {
	trystringsPackage()
	tryosPackage()
	tryError()
	tryContainerAndSort()
	hashAndCryptography()
}
