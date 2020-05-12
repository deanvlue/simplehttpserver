/*
* serves a simpe static file or directory in go
* Usage:
*   -p="8100"   : Puesto a servir
    -d="."      : Directorio a servir
* Nevgar a http://localhost:8100 te da la lista de archivos
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"os"
	"os/signal"
	"syscall"
	"github.com/rs/cors"
)

// AppVersion This is the last modification app version
const AppVersion = "2020-05-21:12:58"

// SetupCloseHandler Handles the exit of the App
func SetupCloseHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Println("\r Ctrl+C  pressed in terminal")
		log.Println("I'll be back")
		os.Exit(0)
	}()
}

func PrintVersion() {
	fmt.Printf("SimpleHttpServer %v\n", AppVersion)
	fmt.Println("Serves a directory")
	fmt.Println("Copyright (C) 2019 Carlos Munoz Inc.")
	fmt.Println("This is free software; see the source for copying conditions.  There is NO")
	fmt.Println("warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE")
}

func main() {

	SetupCloseHandler()

	mux := http.NewServeMux()

	port := flag.String("p", "8100", "puerto de servicio")
	directory := flag.String("d", ".", "el directorio a servir")
	version := flag.Bool("v", false, "version de app")
	flag.Parse()

	if *version {
		PrintVersion()
		os.Exit(0)
	}


	//http.Handle("/", http.FileServer(http.Dir(*directory)))
	mux.Handle("/", http.FileServer(http.Dir(*directory)))
//	http.HandleFunc("/", serveFiles)
	handler := cors.Default().Handler(mux)

	log.Printf("Sirviendo %s en puerto %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, handler))
}

