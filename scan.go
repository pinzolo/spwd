package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func scanText(prompt string) (string, error) {
	if prompt != "" {
		fmt.Print(prompt)
	}
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		return "", err
	}
	return in.Text(), nil
}

func scanBool(prompt string) (bool, error) {
	if prompt != "" {
		fmt.Print(prompt)
	}
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		return false, err
	}
	inText := strings.ToLower(in.Text())

	return inText == "y" || inText == "yes", nil
}

func scanPassword(prompt string) (string, error) {
	if prompt != "" {
		fmt.Print(prompt)
	}
	p, err := terminal.ReadPassword(int(syscall.Stdin))
	defer fmt.Println()
	if err != nil {
		return "", err
	}
	return string(p), nil
}
