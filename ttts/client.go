/******************************************************
 * The file has the logic for running in client mode. *
 ******************************************************/

package ttts


import(
    "fmt"
    "os"
    "net"
    "time"
    "./tic-tac-toe-go/ttt" // https://github.com/jneander/tic-tac-toe-go
)


func runClient(address string, port int, num int) {
    fmt.Println("Starting Client...")

    server := fmt.Sprintf("%s:%d", address, port)

    for i := 0; i < num; i++ {
        time.Sleep(time.Second)

        fmt.Printf("Connecting to server %s\n", server)
        conn, err := net.Dial("tcp", server)

        if err != nil {
            fmt.Fprintf(os.Stderr, "Failed to connect to server: %s\n", server)
            conn.Close()
            os.Exit(1)
        }

        fmt.Printf("Sucess! Starting new game.\n")

        tmp := ""
        game := ttt.NewGame()
        playGame(conn, game, "X")
        readMessage(conn, &tmp)
        conn.Close()
    }
    return
}
