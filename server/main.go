package main

import (
	"github.com/hongruiqi/go-rtmp"
)

type App struct {

}

func main() {
	fs := rtmp.NewFlashServer()
	fs.Handle("flvplayback/", App{})
	server := rtmp.NewServer(fs)
	err := server.ListenAndServe(":1935")
	if err != nil {
		panic(err)
	}
}
