package main

import (
	"bufio"
	"fmt"
	"image/png"
	"log"
	"os"
)

func main() {
	asciiCharset := ".:-=+*#%%"

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter the textfile containing ASCII >>> ")
	scanner.Scan()
	filename := scanner.Text()

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		log.Fatalln(err)
	}

	fmt.Print("Enter the characterset you want to use (or none for default) >>> ")

	scanner.Scan()
	if inp := scanner.Text(); inp != "" {
		asciiCharset = inp
	}

	outputFile := ""

	for {
		fmt.Print("Enter the file you want to save to >>> ")
		scanner.Scan()

		inp := scanner.Text()
		if inp == "" {
			fmt.Println("Please enter a valid filename.")
			continue
		} else {
			outputFile = inp
			break
		}
	}

	fmt.Printf("Using characterset %q , opening file %q , output file to %q\n", asciiCharset, filename, outputFile)

	input, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Create(outputFile)

	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	png.Encode(file, processASCII(asciiCharset, 10, string(input)))
}
