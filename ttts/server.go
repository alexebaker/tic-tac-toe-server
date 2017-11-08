package ttts


import(
    "fmt"
    "os"
    "net"
    "./tic-tac-toe-go/ttt"
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
            conn.Close()
        } else {
            fmt.Printf("Connection recieved, starting new game...\n")
            game := ttt.NewGame()

            if readBoard(conn, game) {
                playGame(conn, game, "O")
            } else {
                fmt.Fprintf(os.Stderr, "Failed to read board.\n")
                conn.Close()
            }
        }
        //conn.Close()
    }
    return
}
