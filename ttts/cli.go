/***********************************************
 * The file initializes the cli arguments.     *
 * The logic for parsing the argument is from: *
 * https://github.com/urfave/cli               *
 ***********************************************/

package ttts


import(
    "os"
    "./cli" // https://github.com/urfave/cli
)


func Run() {
    var address string = ""
    var port int = 0
    var client bool = false
    var numGames int = 1

    app := cli.NewApp()
    app.Name = "Tic Tac Toe Sever"
    app.Usage = "Runs as either a client or server for playing tic tac toe."

    app.Flags = []cli.Flag {
        cli.BoolFlag {
            Name: "client, c",
            Usage: "Run tic tac toe in client mode.",
            Destination: &client,
        },
        cli.StringFlag {
            Name: "server, s",
            Value: "0.0.0.0",
            Usage: "If using --client, then this is the address of the server to connect to, otherwise it is the address to listen on.",
            Destination: &address,
        },
        cli.IntFlag {
            Name: "port, p",
            Value: 707,
            Usage: "If using --client, then this is the port of the server to connect to, otherwise it is the port to listen on.",
            Destination: &port,
        },
        cli.IntFlag {
            Name: "num, n",
            Value: 1,
            Usage: "If using --client, then this is the number of games to play in serial.",
            Destination: &numGames,
        },
    }

    app.Action = func(args *cli.Context) error {
        if client {
            runClient(address, port, numGames)
        } else {
            runServer(address, port)
        }

        return nil
    }

    app.Run(os.Args)
    return
}
