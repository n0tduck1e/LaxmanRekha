package main

import (
	"io/ioutil"
	"log"

	"golang.org/x/crypto/ssh"
)

type sshSess struct {
	client  *server
	session *ssh.Session
}

func newSSHSess(session *ssh.Session, client *server) {
	ssh := &sshSess{
		client:  client,
		session: session,
	}
	ssh.ConfigureServer()
}

func (ssh *sshSess) ConfigureServer() {
	ssh.DeployKeys()
}

func (ssh *sshSess) DeployKeys() {

}

func (ssh *sshSess) Cmd(cmd string) (string, string) {
	stderr, err := ssh.session.StderrPipe()
	stdout, err := ssh.session.StdoutPipe()
	if err != nil {
		log.Println(err)
	}
	ssh.session.Run(cmd)
	errorString, _ := ioutil.ReadAll(stderr)
	output, _ := ioutil.ReadAll(stdout)
	return string(output), string(errorString)
}
