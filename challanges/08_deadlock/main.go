// find the deadlock in this code
//
// the challange may require knowledge about:
//   * for loop
//   * anyonimus self executed func
//   * for loop variable scope
//   * channels
//   * fan in/out pattern
//   * sync lib
//   * fmt Println method
//   * array creation
//   * race condition analisis
//     $ go run -race
//
package main

import (
	crypto_rand "crypto/rand"
	"fmt"
	math_rand "math/rand"
	"runtime"
	"strings"
	"sync"
	"time"
)

func main() {
	for n := range mergeChannels(spawnWorkToChannels()) {
		fmt.Println(n)
	}
}

func spawnWorkToChannels() []chan string {
	var channels []chan string

	cpuCount := runtime.NumCPU()
	for i := 0; i < cpuCount; i++ {

		in := make(chan string)
		out := make(chan string)
		channels = append(channels, out)

		go publisher(in)
		go consumer(in, out)

	}

	return channels
}

func mergeChannels(channels []chan string) chan string {
	out := make(chan string)

	var wg sync.WaitGroup
	wg.Add(len(channels))
	for _, ch := range channels {
		go func(ch <-chan string) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(ch)
	}

	wg.Wait()
	close(out)
	return out
}

func publisher(out chan<- string) {
	for i := 0; i < 10; i++ {
		out <- randString(randomInt(1, 100))
	}

	close(out)
}

func consumer(in <-chan string, out chan<- string) {
	for word := range in {
		out <- strings.ToUpper(word)
	}

	close(out)
}

func randomInt(min, max int) int {
	math_rand.Seed(time.Now().Unix())
	return math_rand.Intn(max-min) + min
}

func randString(n int) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	crypto_rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}
