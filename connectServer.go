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
	Config   *ssh.ClientConfig
}

func newServer() *server {
	client := &server{}
	return client

}

func (client *server) askDetails() {
	fmt.Println("Please Enter the following details")

	client.Username = "duckie"
	client.IP = "127.0.0.1"
	client.SShPort = "22"
	client.createSession("assslayer69")
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
	if err != nil {
		panic(err)
	}
	sess, err := conn.NewSession()
	if err != nil {
		panic(err)
	}
	newSSHSess(sess, client)
}
