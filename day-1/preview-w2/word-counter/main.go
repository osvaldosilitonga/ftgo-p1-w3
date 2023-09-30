package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"
)

func wordCheck(line string, text map[string]int, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)

	words := strings.Split(strings.TrimSpace(line), " ") // Split words

	for _, word := range words {
		text[word] += 1
	}
}

func main() {
	// Recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("[ERROR] :", r)
		}
	}()

	text := make(map[string]int)
	var mutex sync.Mutex

	// Check argumen
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Isi argumen (nama_file.txt) sebanyak 1")
		return
	}

	// Regexp - check ekstensi file dalam argumen
	regx := regexp.MustCompile(`^\w+.txt$`)
	if !regx.MatchString(args[1]) {
		fmt.Println("Extensi file harus .txt")
		return
	}

	// Open File
	file, err := os.Open(args[1])
	if err != nil {
		panic("File not exist!")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	wg := sync.WaitGroup{}

	// Looping file line
	for scanner.Scan() {
		line := scanner.Text()

		mutex.Lock()
		go wordCheck(line, text, &wg)
		mutex.Unlock()
	}

	wg.Wait()

	fmt.Println("Process Done")
	fmt.Println("## Word Count ##")

	// Print Output
	for key, value := range text {
		fmt.Printf("%s : %d\n", key, value)
	}

}
