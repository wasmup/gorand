package main

import (
	_ "embed"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
	"strings"
	"time"
)

//go:embed input.txt
var txt string

func main() {
	{
		// runtime.SetBlockProfileRate(1)

		f, err := os.Create("cpu.out")
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer func() {
			f.Close()
			fmt.Println("...cpu.out done...")
		}()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}
	rand.Seed(time.Now().UnixNano())
	a := strings.Fields(txt)
	for i := 0; i < 1_000_000; i++ {
		rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
		os.Stdout.Write([]byte(strings.Join(a, " ")))
	}
}
