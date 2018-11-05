package main

import (
    "flag"
    "fmt"
)

func start_service(port int) bool {
    if port == 8080 {
        return true
    } else {
        return false
    }
}

func main() {
    var port = flag.Int("port",  8338, "port to listen on")
    var cache = flag.String("cache", "./data", "cache data path")
    var deamon = flag.Bool("deamon", false, "work in deamon mode")

    flag.Parse()

    var succ = start_service(*port)

    if succ {
        fmt.Println("Worm start on port", *port)
    }

    succ = open_table(*cache)

    if succ {
        fmt.Println("Worm open table", *cache)
    }

    if *deamon {
        fmt.Println("deamon")
    }
}
