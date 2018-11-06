package peer

import (
    "fmt"
    "table"
)

type Peer struct {
    port int
    cache *table.CacheTable
}

func StartPeer(port int, cache *table.CacheTable) *Peer {
    fmt.Println("Start peer on port", port)
    var peer *Peer = new(Peer)
    peer.cache = cache
    return peer
}
