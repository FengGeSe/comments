package main

import (
	// This Service
	global "comments/global"
	lib "comments/lib"
	server "comments/svc/server"
)

func main() {
	// Update addresses if they have been overwritten by flags
	lib.SetPid(global.ProjectRealPath + "/runtime/pid")

	server.Run()
}
