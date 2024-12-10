package main

import (
	"fmt"
	"log"
	"golang.org/x/sys/windows"
)

func isProcessElevated() bool {
	// Open the current process token
	var token windows.Token
	err := windows.OpenCurrentProcessToken(windows.TOKEN_QUERY, &token)
	if err != nil {
		log.Fatal("Error opening process token:", err)
	}
	defer token.Close()

	// Check if the token has the "SeDebugPrivilege" privilege (elevated privilege)
	privileges, err := token.GetTokenPrivileges()
	if err != nil {
		log.Fatal("Error getting token privileges:", err)
	}

	for _, privilege := range privileges {
		if privilege.Luid == windows.SeDebugPrivilege.Luid {
			return true
		}
	}
	return false
}

func main() {
	if isProcessElevated() {
		fmt.Println("The process is running with elevated privileges!")
	} else {
		fmt.Println("The process is NOT running with elevated privileges.")
	}
}