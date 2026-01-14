package utils

import "os/exec"

func OpenWin(uri string) {
	exec.Command(`cmd`, `/c`, `start`, uri).Start()
}
func OpenMac(uri string) {
	exec.Command("open", uri).Run()
}
