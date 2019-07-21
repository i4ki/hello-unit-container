package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"nginx/unit"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Unit container example!\n")
	fmt.Fprintf(w, "My PID: %d\n", os.Getpid())
	fmt.Fprintf(w, "rootfs:\n")

	files, err := ioutil.ReadDir("/")
	if err != nil {
		fmt.Fprintf(w, "Failed to read /: %s\n", err)
		return
	}

	for _, f := range files {
		fmt.Fprintf(w, "%s ", f.Name())
	}

	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "UID: %d\n", os.Getuid())
	fmt.Fprintf(w, "GID: %d\n\n", os.Getgid())
	fmt.Fprintf(w, "Environment variables: \n")

	for _, env := range os.Environ() {
		fmt.Fprintf(w, "%s\n", env)
	}
}

func main() {
	fmt.Println("Unit Example started")
	http.HandleFunc("/", handler)
	log.Fatal(unit.ListenAndServe(":8080", nil))
}
