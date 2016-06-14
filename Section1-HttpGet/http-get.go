package main

// step 1 - use http.Get to send a request and
// print a response, handling the error
import (
        "fmt"
        "log"
        "net/http"
)

func main() {
  res, err := http.Get("https://golang.org")
  if err != nil {
        log.Fatal(err)
      }
      fmt.Println(res.Status) 
}
