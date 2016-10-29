package main

import (
  "fmt"
  "net/http"
  "text/template"

  "github.com/jteeuwen/go-pkg-optarg"
)

type Page struct {
  Title string
  Count int
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
  page := Page{"Hello World.", 1}
  tmpl, err := template.ParseFiles("template.html") // ParseFilesを使う
  if err != nil {
    panic(err)
  }

  err = tmpl.Execute(w, page)
  if err != nil {
    panic(err)
  }
}

func main() {
  optarg.Add("p", "port", "Specify port number", "8080")
  var portNum string = "8080"
  for opt := range optarg.Parse() {
    switch opt.ShortName {
    case "p":
      portNum = opt.String()
    }
  }

  http.HandleFunc("/", viewHandler) // hello
  http.ListenAndServe(":"+portNum, nil)
}
