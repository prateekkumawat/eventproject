
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/reservation", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Reservation Service")
    })
    fmt.Println("Reservation service running on port 8080")
    http.ListenAndServe(":8080", nil)
}
