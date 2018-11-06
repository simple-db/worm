package peer

import (
    "fmt"
    "net"
    "strconv"
    "net/rpc"
    "net/http"
)

import (
    "table"
)

type CacheTable = table.CacheTable

type Peer struct {
    cache *CacheTable
}

type GetReq struct {
}

type GetRes struct {
}

type PublishReq struct {
}

type PublishRes struct {
}

type DumpReq struct {
}

type DumpRes struct {
}

func NewPeer(cache *CacheTable) *Peer {
    var peer *Peer = new(Peer)
    peer.cache = cache
    return peer
}

func (peer *Peer) Publish(req *PublishReq, res *PublishRes) error {
    return nil
}

func (peer *Peer) Dump(req *DumpReq, res *DumpRes) error {
    return nil
}

func (peer *Peer) Get(req *GetReq, res *GetRes) error {
    return nil
}

func (peer *Peer) Serve(port int) bool {
    fmt.Println("Start peer on port", port)
    rpc.Register(peer)
    rpc.HandleHTTP()

    l, err := net.Listen("tcp", ":" + strconv.Itoa(port))
    if err != nil {
        fmt.Println("fail to listen on port", port)
    }

    http.Serve(l, nil)
    return true
}

