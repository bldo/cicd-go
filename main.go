package main
import (
    "fmt"
    "net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    line := "***********************"
    output := "   Hello ByteDancer!"
    fmt.Fprintln(w, line)
    fmt.Fprintln(w, output)
    fmt.Fprintln(w, line)
}

func main() {
    http.HandleFunc("/", IndexHandler)
    fmt.Println("server start..")
    err := http.ListenAndServe(":80", nil)
    if err != nil {
        fmt.Printf("failed to start.. %v", err)
    }
}
