package ttts


import(
    "fmt"
    "os"
    "net"
)


func runServer(address string, port int) {
    fmt.Println("Starting Server...")

    server := fmt.Sprintf("%s:%d", address, port)

    fmt.Printf("Listening on %s:%d\n", address, port)
    ln, err := net.Listen("tcp", server)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to listen on: %s\n", server)
        os.Exit(1)
    }

    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Fprintf(os.Stderr, "Failed to accept connection.\n")
            os.Exit(1)
        } else {
            newGame(conn)
        }
    }
    return
}


func newGame(conn net.Conn) {
    var board = make([]byte, 13)
    for {
        _, err := conn.Read(board)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Failed to read board.\n")
            os.Exit(1)
        } else {
            fmt.Printf("Recieved board: %s\n", string(board))
        }
    }
    return
}
