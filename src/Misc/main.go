package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
)

func main() {
	// aByteSlice := []byte("Meron Hayle\n")
	// ioutil.WriteFile("temp.txt", aByteSlice, 0644)
	// f, err := os.Open("temp.txt")
	// if err != nil {
	// 	log.Fatal("Can not read file", err)
	// }
	// // f.Write(aByteSlice)
	// defer f.Close()

	// antotherByteSlice := make([]byte, 100)

	// n, err := f.Read(antotherByteSlice)
	// fmt.Printf("Read %d bytes: %s", n, antotherByteSlice)
	file, err := os.Open("temp.txt")
	if err != nil {
		log.Fatal("Can not open file", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		//Reads a line
		line := scanner.Text()

		if scanner.Err() != nil{
			log.Fatalln("Error Reading file: ", scanner.Err())
		}
		fmt.Println(line)
	}

	// fmt.Println("The Text: ", string(byteNew))
}
