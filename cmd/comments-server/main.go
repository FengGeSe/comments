package main

import (
	// This Service
	global "myservices/comments/global"
	lib "myservices/comments/lib"
	server "myservices/comments/svc/server"
)

func main() {
	// Update addresses if they have been overwritten by flags
	lib.SetPid(global.ProjectRealPath + "/runtime/pid")

	server.Run()
}
