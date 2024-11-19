package main

import (
	"flag"
	"os"
)

var (
	flags = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	help = flag.Bool("help", false, `Print usage instructions and exit.`)
)

func init() {

}

func main() {
	flags.Parse(os.Args[1:])

	return
}
