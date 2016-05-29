package main

import (
    "fmt"
    "time"
    "sync"
)

var wg sync.WaitGroup
var ret_arr []int

func YSort(arr []int) {
    wg.Add(len(arr))
    for i := range arr {
        go func (i int) {
            defer wg.Done()
            time.Sleep(time.Duration(arr[i]) * time.Millisecond)
            fmt.Println(arr[i])
            ret_arr = append(ret_arr, arr[i])
        }(i)
    }
    wg.Wait()
    fmt.Println(ret_arr)
}


func main() {
    arr := []int{5, 2, 6, 3, 1, 4}
    YSort(arr)
}
