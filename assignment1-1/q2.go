package cos418_hw1_1

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Sum numbers from channel `nums` and output sum to `out`.
// You should only output to `out` once.
// Do NOT modify function signature.
func sumWorker(nums chan int, out chan int) {
	// TODO: implement me
	// HINT: use for loop over `nums`
	var sum int
	for x := range nums {
		sum += x
	}
	// fmt.Printf("goroutine computed partial sum %d\n", sum)

	out <- sum
}

// Read integers from the file `fileName` and return sum of all values.
// This function must launch `num` go routines running
// `sumWorker` to find the sum of the values concurrently.
// You should use `checkError` to handle potential errors.
// Do NOT modify function signature.
func sum(num int, fileName string) int {
	// TODO: implement me
	// HINT: use `readInts` and `sumWorkers`
	// HINT: used buffered channels for splitting numbers between workers
	file, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Sprintf("Encountered error while opening the file %s: %v", fileName, err))
	}
	numbers, err := readInts(bufio.NewReader(file))
	if err != nil {
		panic(fmt.Sprintf("Encountered error while reading the file %s: %v", fileName, err))
	}
	file.Close()

	channels := make([]chan int, num)
	out := make(chan int)
	size := len(numbers)/num + 1
	// fmt.Printf("Creating %d buffered channels of size %d\n", num, size)
	for i := 0; i < num; i++ {
		channels[i] = make(chan int, size)
		go sumWorker(channels[i], out)
	}
	for i, n := range numbers {
		channels[i%num] <- n
	}
	for i := 0; i < num; i++ {
		close(channels[i])
	}

	sum := 0
	for i := 0; i < num; i++ {
		val := <-out
		sum += val
		// fmt.Printf("Received %d from channel. Current Sum = %d\n", val, sum)
	}
	close(out)

	return sum
}

// Read a list of integers separated by whitespace from `r`.
// Return the integers successfully read with no error, or
// an empty slice of integers and the error that occurred.
// Do NOT modify this function.
func readInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var elems []int
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return elems, err
		}
		elems = append(elems, val)
	}
	return elems, nil
}
