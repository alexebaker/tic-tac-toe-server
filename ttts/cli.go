package ttts


import(
    "os"
    "github.com/urfave/cli"
)


func Run() {
    var address string = ""
    var port int = 0
    var client bool = false

    app := cli.NewApp()
    app.Name = "Tic Tac Toe Sever"
    app.Usage = "Runs as either a client or server for playing tic tac toe."

    app.Flags = []cli.Flag {
        cli.BoolTFlag {
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
    }

    app.Action = func(args *cli.Context) error {
        if (client) {
            run_client(address, port)
        } else {
            run_server(address, port)
        }

        return nil
    }

    app.Run(os.Args)
    return
}
