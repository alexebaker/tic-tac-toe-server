package ttts


import(
    "fmt"
)


func run_server(address string, port int) {
    fmt.Println("Starting Server...")
    fmt.Printf("Listening on %s:%d\n", address, port)
    return
}
