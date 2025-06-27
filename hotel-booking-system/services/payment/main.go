
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/pay", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Payment processed")
    })
    fmt.Println("Payment service running on port 8080")
    http.ListenAndServe(":8080", nil)
}
