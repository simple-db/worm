package table

import (
    "fmt"
)

type CacheTable struct {
    path string
}

func OpenTable(tablePath string) *CacheTable {
    fmt.Println("open table", tablePath)
    return nil
}
