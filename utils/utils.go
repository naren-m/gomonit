package utils

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// GetScreenSize gets the screen size and returns rows and columns
func GetScreenSize() (int, int, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
		return -1, -1, err
	}

	res := strings.Split(strings.TrimSpace(string(out)), " ")

	rows, err := strconv.ParseInt(res[0], 10, 64)
	if err != nil {
		log.Fatal(err)
		return -1, -1, err
	}
	cols, err := strconv.ParseInt(res[1], 10, 64)
	if err != nil {
		log.Fatal(err)
		return -1, -1, err
	}

	return int(rows), int(cols), nil
}
