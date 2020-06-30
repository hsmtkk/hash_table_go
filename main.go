package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hsmtkk/hash_table_go/pkg/hashtable"
)

func main() {
	rand.Seed(time.Now().Unix())
	table := hashtable.New()
	for i := 0; i < 10; i++ {
		v := rand.Intn(100)
		table.Insert(v)
		fmt.Print(v, " ")
	}
	fmt.Println()
	for i := 0; i < 10; i++ {
		v := i * 10
		exist := table.Lookup(v)
		fmt.Printf("%d %v\n", v, exist)
	}

}
