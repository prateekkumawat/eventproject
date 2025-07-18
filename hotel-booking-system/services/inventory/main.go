
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/inventory", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Inventory Service")
    })
    fmt.Println("Inventory service running on port 8080")
    http.ListenAndServe(":8080", nil)
}
