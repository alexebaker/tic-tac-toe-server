package ttts


import(
    "os"
    "fmt"
    "net"
    "time"
    "strings"
    //"github.com/jneander/tic-tac-toe-go/ttt"
)


func playGame(conn net.Conn, board string, mark string) {
    //hasWinner := false

    for {
        if makeMove(&board, &mark) {
            return
            }

        if !sendBoard(&conn, &board) {
            conn.Close()
            return
        }

        if !readBoard(&conn, &board) {
            conn.Close()
            return
        }
    }
    return
}


func makeMove(board *string, mark *string) bool {
    for i := 0; i < len(*board); i++ {
        if (*board)[i] == ' ' {
            tmp := []rune(*board)
            tmp[i] = []rune(*mark)[0]
            *board = string(tmp)
            return false
        }
    }
    return true
}


func sendBoard(conn *net.Conn, board *string) bool {
    convertToNetBoard(board)

    for attempts := 0; attempts < 10; attempts++ {
        _, err := (*conn).Write([]byte(*board))

        if err != nil {
            fmt.Fprintf(os.Stderr, "Failed to write board, trying again...\n")
        } else {
            return true
        }

        time.Sleep(time.Second)
    }

    fmt.Fprintf(os.Stderr, "Failed to write board after 10 attempts, ending game.\n")
    return false
}


func readBoard(conn *net.Conn, board *string) bool {
    for attempts := 0; attempts < 10; attempts++ {
        newBoard := make([]byte, 13)
        _, err := (*conn).Read(newBoard)

        if err != nil {
            fmt.Fprintf(os.Stderr, "Failed to read board, trying again...\n")
        } else {
            fmt.Printf("Recieved board: %s\n", string(newBoard))
            *board = string(newBoard)
            convertFromNetBoard(board)
            return true
        }

        time.Sleep(time.Second)
    }

    fmt.Fprintf(os.Stderr, "Failed to read board after 10 attempts, ending game.\n")
    return false
}


func convertFromNetBoard(board *string) {
    *board = fmt.Sprintf("%s%s%s", (*board)[0:3], (*board)[4:7], (*board)[8:11])
    *board = strings.Replace(*board, "-", " ", -1)
    return
}


func convertToNetBoard(board *string) {
    *board = strings.Replace(*board, " ", "-", -1)
    *board = fmt.Sprintf("%s|%s|%s\n", (*board)[0:3], (*board)[3:6], (*board)[6:9])
    return
}
