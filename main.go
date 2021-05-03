package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"webAnalyzer/htmlanalyzer"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	if r.Method == "GET" {
		t.Execute(w, nil)
	} else {
		 r.ParseForm()
		url := r.Form["url"]
		if len(url)==0||url[0]==""{
			w.WriteHeader(http.StatusBadRequest)
			t.Execute(w, errors.New("please enter valid URL"))
		}else if err:=htmlanalyzer.ValidateUrl(url[0]);err!=nil{
			w.WriteHeader(http.StatusBadRequest)
			t.Execute(w, err)
		}else {
				if status,err:=htmlanalyzer.Analyser(url[0]);err!=nil{
					w.WriteHeader(http.StatusInternalServerError)
					err=t.Execute(w, err)
					fmt.Print(err)
				}else{
					err=t.Execute(w, status)
					fmt.Print(err)
				}
		}

	}
}

func main() {
	fmt.Println("Server is listing at port 8000")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
