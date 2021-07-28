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

func setup(ip,username, pass, port string) bool{
	//client := newServer()
	fmt.Println(ip)
	fmt.Println(username)
	fmt.Println(pass)
	fmt.Println(port)
	//client.createSession(client.SSHPass)
	return true
}
/*
func (client *server) askDetails() {

	fmt.Scanf("%s", &client.IP)

	fmt.Scanf("%s", &client.SShPort)

	fmt.Scanf("%s", &client.Username)
	bytepwd, _ := term.ReadPassword(int(syscall.Stdin))
	client.SSHPass = string(bytepwd)
//		client.Username = "ubuntu"
//		client.IP = "192.168.0.104"
//		client.SShPort = "22"
//		client.SSHPass = "<REDACTED>"
//		client.createSession(client.SSHPass)
	client.createSession(client.SSHPass)
}
*/

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
