package main

import (
	"net"
	"log"
	"io"
	"os"
	"fmt"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(err)
	}
	done := make(chan struct{})
	go func() {
		_, err := io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		if err != nil {
			log.Println(err)
		}
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	for {
		mustCopy1(conn, os.Stdin)
	}

	defer conn.Close()
	<-done

}
func mustCopy1(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
