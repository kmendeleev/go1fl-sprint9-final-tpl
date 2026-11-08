package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	if size <= 0 {
		return []int{}
	}
	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(1_000_000_000) + 1
	}
	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}
	largest := data[0]
	for _, value := range data[1:] {
		if value > largest {
			largest = value
		}
	}
	return largest
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	if len(data) < CHUNKS {
		return maximum(data)
	}
	partSize := len(data) / CHUNKS
	maxima := make([]int, CHUNKS)
	var wg sync.WaitGroup
	for i := 0; i < CHUNKS; i++ {
		start := i * partSize
		end := start + partSize
		if i == CHUNKS-1 {
			end = len(data)
		}
		wg.Add(1)
		go func(index int, part []int) {
			defer wg.Done()
			maxima[index] = maximum(part)
		}(i, data[start:end])
	}
	wg.Wait()
	return maximum(maxima)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d частях\n", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)
}
