// There are all tasks here.
// All tasks were wraped in functions.
package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	FirstTask()
	SecondTask()
	ThirdTask()
	FourthTask()
	FifthTask()
	SixthTask()
	SeventhTask()
	EigthTask()
}

// First task

func getMapCounts(words []string) map[string]int {
	counts := make(map[string]int)

	for _, word := range words {
		if _, ok := counts[word]; !ok {
			counts[word] = 1
		} else {
			counts[word]++
		}
	}

	return counts
}

func getSortedList(counts map[string]int) [][]string {
	var words [][]string
	var keys []string

	// Sort map
	for key := range counts {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return counts[keys[i]] > counts[keys[j]]
	})

	// Create result slice
	for _, key := range keys {
		words = append(words, []string{key, strconv.Itoa(counts[key])})
	}

	return words
}

func countWords(str string) {
	words := strings.Split(str, " ")

	counts := getMapCounts(words)
	wordsPairs := getSortedList(counts)

	// Output result
	for _, pair := range wordsPairs {
		fmt.Printf("%s: %s\n", pair[0], pair[1])
	}
}

func FirstTask() {
	countWords("hello world hello hello everyone")
}

// Second task

func getRectangleData(width, height int) (int, int, float64) {
	square := width * height
	perimeter := width*2 + height*2
	diagonal := math.Sqrt((math.Pow(float64(width), 2) + math.Pow(float64(height), 2)))
	return square, perimeter, diagonal
}

func SecondTask() {
	s, p, d := getRectangleData(5, 4)
	fmt.Println("Площадь:", s)
	fmt.Println("Периметр:", p)
	fmt.Printf("Диагональ: %.4f\n", d)
}

// Third task

func countSum(N int) func() int {
	s := 0
	return func() int {
		if s < N {
			s++
		}
		return s
	}
}

func ThirdTask() {
	first := countSum(10)
	fmt.Println(first())
	fmt.Println(first())
	fmt.Println(first())
	fmt.Println(first())
	fmt.Println(first())
	fmt.Println(first())
	fmt.Println(first())
	fmt.Println(first())
	fmt.Println(first())
	fmt.Println(first())
}

// Fourth task

func multiply(n int) func(int) int {
	a := n

	return func(b int) int {
		return a * b
	}
}

func FourthTask() {
	multiplyTen := multiply(10)
	fmt.Println(multiplyTen(1))
	fmt.Println(multiplyTen(2))
	fmt.Println(multiplyTen(3))
}

// Fifth task

func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func FifthTask() {
	fmt.Println("Сумма:", applyOperation(5, 3, func(i1, i2 int) int { return i1 + i2 }))
	fmt.Println("Разность:", applyOperation(5, 3, func(i1, i2 int) int { return i1 - i2 }))
	fmt.Println("Произведение:", applyOperation(5, 3, func(i1, i2 int) int { return i1 * i2 }))
	fmt.Println("Частное:", applyOperation(6, 3, func(i1, i2 int) int { return i1 / i2 }))
}

// Sixth task

func recursiveFactorial(n int) int {
	if n == 1 {
		return n
	}
	return n * recursiveFactorial(n-1)
}

func SixthTask() {
	fmt.Println("Факториал числа 5:", recursiveFactorial(5))
}

// Seventh task

func SeventhTask() {
	sum := func(numbers ...int) int {
		s := 0

		for _, num := range numbers {
			s += num
		}

		return s
	}

	fmt.Println(sum(1, 5, 8, 2, 9))
}

// Eigth task

func calculateFare(taxiType string, kilometers, minutes int) (finalCost int, taxiDefinition string) {
	// Cost for taking a seat
	baseCost := 50

	// Cost for a minute
	timeCost := 5

	// Cost for a kilometer
	taxiCosts := map[string]int{
		"econom":   10,
		"standart": 15,
		"business": 25,
	}

	taxiDefinitions := map[string]string{
		"econom":   "An econom taxi will arrive.",
		"standart": "A standart taxi will arrive.",
		"business": "A business taxi will arrive.",
	}

	if taxiCost, ok := taxiCosts[taxiType]; !ok {
		return 0, ""
	} else {
		return baseCost + taxiCost*kilometers + timeCost*minutes, taxiDefinitions[taxiType]
	}
}

func EigthTask() {
	cost, def := calculateFare("econom", 10, 60)
	fmt.Println("Стоимость для эконома на 10 километров за 60 минут:", cost)
	fmt.Println("Описание такси:", def)
}
