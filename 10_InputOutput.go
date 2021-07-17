package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"syscall"
)

/*
- import io, ioutil, bufio, fmt, os, path, path/filepath
- io.Writer || io.Reader == for consequential access
- io.[Reader || writer Seeker ].At == for specified access
- io.Copy
- io.WriterTo || io.ReadFrom implemented vy the interface ==> io.Copy with no additionall memory allocation
- Closer, ByteReader, ByteWriter == return by to the source
- Multi [Reader, Writer, Limit] == cat file1 file2 file3
- buffering of the Read || Write == bufio.*
*/

func main() {
	path := "test"
	O_RDWR := syscall.O_RDWR
	var file *os.File // файловый дескриптор в Go
	file, err := os.OpenFile(path, O_RDWR, 0644)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File do not exist")
		}
		// другие ошибки, например нет прав
	}
	/* read with offset */
	N := 42 // мы заранее знаем сколько хотим прочитать
	buf := make([]byte, N)
	// подготавливаем буфер нужного размера
	file2, _ := os.Open(path) // открываем файл
	offset := 0
	for offset < N {
		read, err := file2.Read(buf[offset:])
		offset += read
		if err == io.EOF {
			// что если не дочитали ?
			break
		}
		if err != nil {
			log.Panicf("failed to read: %v", err)
		}
	}

	/* read all */

	// b := make([]byte, 1024*1024)
	// file3, _ := os.Open(path)
	// read, err := io.ReadFull(file3, b) // содержит цикл внутри

	b, err := ioutil.ReadFile(path)

	fmt.Println(buf, b)
	defer file.Close()
}
