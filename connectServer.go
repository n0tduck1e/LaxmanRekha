package main

import (
	"fmt"

	"golang.org/x/crypto/ssh"
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
	fmt.Println("Please Enter the following details")

	client.Username = "duckie"
	client.IP = "172.17.0.2"
	client.SShPort = "22"
	client.SSHPass = "assslayer69"
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
