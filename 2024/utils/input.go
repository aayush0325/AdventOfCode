package utils

import (
	"bufio"
	"os"
	"strconv"
)

/**
* function to parse input from a file returns a slice of strings seperated by newlines or space
**/
func ReadFileAsString(filename string) ([]string, error) {
	content, err := os.Open(filename)

	if err != nil {
		return nil ,err
	}	
	defer content.Close()

	scanner := bufio.NewScanner(content)
	scanner.Split(bufio.ScanWords)

	result := make([]string, 0)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

/**
* function to parse input from a file returns a slice of integers seperated by newlines or space
**/
func ReadFileAsInt(filename string) ([]int, error) {
	content, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(content)
	scanner.Split(bufio.ScanWords)

	result := make([]int, 0)

	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text()) 
		
		if err != nil {
			return nil, err
		}
		result = append(result, int(value))
		
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}