package main

import (
	"io"
	"net/http"
	_ "net/http/pprof"
)

// go build -gcflags -S main.go
// go build -o main main.go ; go tool objdump -s main.main main
// godbolt.org
// go build -o main main.go -> ./main -> go tool pprof main http://127.0.0.1:8080/debug/pprof/profile ->  ab -k -c 8 -n 100000 "http://127.0.0.1:8080/"
// go build -o main main.go ; objdump -S main

//go:noinline
func sayHello(writer io.Writer) {
	writer.Write([]byte("Hello world!"))
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		sayHello(writer)
	})
	http.ListenAndServe(":8080", nil)
}
