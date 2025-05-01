package commands

import (
	"fmt"
	"github.com/feralonsomaccari/my-cli-tools/utils"
	"os/exec"
	"strings"
)

func OpenRepository() {
	repoUrl, err := exec.Command("git", "remote", "get-url", "origin").CombinedOutput()
	if err != nil {
		fmt.Println("Errors:", err)
		return
	}

	formattedUrl := utils.SshToHttps(string(repoUrl))

	exec.Command("open", formattedUrl).Start()

}

func OpenPR() {

	repoUrl, err := exec.Command("git", "remote", "get-url", "origin").CombinedOutput()
	if err != nil {
		fmt.Print("Errors:", err)
		return
	}

	branchName, err := exec.Command("git", "branch", "--show-current").CombinedOutput()
	if err != nil {
		fmt.Print("Errors:", err)
		return
	}

	formattedUrl := utils.SshToHttps(string(repoUrl))
	branchUrl := formattedUrl + "/pull/" + strings.TrimSpace(string(branchName))

	exec.Command("open", branchUrl).Start()

}
