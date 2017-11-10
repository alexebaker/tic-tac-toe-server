# Tic Tac Toe Server

UNM CS 585: Tic Tac Toe Server

This project is written in [go](https://golang.org/)


## Getting started

This project uses two submodules you will need to download first:

```bash
git submodule init
git submodule update
```

## Code Layout

main.go is simply the entry point to the program. The logic for the server and client are in the `ttts/` directory.

cli.go is initializes the cli arguments. This was found in the following repo: https://github.com/urfave/cli

server.go is called when run in server mode and will listen for connections.

client.go is called when run in client mode and will try and connect to a server.

ttt.go is called to play a ttt game. Both the client and server use the same logic for playing tic tac toe. The logic of the game was found here: https://github.com/jneander/tic-tac-toe-go

## Usage

You will need to compile the source code first. You can run the make file to compile the source code:

```bash
make
```

The server takes command line arguments outlined below:

```bash
>>>./ttt-server -h
NAME:
   Tic Tac Toe Sever - Runs as either a client or server for playing tic tac toe.

USAGE:
   tic-tac-toe-server [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --client, -c              Run tic tac toe in client mode.
   --server value, -s value  If using --client, then this is the address of the server to connect to, otherwise it is the address to listen on. (default: "0.0.0.0")
   --port value, -p value    If using --client, then this is the port of the server to connect to, otherwise it is the port to listen on. (default: 707)
   --num value, -n value     If using --client, then this is the number of games to play in serial. (default: 1)
   --help, -h                show help
   --version, -v             print the version
```

Run in server mode and listen on port 707 on all interfaces:

```bash
./ttt-server -s 0.0.0.0 -p 707
```

Add a `-c` option runs the code in client mode. To connect to a server at 127.0.0.1 listening on port 707:

```bash
./ttt-server -c -s 127.0.0.1 -p 707
```

You can also add a `-n` option to play more than one game at a time:

```bash
./ttt-server -c -s 127.0.0.1 -p 707 -n 20
```


## References

This project uses two outside code sources. One is used for parsing cli arguments and the other is used for the tic tac toe logic.

The cli project can be found here: https://github.com/urfave/cli

The Tic Tac Toe code can be found here: https://github.com/jneander/tic-tac-toe-go


## Authors

* [Alexander Baker](mailto:alexebaker@unm.edu)
