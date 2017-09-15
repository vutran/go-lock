package main

import (
	"log"
	"os/exec"
	"runtime"
)

func main() {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "darwin":
		cmd = "/System/Library/CoreServices/Menu Extras/User.menu/Contents/Resources/CGSession"
		args = []string{"-suspend"}
	case "linux":
		cmd = "gnome-screensaver-command"
		args = []string{"-l"}
	case "windows":
		cmd = "rundll32.exe"
		args = []string{"user32.dll,LockWorkStation"}
	}

	if err := exec.Command(cmd, args...).Run(); err != nil {
		log.Fatal(err)
	}
}
