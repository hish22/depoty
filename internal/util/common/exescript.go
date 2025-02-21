package common

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
)

func ExecutePrevScript(script string, needle string) {
	scriptToEx := fmt.Sprintf(`%s %s`, script, needle)

	startScript := exec.Command("powershell", "-Command",
		"Start-Process", "powershell",
		"-ArgumentList", `'-ExecutionPolicy', 'Bypass', '-Command', "`+scriptToEx+`"`,
		"-Verb", "runAs")

	outputOfStd, err := startScript.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		log.Fatal("Unpredictable behavior happens while returning a StdoutPipe!")
	}

	StderrPipe, err := startScript.StderrPipe()

	if err != nil {
		fmt.Println(err)
		log.Fatal("Unpredictable behavior happens while returning a StderrPipe!")
	}

	err = startScript.Start()

	if err != nil {
		fmt.Println("To preform the init command you must be opening your cmd/powershell as an administrator!")
		log.Fatal("Unpredictable behavior happens while starting the process!")
	}

	outputBuffer := bufio.NewScanner(outputOfStd)
	errBuffer := bufio.NewScanner(StderrPipe)

	go func() {
		for outputBuffer.Scan() {
			fmt.Println(outputBuffer.Text())
		}
		if err := outputBuffer.Err(); err != nil {
			log.Fatal("Unpredictable behavior happens while outputing default buffer")
		}
	}()

	go func() {
		for errBuffer.Scan() {
			fmt.Println(errBuffer.Text())
		}
		if err := errBuffer.Err(); err != nil {
			log.Fatal("Unpredictable behavior happens while outputing Error buffer")
		}
	}()

	err = startScript.Wait()

	if err != nil {
		log.Fatal("Unpredictable behavior happens while wating process!")
	}

	fmt.Println("Choco Installed successfully")

}

func ExecuteScript(script string, needle string) string {
	scriptToEx := fmt.Sprintf(`%s %s`, script, needle)

	startScript := exec.Command("powershell", "-Command", scriptToEx)

	outputOfStd, err := startScript.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		log.Fatal("Unpredictable behavior happens while returning a StdoutPipe!")
	}

	StderrPipe, err := startScript.StderrPipe()

	if err != nil {
		fmt.Println(err)
		log.Fatal("Unpredictable behavior happens while returning a StderrPipe!")
	}

	err = startScript.Start()

	if err != nil {
		fmt.Println("To preform the init command you must be opening your cmd/powershell as an administrator!")
		log.Fatal("Unpredictable behavior happens while starting the process!")
	}

	outputBuffer := bufio.NewScanner(outputOfStd)
	errBuffer := bufio.NewScanner(StderrPipe)

	resultBufer := ""

	go func() {
		for outputBuffer.Scan() {
			resultBufer += outputBuffer.Text() + "\n"
			//fmt.Println(outputBuffer.Text())
		}
		if err := outputBuffer.Err(); err != nil {
			log.Fatal("Unpredictable behavior happens while outputing default buffer")
		}
	}()

	go func() {
		for errBuffer.Scan() {
			//fmt.Println(errBuffer.Text())
		}
		if err := errBuffer.Err(); err != nil {
			log.Fatal("Unpredictable behavior happens while outputing Error buffer")
		}
	}()

	err = startScript.Wait()

	if err != nil {
		log.Fatal("Unpredictable behavior happens while wating process!")
	}

	return resultBufer

}
