package internal

import (
	"bufio"
	"log"
	"os/exec"
	"strings"
)

func ArrayContains(arr []string, target string) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}

func GoBash(command string, args ...string) string {
	var sb strings.Builder
	cmd := exec.Command(command, args...)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		m := scanner.Text()
		sb.WriteString(m + "\n")
	}

	scanner = bufio.NewScanner(stderr)
	for scanner.Scan() {
		m := scanner.Text()
		log.Println(m)
		sb.WriteString(m + "\n")
	}

	cmd.Wait()

	return sb.String()
}
