package ttts


import(
    "fmt"
    "os"
    "net"
    "./tic-tac-toe-go/ttt"
)


func runClient(address string, port int) {
    fmt.Println("Starting Client...")

    server := fmt.Sprintf("%s:%d", address, port)

    fmt.Printf("Connecting to server %s\n", server)
    conn, err := net.Dial("tcp", server)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to connect to server: %s\n", server)
        conn.Close()
        os.Exit(1)
    }

    fmt.Printf("Sucess! Starting new game.\n")

    game := ttt.NewGame()
    playGame(conn, game, "X")
    conn.Close()
    return
}
