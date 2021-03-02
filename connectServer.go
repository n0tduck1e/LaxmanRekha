package main

import (
	"fmt"
	"syscall"

	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

type server struct {
	Username string
	IP       string
	SShPort  string
	SShKey   string
	SSHPass  string
	Config   *ssh.ClientConfig
}

func newServer() *server {
	client := &server{}
	return client

}

func (client *server) askDetails() {
	fmt.Println(higreen("Please Enter the following details"))
	fmt.Printf(hiyellow("Server IP >"))
	fmt.Scanf("%s", &client.IP)
	fmt.Printf(hiyellow("Username >"))
	fmt.Scanf("%s", &client.Username)
	fmt.Printf(hiyellow("Port >"))
	fmt.Scanf("%s", &client.SShPort)
	fmt.Printf(hiyellow("Password (will not be printed to screen)>"))
	fmt.Println()
	bytepwd, _ := term.ReadPassword(int(syscall.Stdin))
	client.SSHPass = string(bytepwd)
	/*
		client.Username = "ubuntu"
		client.IP = "192.168.0.104"
		client.SShPort = "22"
		client.SSHPass = "assslayer69"
		client.createSession(client.SSHPass)
	*/
	client.createSession(client.SSHPass)
}

func (client *server) createSession(password string) {
	client.Config = &ssh.ClientConfig{
		User: client.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	conn, err := ssh.Dial("tcp", client.IP+":"+client.SShPort, client.Config)
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	newSSHSess(conn, client)
}
