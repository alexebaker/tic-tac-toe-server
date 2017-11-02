package ttts


import(
    "os"
    "fmt"
    "net"
    "time"
    "strings"
)


func playGame(conn net.Conn, board string, mark string) {
    hasWinner := false
    winner := ""

    for !hasWinner {
        hasWinner, winner = makeMove(&board, &mark)

        if !sendBoard(&conn, &board) {
            conn.Close()
            return
        }

        if hasWinner {
            for attempts := 0; attempts < 10; attempts++ {
                _, err := conn.Write([]byte(winner))

                if err != nil {
                    fmt.Fprintf(os.Stderr, "Failed to write winner, trying again...\n")
                } else {
                    break
                }

                time.Sleep(time.Second)
            }
        }

        if !readBoard(&conn, &board) {
            conn.Close()
            return
        }
    }

    fmt.Println("Game Over")
    return
}


func makeMove(board *string, mark *string) (bool, string) {
    for i := 0; i < len(*board); i++ {
        if (*board)[i] == ' ' {
            tmp := []rune(*board)
            tmp[i] = []rune(*mark)[0]
            *board = string(tmp)
            return false, ""
        }
    }
    return true, "tie\n"
}


func sendBoard(conn *net.Conn, board *string) bool {
    convertToNetBoard(board)

    for attempts := 0; attempts < 10; attempts++ {
        fmt.Printf("Sending board: %s", *board)
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
        newBoard := make([]byte, 12)
        _, err := (*conn).Read(newBoard)

        if err != nil {
            fmt.Fprintf(os.Stderr, "Failed to read board, trying again...\n")
        } else {
            fmt.Printf("Recieved board: %s", string(newBoard))
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

/********************************************************************
 * Tic Tac Toe code from https://github.com/jneander/tic-tac-toe-go *
 ********************************************************************/


 type Board struct {
  spaces []string
}

func ( b Board ) Blank() string {
  return " "
}

func ( b *Board ) Spaces() []string {
  dup := make( []string, len( b.spaces ) )
  copy( dup, b.spaces )
  return dup
}

func ( b *Board ) Mark( pos int, mark string ) {
  if pos >= 0 && pos < len( b.spaces ) {
    b.spaces[ pos ] = mark
  }
}

func ( b Board ) SpacesWithMark( mark string ) []int {
  count, result := 0, make( []int, len(b.Spaces()) )
  for i,v := range b.Spaces() {
    if v == mark {
      result[count] = i
      count++
    }
  }
  return result[:count]
}

func ( b *Board ) Reset() {
  setBoard( b )
}

func NewBoard() *Board {
  b := new( Board )
  setBoard( b )
  return b
}

func setBoard( b *Board ) {
  b.spaces = make( []string, 9 )
  for i := range b.spaces {
    b.spaces[ i ] = " "
  }
}


type Game interface {
  Board() *Board
  IsOver() bool
  IsValidMove( int ) bool
  ApplyMove( int, string )
  Winner() ( string, bool )
}

type game struct {
  board *Board
}

func NewGame() *game {
  g := new( game )
  g.board = NewBoard()
  return g
}

func ( g *game ) Board() *Board {
  return g.board
}

func ( g *game ) IsOver() bool {
  return winningSetExists( g.board ) || boardIsFull( g.board )
}

func ( g *game ) Winner() ( string, bool ) {
  spaces := g.Board().Spaces()
  for _,set := range solutions() {
    if allSpacesMatch( g.Board(), set ) {
      return spaces[ set[0] ], true
    }
  }
  return "", false
}

func ( g *game ) IsValidMove( space int ) bool {
  board := g.board
  isInRange := space >= 0 && space < len( board.Spaces() )
  return isInRange && board.Spaces()[ space ] == board.Blank()
}

func ( g *game ) ApplyMove( pos int, mark string ) {
  if ( g.IsValidMove( pos ) ) {
    g.board.Mark( pos, mark )
  }
}

func ( g *game ) Reset() {
  g.board.Reset()
}

// PRIVATE

func boardIsFull( board *Board ) bool {
  for _,mark := range board.Spaces() {
    if mark == board.Blank() { return false }
  }
  return true
}

func winningMark( board *Board ) ( string, bool ) {
  for _,set := range solutions() {
    if allSpacesMatch( board, set ) {
      return board.Spaces()[ set[0] ], true
    }
  }
  return "", false
}

func winningSetExists( board *Board ) ( exists bool ) {
  for _,set := range solutions() {
    exists = exists || allSpacesMatch( board, set )
  }
  return
}

func allSpacesMatch( board *Board, pos []int ) bool {
  spaces := board.Spaces()
  mark := spaces[ pos[ 0 ] ]
  result := mark != board.Blank()
  for _,i := range pos {
    result = result && spaces[ i ] == mark
  }
  return result
}

func solutions() [][]int {
  return [][]int{ []int{ 0, 1, 2 }, []int{ 3, 4, 5 }, []int{ 6, 7, 8 },
                  []int{ 0, 3, 6 }, []int{ 1, 4, 7 }, []int{ 2, 5, 8 },
                  []int{ 0, 4, 8 }, []int{ 2, 4, 6 } }
}
