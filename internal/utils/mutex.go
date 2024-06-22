package utils

import ( 
    "log"
    "sync"
)

// 使用信号量实现互斥机制
var (
    mu        sync.Mutex
    isReading = false
    isWriting = false
)

// 当创建一个读进程时，需满足没有进程正在写入才可创建，否则弹窗；
func StartRead() bool {
    mu.Lock()
    defer mu.Unlock()
    if isWriting {
        log.Println("Failed creating Reading Process: Other Process is already writing!")
        return false
    }
    isReading = true
    return true
}

func EndRead() {
    mu.Lock()
    defer mu.Unlock()
    isReading = false
}

// 当创建一个写进程时，需满足没有进程正在读写才可创建，否则弹窗：
func StartWrite() bool {
    mu.Lock()
    defer mu.Unlock()
    if isReading || isWriting {
        log.Println("Failed creating Writing Process: Other Process is already writing or reading!")
        return false
    }
    isWriting = true
    return true
}

func EndWrite() {
    mu.Lock()
    defer mu.Unlock()
    isWriting = false
}
