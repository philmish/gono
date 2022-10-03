package main

import (
	"flag"
	"fmt"

	"github.com/philmish/gono/api"
)

var (
    env *string
    addr *string
)

func init() {
    env = flag.String("env", "dev", "current running mode (dev,prod)")
    addr = flag.String("uri", ":7195", "address the server is running on")
}

func main() {
    fmt.Println("Env: ", *env)
    fmt.Println("Addr: ", *addr)

    api.Serve(*addr, *env)
}
