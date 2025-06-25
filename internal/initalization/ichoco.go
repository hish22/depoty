//go:build windows

package initalization

import (
	"bufio"
	"bytes"
	"depoty/internal/util/scripts"
	"errors"
	"fmt"
	"log"
	"os/exec"
)

func InstallChoco() {
	chocoVersion, err := checkChoco()

	if err != nil {
		fmt.Println("choco is not installed!")
		fmt.Println("Installing Choco...")
		installationProcess()

	} else {
		fmt.Printf("Choco is found, and its version is %v \n", chocoVersion)
	}

}

func checkChoco() (string, error) {

	CheckVersion := exec.Command("powershell", "-ExecutionPolicy", "Bypass", "-Command", scripts.CheckVersionScript)

	var stderr, stdout bytes.Buffer

	CheckVersion.Stdout = &stdout
	CheckVersion.Stderr = &stderr

	err := CheckVersion.Start()

	if err != nil {
		fmt.Printf("Error executing command: %v\nOutput: %s\n", err, "choco --version")
		return "", err
	}

	err = CheckVersion.Wait()

	if err != nil {
		fmt.Printf("Error executing command: %v\nOutput: %s\n", err, "choco --version")
		return "", err
	}

	if len(stderr.Bytes()) != 0 {
		return "", errors.New("choco not found")
	}

	return stdout.String(), nil

}

func installationProcess() {

	// startInstalling := exec.Command("powershell", "-Command", "Start-Process", "powershell", "-ArgumentList", fmt.Sprintf(`"-ExecutionPolicy Bypass -Command %s"`, scripts.InstallChocoScript), "-Verb", "runAs")
	startInstalling := exec.Command("powershell", "-ExecutionPolicy", "Bypass", "-Command", scripts.InstallChocoScript)

	outputOfInstallation, err := startInstalling.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		log.Fatal("Unpredictable behavior happens while returning a StdoutPipe!")
	}

	InstallationerrPipe, err := startInstalling.StderrPipe()

	if err != nil {
		fmt.Println(err)
		log.Fatal("Unpredictable behavior happens while returning a StderrPipe!")
	}

	err = startInstalling.Start()

	if err != nil {
		fmt.Println("To preform the init command you must be opening your cmd/powershell as an administrator!")
		log.Fatal("Unpredictable behavior happens while starting the process!")
	}

	outputBuffer := bufio.NewScanner(outputOfInstallation)
	errBuffer := bufio.NewScanner(InstallationerrPipe)

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

	err = startInstalling.Wait()

	if err != nil {
		log.Fatal("Unpredictable behavior happens while wating process!")
	}

	fmt.Println("Choco Installed successfully")

}
