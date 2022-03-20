package main

import (
	"fmt"
	"time"
)

type Fibonacci struct {
	cache map[uint]uint64
}

func NewFibonacci() Fibonacci {
	f := Fibonacci{}
	f.cache = make(map[uint]uint64)
	f.cache[0] = 0
	f.cache[1] = 1
	return f
}

func (f *Fibonacci) Get(idx uint) uint64 {
	v, ok := f.cache[idx]
	if ok {
		return v
	}
	switch idx {
	case 0:
		return 0
	case 1:
		return 1
	default:
		v = f.Get(idx-1) + f.Get(idx-2)
		f.cache[idx] = v
		return v
	}
}

func main() {
	var idx uint
	f := NewFibonacci()
	for {
		fmt.Print("Enter index of Fibonacci sequence: ")
		_, err := fmt.Scan(&idx)
		if err != nil {
			fmt.Println(err)
			continue
		}
		startCalc := time.Now()
		v := f.Get(idx)
		endCalc := time.Now()
		fmt.Printf("Fibonacci number index %3d is %d in %v us\n", idx, v, endCalc.UnixMicro()-startCalc.UnixMicro())
	}
}
