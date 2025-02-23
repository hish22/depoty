package common

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

func ExecutePrevScript(script string, needle string) bool {
	scriptToEx := fmt.Sprintf(`%s %s`, script, needle)

	// startScript := exec.Command("powershell", "-Command",
	// 	"Start-Process", "powershell",
	// 	"-ArgumentList", `'-ExecutionPolicy', 'Bypass', '-Command', "`+scriptToEx+`"`,
	// 	"-Verb", "runAs")

	startScript := exec.Command("powershell", "-ExecutionPolicy", "Bypass", "-Command", scriptToEx)

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
		fmt.Printf("To preform the init command you must be opening your cmd/powershell as an administrator!")
		fmt.Printf("%s", "Unpredictable behavior happens while starting the process!")
	}

	const MaxSize = 1024 * 1024

	// Increase the size of the buffer.
	buffer := make([]byte, MaxSize)

	outputBuffer := bufio.NewScanner(outputOfStd)

	outputBuffer.Buffer(buffer, MaxSize)

	errBuffer := bufio.NewScanner(StderrPipe)

	errBuffer.Buffer(buffer, MaxSize)

	go func() {
		for outputBuffer.Scan() {
			fmt.Println(outputBuffer.Text())
		}
		if err := outputBuffer.Err(); err != nil {
			fmt.Println(err)
			fmt.Printf("%s", "Unpredictable behavior happens while outputing default buffer")
		}
	}()

	go func() {
		for errBuffer.Scan() {
			fmt.Println(errBuffer.Text())
		}
		if err := errBuffer.Err(); err != nil {
			fmt.Printf("%s", "Unpredictable behavior happens while outputing Error buffer")
		}
	}()

	err = startScript.Wait()

	if err != nil {
		fmt.Printf("%s", "Unpredictable behavior happens while wating process!")
	}

	return startScript.ProcessState.Success()
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
