package main
import (
 "fmt"
 "net/http"
 "os"
)
func main() {
 var PORT string
 if PORT = os.Getenv("PORT"); PORT == "" {
  PORT = "3001"
 }
 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello World ")
 })
 http.ListenAndServe(":" + PORT, nil)
}
