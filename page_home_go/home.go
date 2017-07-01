package main

import (
    "html/template"
    "net/http"
    "flag"
    "io"
    "os/exec"
    "log"
)

// Command-line flags.
var httpAddr = flag.String("http", ":8080", "Listen address")

var words []string

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
	    t, _ := template.ParseFiles("index.gtpl")
	    t.Execute(w, nil)
	} else {
		r.ParseForm()
		if r.Form["cv"] != nil {
			io.WriteString(w, "CV INSTANCE HAS BEEN CREATED")
			runDocker("cv")
		}
	}
}

func runDocker(instance string) {
	cmd := exec.Command("docker", "start", "website_" + instance + "_1")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()
	http.HandleFunc("/", home)
	http.ListenAndServe(*httpAddr, nil)
}