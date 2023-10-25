package main

import (
	"ascii/funcs"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func CheckFileHashing(fileName string) bool {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := make([]byte, 30*1024)
	sha256 := sha256.New()
	for {
		n, err := file.Read(buf)
		if n > 0 {
			_, err := sha256.Write(buf[:n])
			if err != nil {
				log.Fatal(err)
			}
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("Read %d bytes: %v", n, err)
			break
		}
	}

	sum1 := sha256.Sum(nil)
	sum := hex.EncodeToString(sum1)

	shadow := "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
	thinkertoy := "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3"
	standard := "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"

	if fileName == "standard.txt" && string(sum) == standard {
		return true
	} else if fileName == "thinkertoy.txt" && string(sum) == thinkertoy {
		return true
	} else if fileName == "shadow.txt" && string(sum) == shadow {
		return true
	}
	return false
}

func main() {
	args := os.Args[1:]
	if funcs.IsValid(args) {

		file := "standard.txt"
		if len(args) == 2 {
			file = args[1] + ".txt"
		}
		// fmt.Print(args[0])

		if !CheckFileHashing(file) {
			fmt.Println("ERROR: the file was changed!")
			return
		}

		if len(args[0]) == 0 {
			return
		}

		// fmt.Print(args[0])
		// check new line
		// atr := strings.ReplaceAll(args[0], "\n", "\\n")
		arr := strings.Split(args[0], "\\n")
		arr = funcs.SepNewLine(arr)

		counter := 0
		for _, w := range arr {
			if w == "" {
				counter++
			}
		}
		dlina := len(arr)
		if counter == len(arr) {
			dlina = len(arr) - 1
		}

		// for i := 0; i < len(arr); i++ {
		// 	fmt.Println(i, []byte(arr[i]))
		// 	// fmt.Printf("%v - '%v' \n", i, arr[i])
		// }

		text, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("Error: the file is incorrect")
		}

		mp := funcs.StoreInMap(string(text))

		for i := 0; i < dlina; i++ {
			if arr[i] == string("") {
				fmt.Println()
			} else {
				funcs.GetWord(arr[i], mp)
			}
		}
	}
}
