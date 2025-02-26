package common

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
	"sync"
)

func ExecutePrevScript(script string, needle string) (bool, error) {
	scriptToEx := fmt.Sprintf(`%s %s`, script, needle)
	cmd := exec.Command("powershell", "-ExecutionPolicy", "Bypass", "-Command", scriptToEx)

	// Get pipes for stdout and stderr
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return false, fmt.Errorf("failed to create stdout pipe: %v", err)
	}
	defer stdoutPipe.Close()

	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return false, fmt.Errorf("failed to create stderr pipe: %v", err)
	}
	defer stderrPipe.Close()

	// Create buffered scanners
	const maxBufferSize = 1024 * 1024
	stdoutScanner := bufio.NewScanner(stdoutPipe)
	stderrScanner := bufio.NewScanner(stderrPipe)

	stdoutScanner.Buffer(make([]byte, maxBufferSize), maxBufferSize)
	stderrScanner.Buffer(make([]byte, maxBufferSize), maxBufferSize)

	// Channels for collecting output and errors
	var wg sync.WaitGroup
	var stdoutLines, stderrLines []string

	// Read stdout
	wg.Add(1)
	go func() {
		defer wg.Done()
		for stdoutScanner.Scan() {
			line := stdoutScanner.Text()
			fmt.Println(line)
			stdoutLines = append(stdoutLines, line)
		}
		if err := stdoutScanner.Err(); err != nil {
			fmt.Printf("Error reading stdout: %v\n", err)
		}
	}()

	// Read stderr
	wg.Add(1)
	go func() {
		defer wg.Done()
		for stderrScanner.Scan() {
			line := stderrScanner.Text()
			fmt.Println(line)
			stderrLines = append(stderrLines, line)
		}
		if err := stderrScanner.Err(); err != nil {
			fmt.Printf("Error reading stderr: %v\n", err)
		}
	}()

	// Start and wait for command
	if err := cmd.Start(); err != nil {
		return false, fmt.Errorf("failed to start command: %v", err)
	}

	// Wait for goroutines to complete
	wg.Wait()

	if err := cmd.Wait(); err != nil {
		if len(stderrLines) > 0 {
			return false, fmt.Errorf("command failed with error output: %s", strings.Join(stderrLines, "\n"))
		}
		return false, fmt.Errorf("command failed: %v", err)
	}

	return cmd.ProcessState.Success(), nil
}

func ExecuteScript(script string, needle string) string {
	scriptToEx := fmt.Sprintf(`%s %s`, script, needle)

	startScript := exec.Command("powershell", "-Command", scriptToEx)

	outputOfStd, err := startScript.StdoutPipe()

	if err != nil {
		fmt.Printf("%s", err)
		fmt.Printf("%s", "Unpredictable behavior happens while returning a StdoutPipe!")
	}

	StderrPipe, err := startScript.StderrPipe()

	if err != nil {
		fmt.Printf("%s", err)
		fmt.Printf("%s", "Unpredictable behavior happens while returning a StderrPipe!")
	}

	err = startScript.Start()

	if err != nil {
		fmt.Println("To preform the init command you must be opening your cmd/powershell as an administrator!")
		fmt.Printf("%s", "Unpredictable behavior happens while starting the process!")
	}

	outputBuffer := bufio.NewScanner(outputOfStd)
	errBuffer := bufio.NewScanner(StderrPipe)

	var resultBufer strings.Builder

	go func() {
		for outputBuffer.Scan() {
			resultBufer.WriteString(outputBuffer.Text() + "\n")
			//fmt.Printf(outputBuffer.Text())
		}
		if err := outputBuffer.Err(); err != nil {
			fmt.Println(err)
			fmt.Printf("%s", "Unpredictable behavior happens while outputing default buffer")
		}
	}()

	go func() {
		for errBuffer.Scan() {
			//fmt.Printf(errBuffer.Text())
		}
		if err := errBuffer.Err(); err != nil {
			fmt.Printf("%s", "Unpredictable behavior happens while outputing Error buffer")
		}
	}()

	err = startScript.Wait()

	if err != nil {
		fmt.Printf("%s", "Unpredictable behavior happens while wating process!")
	}

	return resultBufer.String()

}
