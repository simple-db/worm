/**
 * File: table.go
 * Author: wtzhuque@163.com
 * Description: implement data cache table
 */
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

func (table *CacheTable) open(path string) bool {
    return false
}

func (table *CacheTable) Put(slice []byte, n int, file string, idx int) bool {
    return false
}

func (table *CacheTable) Get() bool {
    return false
}
