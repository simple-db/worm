package main

import (
    "flag"
    "fmt"
    "table"
    "peer"
)

func main() {
    var port = flag.Int("port",  8338, "port to listen on")
    var cachePath = flag.String("cache", "./data", "cache data path")
    var deamon = flag.Bool("deamon", false, "work in deamon mode")

    flag.Parse()

    var cache = table.OpenTable(*cachePath)

    if cache != nil {
        fmt.Println("Worm open table", *cachePath)
    }

    var peer = peer.NewPeer(cache)

    if peer != nil {
        fmt.Println("Worm start on port", *port)
    }

    peer.Serve(*port)

    if *deamon {
        fmt.Println("deamon")
    }
}
