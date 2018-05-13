package main

import (
	"fmt"
	"goga/tspga"
	"math"
	"testing"
)

func TestDjibouti(t *testing.T) {
	c := make(chan float64)
	const popSize = int(100)
	const generations = int(1000)
	const worlds = int(100)
	for i := 0; i < worlds; i++ {
		go func(k int) {
			tsp := tspga.NewTSP(38, k)
			tsp.AddCity("1", 11003.611100, 42102.500000)
			tsp.AddCity("2", 11108.611100, 42373.888900)
			tsp.AddCity("3", 11133.333300, 42885.833300)
			tsp.AddCity("4", 11155.833300, 42712.500000)
			tsp.AddCity("5", 11183.333300, 42933.333300)
			tsp.AddCity("6", 11297.500000, 42853.333300)
			tsp.AddCity("7", 11310.277800, 42929.444400)
			tsp.AddCity("8", 11416.666700, 42983.333300)
			tsp.AddCity("9", 11423.888900, 43000.277800)
			tsp.AddCity("10", 11438.333300, 42057.222200)
			tsp.AddCity("11", 11461.111100, 43252.777800)
			tsp.AddCity("12", 11485.555600, 43187.222200)
			tsp.AddCity("13", 11503.055600, 42855.277800)
			tsp.AddCity("14", 11511.388900, 42106.388900)
			tsp.AddCity("15", 11522.222200, 42841.944400)
			tsp.AddCity("16", 11569.444400, 43136.666700)
			tsp.AddCity("17", 11583.333300, 43150.000000)
			tsp.AddCity("18", 11595.000000, 43148.055600)
			tsp.AddCity("19", 11600.000000, 43150.000000)
			tsp.AddCity("20", 11690.555600, 42686.666700)
			tsp.AddCity("21", 11715.833300, 41836.111100)
			tsp.AddCity("22", 11751.111100, 42814.444400)
			tsp.AddCity("23", 11770.277800, 42651.944400)
			tsp.AddCity("24", 11785.277800, 42884.444400)
			tsp.AddCity("25", 11822.777800, 42673.611100)
			tsp.AddCity("26", 11846.944400, 42660.555600)
			tsp.AddCity("27", 11963.055600, 43290.555600)
			tsp.AddCity("28", 11973.055600, 43026.111100)
			tsp.AddCity("29", 12058.333300, 42195.555600)
			tsp.AddCity("30", 12149.444400, 42477.500000)
			tsp.AddCity("31", 12286.944400, 43355.555600)
			tsp.AddCity("32", 12300.000000, 42433.333300)
			tsp.AddCity("33", 12355.833300, 43156.388900)
			tsp.AddCity("34", 12363.333300, 43189.166700)
			tsp.AddCity("35", 12372.777800, 42711.388900)
			tsp.AddCity("36", 12386.666700, 43334.722200)
			tsp.AddCity("37", 12421.666700, 42895.555600)
			tsp.AddCity("38", 12645.000000, 42973.333300)
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
	fmt.Printf("Best Djibouti tour distance: %v\n", min)
}

func TestDjiboutiRandom(t *testing.T) {
	c := make(chan float64)
	const popSize = int(100)
	const worlds = int(1000000)
	for i := 0; i < worlds; i++ {
		go func(k int) {
			tsp := tspga.NewTSP(38, k)
			tsp.AddCity("1", 11003.611100, 42102.500000)
			tsp.AddCity("2", 11108.611100, 42373.888900)
			tsp.AddCity("3", 11133.333300, 42885.833300)
			tsp.AddCity("4", 11155.833300, 42712.500000)
			tsp.AddCity("5", 11183.333300, 42933.333300)
			tsp.AddCity("6", 11297.500000, 42853.333300)
			tsp.AddCity("7", 11310.277800, 42929.444400)
			tsp.AddCity("8", 11416.666700, 42983.333300)
			tsp.AddCity("9", 11423.888900, 43000.277800)
			tsp.AddCity("10", 11438.333300, 42057.222200)
			tsp.AddCity("11", 11461.111100, 43252.777800)
			tsp.AddCity("12", 11485.555600, 43187.222200)
			tsp.AddCity("13", 11503.055600, 42855.277800)
			tsp.AddCity("14", 11511.388900, 42106.388900)
			tsp.AddCity("15", 11522.222200, 42841.944400)
			tsp.AddCity("16", 11569.444400, 43136.666700)
			tsp.AddCity("17", 11583.333300, 43150.000000)
			tsp.AddCity("18", 11595.000000, 43148.055600)
			tsp.AddCity("19", 11600.000000, 43150.000000)
			tsp.AddCity("20", 11690.555600, 42686.666700)
			tsp.AddCity("21", 11715.833300, 41836.111100)
			tsp.AddCity("22", 11751.111100, 42814.444400)
			tsp.AddCity("23", 11770.277800, 42651.944400)
			tsp.AddCity("24", 11785.277800, 42884.444400)
			tsp.AddCity("25", 11822.777800, 42673.611100)
			tsp.AddCity("26", 11846.944400, 42660.555600)
			tsp.AddCity("27", 11963.055600, 43290.555600)
			tsp.AddCity("28", 11973.055600, 43026.111100)
			tsp.AddCity("29", 12058.333300, 42195.555600)
			tsp.AddCity("30", 12149.444400, 42477.500000)
			tsp.AddCity("31", 12286.944400, 43355.555600)
			tsp.AddCity("32", 12300.000000, 42433.333300)
			tsp.AddCity("33", 12355.833300, 43156.388900)
			tsp.AddCity("34", 12363.333300, 43189.166700)
			tsp.AddCity("35", 12372.777800, 42711.388900)
			tsp.AddCity("36", 12386.666700, 43334.722200)
			tsp.AddCity("37", 12421.666700, 42895.555600)
			tsp.AddCity("38", 12645.000000, 42973.333300)
			tsp.CreateInitialPopulation(popSize)
			_, fittestVal, _ := tsp.SelectFittest()
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
	fmt.Printf("Best Djibouti random tour distance: %v\n", min)
}

func TestCustom(t *testing.T) {
	c := make(chan float64)
	const popSize = int(100)
	const generations = int(1000)
	const worlds = int(100)
	for i := 0; i < worlds; i++ {
		go func(k int) {
			tsp := tspga.NewTSP(26, k)
			tsp.AddCity("A", 20.0, 30.0)
			tsp.AddCity("B", 25.0, 65.0)
			tsp.AddCity("C", 60.0, 130.0)
			tsp.AddCity("D", 190.0, 10.0)
			tsp.AddCity("E", 150.0, 150.0)
			tsp.AddCity("F", 5.0, 105.0)
			tsp.AddCity("G", 75.0, 15.0)
			tsp.AddCity("H", 120.0, 130.0)
			tsp.AddCity("I", 40.0, 85.0)
			tsp.AddCity("J", 30.0, 60.0)
			tsp.AddCity("K", 50.0, 30.0)
			tsp.AddCity("L", 105.0, 15.0)
			tsp.AddCity("M", 80.0, 135.0)
			tsp.AddCity("N", 90.0, 180.0)
			tsp.AddCity("O", 140.0, 155.0)
			tsp.AddCity("P", 115.0, 115.0)
			tsp.AddCity("Q", 70.0, 160.0)
			tsp.AddCity("R", 10.0, 70.0)
			tsp.AddCity("S", 25.0, 75.0)
			tsp.AddCity("T", 185.0, 20.0)
			tsp.AddCity("U", 20.0, 30.0)
			tsp.AddCity("V", 175.0, 195.0)
			tsp.AddCity("W", 120.0, 110.0)
			tsp.AddCity("X", 90.0, 140.0)
			tsp.AddCity("Y", 15.0, 100.0)
			tsp.AddCity("Z", 195.0, 45.0)
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
	fmt.Printf("Best custom tour distance: %v\n", min)
}
