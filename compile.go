package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/google/uuid"
)

func Compile(code string) (string, error) {
	uuid := uuid.New().String()
	path := fmt.Sprintf("compile/%s.go", uuid)
	file, err := os.Create(path)
	if err != nil {
		return "", err
	}
	file.Write([]byte(code))
	out, err := exec.Command("go", "run", path).Output()
	if err != nil {
		if err := os.Remove(path); err != nil {
			return "", err
		}
		return "", err
	}
	if err := os.Remove(path); err != nil {
		return "", err
	}
	return string(out), nil
}
