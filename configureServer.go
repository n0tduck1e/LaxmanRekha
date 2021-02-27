package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"golang.org/x/crypto/ssh"
)

type sshSess struct {
	client *server
	conn   *ssh.Client
}

func newSSHSess(conn *ssh.Client, client *server) {
	ssh := &sshSess{
		client: client,
		conn:   conn,
	}
	ssh.ConfigureServer()
}

func (ssh *sshSess) ConfigureServer() {
	//ssh.DeployKeys()
	ssh.DeployFirewall()
}

func (ssh *sshSess) DeployKeys() {
	publicKeyBytes := generateKeyPair("id_rsa", "id_rsa.pub")

	// making assumption that ssh will always land in home directory on login
	//TODO change it for windows hosts
	command := fmt.Sprintf("ls -la .ssh || mkdir .ssh/")
	ssh.Cmd(command)

	command = fmt.Sprintf("echo '%v' >> .ssh/authorized_keys", string(publicKeyBytes))
	fmt.Println(command)
	ssh.Cmd(command)
}

func (ssh *sshSess) ConfigureSSHPrivateKeyOnly() {

}

func (ssh *sshSess) DeployFirewall() {
	ssh.Cmd("(ls /sbin/iptables || which iptables) || apt install iptables")
	fmt.Println("Configuring Iptables on the remote box...")

	cmd := fmt.Sprintf("echo %v | sudo -S iptables -L", ssh.client.SSHPass)
	tmp, _ := ssh.Cmd(cmd)
	fmt.Println(tmp)

}

func (ssh *sshSess) Cmd(cmd ...string) (string, error) {
	var command, input string
	session, err := ssh.conn.NewSession()
	defer session.Close()

	if len(cmd) > 1 {
		input = cmd[0]
		command = cmd[1]

		stdin, err := session.StdinPipe()
		if err != nil {
			panic(err)
		}
		io.WriteString(stdin, input)

	} else {
		command = cmd[0]
	}

	if err != nil {
		panic(err)
	}
	stderr, err := session.StderrPipe()
	if err != nil {
		log.Println(err)
	}
	stdout, err := session.StdoutPipe()
	if err != nil {
		log.Println(err)
	}
	session.Run(command)
	errorString, _ := ioutil.ReadAll(stderr)
	output, _ := ioutil.ReadAll(stdout)
	err = fmt.Errorf("%s", string(errorString))

	return string(output), err
}
