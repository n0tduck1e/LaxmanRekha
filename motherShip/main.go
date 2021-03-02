package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type malware struct {
	Name    string
	Md5sum  string
	Sha1sum string
}

var samples = make(map[string]malware)

func addSample(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if _, ok := samples[r.PostForm["sha256sum"][0]]; ok {
		w.Write([]byte("Alreay have this sample.\n"))
		w.Write([]byte("Thank you for your time\n"))
		return
	}

	fmt.Println(r.PostForm)
	samples[r.PostForm["sha256sum"][0]] = malware{
		Name:    r.PostForm["name"][0],
		Md5sum:  r.PostForm["md5sum"][0],
		Sha1sum: r.PostForm["sha1sum"][0],
	}

	fmt.Println(samples)
	w.Write([]byte("Sample Added.\n"))
	w.Write([]byte("Thank you"))

}

func home(w http.ResponseWriter, r *http.Request) {
	temp, _ := template.ParseFiles("malwareSamples.html")
	temp.Execute(w, samples)
}

func main() {

	http.HandleFunc("/addSample", addSample)
	http.HandleFunc("/", home)
	log.Println("Starting listening on Port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
