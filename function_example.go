package main

import (
	"fmt"
)

// Assign a funtion to a type.
type Protocol struct {
	Name         string
	Job          string
	NumberOfPort int
}

func (p *Protocol) Does() {
	// I want it to say SMB works over Port 445
  // And, SSH works over port 22
	fmt.Printf("The %s Protocol performs %s", p.Name, p.Job)
	fmt.Printf(" and works over port %d", p.NumberOfPort)
	fmt.Println()
}


func main() {
	var smb Protocol
	smb.Name = "Server Message Block"
	smb.NumberOfPort = 445
	smb.Job = "File Sharing"
	smb.Does()

	ssh := Protocol{
		Name:         "Secure Shell",
		Job:          "Remote Access",
		NumberOfPort: 22,
	}

	ssh.Does()

}
