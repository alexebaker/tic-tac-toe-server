package ttts


import(
    "os"
    "fmt"
    "net"
    "time"
    "strings"
    "./tic-tac-toe-go/ttt"
)


func playGame(conn net.Conn, game ttt.Game, mark string) {
    ai := ttt.NewImpossibleComputer()
    ai.SetMark(mark)

    hasWinner := false
    winner := ""


    for !game.IsOver() {
        game.ApplyMove(ai.Move(*game.Board()), ai.GetMark())

        if !sendBoard(conn, game) {
            return
        }

        if !game.IsOver() {
            if !readBoard(conn, game) {
                return
            }
        }
    }

    winner, hasWinner = game.Winner()
    msg := ""
    if hasWinner {
        msg = fmt.Sprintf("%s wins!", winner)
    } else {
        msg = "tie"
    }

    sendMessage(conn, msg)

    fmt.Printf("%s\nGame Over\n\n", msg)
    return
}


func sendMessage(conn net.Conn, msg string) bool {
    for attempts := 0; attempts < 10; attempts++ {
        _, err := conn.Write([]byte(msg))

        if err != nil {
            fmt.Fprintf(os.Stderr, "Failed to write message, trying again...\n")
        } else {
            return true
        }

        time.Sleep(time.Second)
    }

    fmt.Fprintf(os.Stderr, "Failed to write message after 10 attempts, closing connection.\n")
    conn.Close()
    return false
}


func readMessage(conn net.Conn, msg *string) bool {
    for attempts := 0; attempts < 10; attempts++ {
        bytes := make([]byte, 11)
        _, err := conn.Read(bytes)

        if err != nil {
            fmt.Fprintf(os.Stderr, "Failed to read message, trying again...\n")
        } else {
            *msg = string(bytes)
            return true
        }

        time.Sleep(time.Second)
    }

    fmt.Fprintf(os.Stderr, "Failed to read message after 10 attempts, closing connection.\n")
    conn.Close()
    return false
}


func sendBoard(conn net.Conn, game ttt.Game) bool {
    board := board2str(game)
    fmt.Printf("Sending board: %s\n", board)
    return sendMessage(conn, board)
}


func readBoard(conn net.Conn, game ttt.Game) bool {
    board := ""
    success := readMessage(conn, &board)
    fmt.Printf("Recieved board: %s\n", string(board))
    str2board(strings.ToUpper(string(board)), game)
    return success
}


func convertFromNetBoard(str string) string {
    str = fmt.Sprintf("%s%s%s", str[0:3], str[4:7], str[8:11])
    str = strings.Replace(str, "-", " ", -1)
    return str
}


func convertToNetBoard(str string) string {
    str = strings.Replace(str, " ", "-", -1)
    str = fmt.Sprintf("%s|%s|%s", str[0:3], str[3:6], str[6:9])
    return str
}


func str2board(str string, game ttt.Game) {
    str = convertFromNetBoard(str)
    for i,_ := range game.Board().Spaces() {
        game.ApplyMove(i, string(str[i]))
    }
    return
}


func board2str(game ttt.Game) string {
    str := make([]string, 9)
    for i,m := range game.Board().Spaces() {
        str[i] = m
    }
    return convertToNetBoard(strings.Join(str, ""))
}
