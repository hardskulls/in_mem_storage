package main

import (
	"in_mem_storage/app"
	"os"
	"strconv"
)

func main() {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	srv := app.New(port)

	err := srv.Run()
	if err != nil {
		panic(err)
	}
}
