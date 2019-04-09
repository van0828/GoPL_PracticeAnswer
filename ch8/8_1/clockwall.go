package main

import (
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	for _, v := range os.Args[1:] {
		keyValue := strings.Split(v, "=")
		go conTcp(keyValue[1])
	}
	for {
		time.Sleep(1 * time.Second)
	}
}

func conTcp(url string) {
	conn, err := net.Dial("tcp", url)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustcopy(os.Stdout, conn)
}

func mustcopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
