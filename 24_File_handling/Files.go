package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
func main() {
	f, err := os.Open("exmp.txt")
	if err != nil {
		// log the error
		panic(err)
	}
	fileInfo, err := f.Stat()
	if err != nil {
		// log the error
		panic(err)
	}
	fmt.Println("file Nmae:", fileInfo.Name())
	fmt.Println("file or folder:", fileInfo.IsDir())
	fmt.Println("file size:", fileInfo.Size())
	fmt.Println("file permission:", fileInfo.Mode())
	fmt.Println("file modified at:", fileInfo.ModTime())

}   */

/*
// Read file
func main() {
	f, err := os.Open("exmp.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	buf := make([]byte, 15)

	d, err := f.Read(buf)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(buf); i++ {

		fmt.Println("data", d, string(buf[i]))
	}
}   */

/*
// other ways -  only for small files

func main() {
	data, err := os.ReadFile("exmp.txt")
	if err != nil {
		panic(err)

	}
	fmt.Println(string(data))

}  */

/*
// read Folder

func main() {
	dir, err := os.Open("../")
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	fileInfo, err := dir.ReadDir(-1)

	for _, fi := range fileInfo {

		fmt.Println(fi.Name(), fi.IsDir())
	}

}  */

/*
// Create a file
func main() {

	f, err := os.Create("exmp2.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()
	f.WriteString("Hi go")
	f.WriteString("It is a very nice language")
}

//other ways
func main() {

	f, err := os.Create("exmp2.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()
	bytes := []byte("Hello Go lang")
	f.Write(bytes)
} */

// Read and write to anpother file(Streaming fashion)

func main() {
	sourceFile, err := os.Open("exmp.txt")
	if err != nil {
		panic(err)
	}

	defer sourceFile.Close()

	destFile, err := os.Create("exmp2.txt")
	if err != nil {
		panic(err)
	}
	defer destFile.Close()

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destFile)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}
		e := writer.WriteByte(b)
		if e != nil {
			panic(e)
		}
	}

	writer.Flush()

	fmt.Println("Written to new file successfully")

}

/*
//Delete a File

func main() {

	err := os.Remove("exmp2.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("File deleted successfully")
}
*/
