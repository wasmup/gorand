package main

import (
	_ "embed"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
	"strings"
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

	a := strings.Fields(txt)
	s := shuffleText1(a)
	fmt.Println(s)

}

func shuffleText2(a []string) string {
	sb := &strings.Builder{}
	r := rand.New(rand.NewSource(0))
	for i := 0; i < 1_000_000; i++ {
		r.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
		sb.WriteString(strings.Join(a, " "))
		sb.WriteRune('\n')
	}
	return sb.String()
}

func shuffleText1(a []string) string {
	sb := &strings.Builder{}
	for i := 0; i < 1_000_000; i++ {
		rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
		sb.WriteString(strings.Join(a, " "))
		sb.WriteRune('\n')
	}
	return sb.String()
}
