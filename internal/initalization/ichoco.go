package initalization

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
)

func InstallChoco() {
	chocoVersion, err := checkChoco()

	if err != nil {
		fmt.Println("choco is not installed!")
		fmt.Println("Installing Choco...")
		installationProcess()

	} else {
		fmt.Printf("Choco is found, and its version is %v", chocoVersion)
	}

}

func checkChoco() (string, error) {

	versionScript := getScript("checkChocoVersion.ps1")

	CheckVersion := exec.Command("powershell", "-ExecutionPolicy", "Bypass", "-File", versionScript)

	var stderr, stdout bytes.Buffer

	CheckVersion.Stdout = &stdout
	CheckVersion.Stderr = &stderr

	err := CheckVersion.Run()

	if err != nil {
		log.Fatal("Unpredictable behavior happens while executing version command!")
	}

	if len(stderr.Bytes()) != 0 {
		return "", errors.New("choco not found")
	}

	return stdout.String(), nil

}

func getScript(scriptName string) string {
	mainDir, err := os.Getwd()

	if err != nil {
		log.Fatal("Unpredictable behavior happens while fetching path!")
	}

	return path.Join(mainDir, `internal\util\scripts\`+scriptName)
}

func installationProcess() {

	installationScript := getScript(`installChoco.ps1`)

	startInstalling := exec.Command("powershell", "-Command", "Start-Process", "powershell", "-ArgumentList", fmt.Sprintf(`"-ExecutionPolicy Bypass -File %s"`, installationScript), "-Verb", "runAs")

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
