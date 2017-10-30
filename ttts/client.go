package ttts


import(
    "fmt"
    "os"
    "net"
    "time"
)


func runClient(address string, port int) {
    fmt.Println("Starting Client...")

    server := fmt.Sprintf("%s:%d", address, port)

    fmt.Printf("Connecting to server %s\n", server)
    conn, err := net.Dial("tcp", server)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to connect to server: %s\n", server)
        os.Exit(1)
    } else {
        fmt.Printf("Sucess!\n")
    }

    var board string = "X--|---|---\n"
    for {
        fmt.Printf("Sending board: %s", board)
        _, err := conn.Write([]byte(board))
        if err != nil {
            fmt.Fprintf(os.Stderr, "Failed to write board (%s) to server: %s\n", board, server)
            os.Exit(1)
        } else {
            fmt.Printf("Sucess!\n")
        }
        time.Sleep(time.Second * 1)
    }
    return
}
