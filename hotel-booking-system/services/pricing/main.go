
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/pricing", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Pricing Service")
    })
    fmt.Println("Pricing service running on port 8080")
    http.ListenAndServe(":8080", nil)
}
