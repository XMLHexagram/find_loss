package main

import "fmt"

func DealError(err error)  {
	if err != nil {
		fmt.Println(err)
	}
}
