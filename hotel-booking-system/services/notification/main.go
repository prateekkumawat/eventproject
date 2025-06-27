
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/notify", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Notification sent")
    })
    fmt.Println("Notification service running on port 8080")
    http.ListenAndServe(":8080", nil)
}
