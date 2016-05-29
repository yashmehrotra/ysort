package ysort

import (
	"sync"
	"time"
)

func Sort(arr []int) []int {
	var wg sync.WaitGroup
	var ret_arr []int
	wg.Add(len(arr))
	for i := range arr {
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Duration(arr[i]) * time.Millisecond)
			ret_arr = append(ret_arr, arr[i])
		}(i)
	}
	wg.Wait()
	return ret_arr
}
