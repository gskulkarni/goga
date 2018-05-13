package main

import (
	"bufio"
	"flag"
	"fmt"
	"goga/tspga"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"runtime"
	"runtime/pprof"
	"strconv"
	"strings"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}

	resp, err := http.Get("http://www.math.uwaterloo.ca/tsp/world/ch71009.tsp")
	if err != nil {
		fmt.Printf("Error getting China tsp data: %v", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	scanner := bufio.NewScanner(strings.NewReader(string(body)))
	cities := []tspga.City{}
	tspSize := 0
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")
		if _, err := strconv.ParseInt(words[0], 10, 32); err == nil {
			X, _ := strconv.ParseFloat(words[1], 64)
			Y, _ := strconv.ParseFloat(words[2], 64)
			city := tspga.City{
				Name: words[0],
				X:    X,
				Y:    Y,
			}
			cities = append(cities, city)
			tspSize++
		}
	}
	c := make(chan float64)
	const popSize = int(1000)
	const generations = int(1000)
	const worlds = int(100)
	for i := 0; i < worlds; i++ {
		go func(k int) {
			tsp := tspga.NewTSP(tspSize, k)
			for _, city := range cities {
				tsp.AddCity(city.Name, city.X, city.Y)
			}
			tsp.CreateInitialPopulation(popSize)
			_, fittestVal := tsp.Evolve()
			for g := 1; g < generations; g++ {
				_, fittestVal = tsp.Evolve()
			}
			c <- fittestVal
		}(i)
	}
	min := math.MaxFloat64
	for i := 0; i < worlds; i++ {
		result := <-c
		if result < min {
			min = result
		}
	}
	fmt.Printf("Best China tour distance: %v\n", min)
}
