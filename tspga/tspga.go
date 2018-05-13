package tspga

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const mutationProbability = float64(0.05)

// City is an individual city
type City struct {
	Name string
	X    float64
	Y    float64
}

// TSP is an instance of the TRaveling Salesman Problem
type TSP struct {
	Size                int
	Cities              []*City
	Population          Population
	Source              *rand.Rand
	MutationProbability float64
}

// Tour is a specific sequence of cities
type Tour struct {
	Cities []*City
	Score  float64
}

// Population is a collection of tours
type Population struct {
	Tours []*Tour
	Size  int
	Score float64
}

// NewTSP returns a new instance of TSP
func NewTSP(size, seed int) *TSP {
	tsp := &TSP{
		Size:                size,
		Cities:              nil,
		Source:              rand.New(rand.NewSource(time.Now().UnixNano() + int64(seed))),
		MutationProbability: mutationProbability,
	}
	return tsp
}

// AddCity appends a new city to the problem instance
func (t *TSP) AddCity(name string, x, y float64) {
	city := &City{
		Name: name,
		X:    x,
		Y:    y,
	}
	t.Cities = append(t.Cities, city)
}

// CreateInitialPopulation creates an initial population of the specified size
// Must call AddCity first
func (t *TSP) CreateInitialPopulation(count int) {
	t.Population = Population{}
	for i := 0; i < count; i++ {
		perm := t.Source.Perm(t.Size)
		tour := &Tour{}
		for j := 0; j < t.Size; j++ {
			tour.Cities = append(tour.Cities, t.Cities[perm[j]])
		}
		t.Population.Tours = append(t.Population.Tours, tour)
		t.Population.Tours[i].Score = t.IndividualFitness(i)
		t.Population.Size++
		t.Population.Score += t.Population.Tours[i].Score
	}
}

// IndividualFitness computes the fitness of an individual tour specified by its index
func (t *TSP) IndividualFitness(idx int) float64 {
	f := 0.0
	for i := 0; i < t.Size-1; i++ {
		f += t.distance(t.Population.Tours[idx].Cities[i], t.Population.Tours[idx].Cities[i+1])
	}
	f += t.distance(t.Population.Tours[idx].Cities[t.Size-1], t.Population.Tours[idx].Cities[0])
	return f
}

func (t *TSP) distance(city1, city2 *City) float64 {
	d := math.Sqrt(math.Pow((city1.X-city2.X), 2) + math.Pow((city1.Y-city2.Y), 2))
	return d
}

// SelectFittest returns the fittest individual tour
func (t *TSP) SelectFittest() (*Tour, float64, int) {
	fittestVal := math.MaxFloat64
	fittest := &Tour{}
	fittestIdx := -1
	for i := 0; i < t.Population.Size; i++ {
		fitVal := t.Population.Tours[i].Score
		if fitVal < fittestVal {
			fittestVal = fitVal
			fittest = t.Population.Tours[i]
			fittestIdx = i
		}
	}
	return fittest, fittestVal, fittestIdx
}

// SelectParents returns a pair of parents weighted by fitness
func (t *TSP) SelectParents() (*Tour, *Tour) {
	t1 := -1
	p := t.Source.Float64()
	if p <= 0.5 {
		_, _, t1 = t.SelectFittest()
	} else {
		t1 = t.Source.Intn(t.Population.Size)
	}
	t2 := t.Source.Intn(t.Population.Size)
	for t1 == t2 {
		t2 = t.Source.Intn(t.Population.Size)
	}
	return t.Population.Tours[t1], t.Population.Tours[t2]
}

// PerformCrossover generates a child tour from a given pair of parents
func (t *TSP) PerformCrossover(tour1, tour2 *Tour) *Tour {
	// Select [start, end] cities from tour1 and the remaining from tour2
	start := t.Source.Intn(t.Size)
	end := t.Source.Intn(t.Size)
	if start > end {
		tmp := start
		start = end
		end = tmp
	}
	cities := make([]*City, t.Size)
	fromTour1 := make(map[string]bool)
	for i := start; i <= end; i++ {
		cities[i] = tour1.Cities[i]
		fromTour1[cities[i].Name] = true
	}
	idx := 0
	for i := 0; i < t.Size; i++ {
		if _, ok := fromTour1[tour2.Cities[i].Name]; !ok {
			if idx == start {
				idx += end - start + 1
			}
			cities[idx] = tour2.Cities[i]
			idx++
		}
	}
	tour := &Tour{
		Cities: cities,
	}
	return tour
}

// PerformMutation mutates a specific individual given the index
func (t *TSP) PerformMutation(idx int) {
	// Swap two positions with a probability
	p1 := t.Source.Intn(t.Size)
	p2 := t.Source.Intn(t.Size)
	p := t.Source.Float64()
	if p <= mutationProbability {
		temp := t.Population.Tours[idx].Cities[p1]
		t.Population.Tours[idx].Cities[p1] = t.Population.Tours[idx].Cities[p2]
		t.Population.Tours[idx].Cities[p2] = temp
	}
}

// Evolve does the evolution of the population and returns the best tour
func (t *TSP) Evolve() (*Tour, float64) {
	// Always select the fittest for the next generation
	fittest, fittestVal, _ := t.SelectFittest()
	newPop := Population{
		Tours: nil,
		Size:  t.Population.Size,
		Score: fittestVal,
	}
	newPop.Tours = append(newPop.Tours, fittest)
	for i := 1; i < t.Population.Size; i++ {
		// Each of the remaining can die with a probability
		if t.eliminated(i, fittestVal) {
			tour1, tour2 := t.SelectParents()
			child := t.PerformCrossover(tour1, tour2)
			newPop.Tours = append(newPop.Tours, child)
		} else {
			newPop.Tours = append(newPop.Tours, t.Population.Tours[i])
		}
	}
	t.Population = newPop
	for i := 1; i < t.Population.Size; i++ {
		t.PerformMutation(i)
		t.Population.Tours[i].Score = t.IndividualFitness(i)
		t.Population.Score += t.Population.Tours[i].Score
	}
	//t.dumpPopulation()
	fittest, fittestVal, _ = t.SelectFittest()
	return fittest, fittestVal
}

func (t *TSP) eliminated(idx int, fittest float64) bool {
	if t.Population.Tours[idx].Score == fittest {
		return true
	}
	normalizedScore := t.Population.Tours[idx].Score / (t.Population.Score - fittest)
	p := t.Source.Float64()
	if p >= normalizedScore {
		return true
	}
	return false
}

func (t *TSP) dumpPopulation() {
	for p := 0; p < t.Population.Size; p++ {
		fmt.Printf("\nTour %d\n", p)
		for i := 0; i < t.Size; i++ {
			fmt.Printf("%s ", t.Population.Tours[p].Cities[i].Name)
		}
		fmt.Printf(" fitness = %v ", t.IndividualFitness(p))
	}
}
