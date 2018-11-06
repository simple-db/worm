package table

import (
    "fmt"
)

type CacheTable struct {
    path string
}

func OpenTable(tablePath string) *CacheTable {
    fmt.Println("open table", tablePath)
    var table = new(CacheTable)

    if table.open(tablePath) {
        return table
    }

    return nil
}

func (table CacheTable) open(path string) bool {
    return false
}

func (table CacheTable) put(slice []byte, n int, file string, idx int) bool {
    return false
}

func (table CacheTable) get() bool {
    return false
}
