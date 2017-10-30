package ttts


import(
    "fmt"
)


func run_client(address string, port int) {
    fmt.Println("Starting Client...")
    fmt.Printf("Connecting to server %s:%d\n", address, port)
    return
}
