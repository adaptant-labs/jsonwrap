package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)


func jsonFromFile(filename string) (string, error) {
	jsonBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

func jsonFromStdin() (string, error) {
	var output []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	err := scanner.Err()
	if err != nil {
		return "", err
	}

	return strings.Join(output, ""), nil
}
