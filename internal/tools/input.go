package tools

import (
	"bufio"
	"os"
)

func ReadInput(inputName string) chan string {
	lines := make(chan string)
	go func() {
		defer close(lines)

		file, err := os.Open(inputName)
		PanicOnError(err)
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines <- scanner.Text()
		}

		err = scanner.Err()
		PanicOnError(err)

	}()
	return lines
}
