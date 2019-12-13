package cmd

import (
	"fmt"
	"os/exec"
	"runtime"
)

func openWebBrowser(repoUrl string) error {

	fmt.Printf("Open %s...\n", repoUrl)

	var args []string
	switch runtime.GOOS {
	case "windows":
		args = []string{"rundll32.exe", "url.dll,FileProtocolHandler", repoUrl}
	case "darwin":
		args = []string{"open", repoUrl}
	case "linux":
		args = []string{"xdg-open", repoUrl}
	default:
		return fmt.Errorf("Browsing git hosting service on %s is not supported.\n", runtime.GOOS)
	}

	cmd := exec.Command(args[0], args[1:]...)
	if cmd != nil {
		cmd.Start()
	}

	return nil
}
