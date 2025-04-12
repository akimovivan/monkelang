package main

import (
	"fmt"
	"os"

	"github.com/akimovivan/monkelang/repl"
)

func main() {
	fmt.Print(`
 _ __ ___   ___  _ __ | | _____
| '_ ' _ \ / _ \| '_ \| |/ / _ \
| | | | | | (_) | | | |   <  __/
|_| |_| |_|\___/|_| |_|_|\_\___|
`)
	repl.Start(os.Stdin, os.Stdout)
}
