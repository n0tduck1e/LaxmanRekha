package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"

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
	ssh.DeployKeys()
	ssh.OnBoxDefence()
	ssh.DeployFirewall()
}

func (ssh *sshSess) DeployKeys() {
	publicKeyBytes := generateKeyPair("id_rsa", "id_rsa.pub")

	// making assumption that ssh will always land in home directory on login
	//TODO change it for windows hosts
	fmt.Println(hiwhite("Deploying SSH keys on the box."))
	fmt.Println(white("Its better to use them over passwords."))
	fmt.Println(higreen("[+] Creating public directory"))
	command := fmt.Sprintf("ls -la .ssh || mkdir .ssh/")
	ssh.Cmd(command)

	command = fmt.Sprintf("echo '%v' >> .ssh/authorized_keys", string(publicKeyBytes))
	fmt.Println(higreen("[+] SSH Keys Sucessfully Depolyed"))
	ssh.Cmd(command)
}

func (ssh *sshSess) DeployFirewall() {
	ssh.Cmd("(ls /sbin/iptables || which iptables) || apt install iptables")
	fmt.Println("Configuring Iptables on the remote box...")

	//cmd := fmt.Sprintf("echo %v | sudo -S iptables -L", ssh.client.SSHPass)
	cmd := "ss -plunt | awk '{print $5}' | grep -v -E '(Local|127.0.0.1)' | awk -F':' '{print $NF}' | sort -u"
	tmp, _ := ssh.Cmd(cmd)
	tmp = strings.Trim(tmp, "\r\n")
	ports := strings.Split(tmp, "\n")

	//Basic IPTable Setup
	iptableRules := []string{
		//Dont break existing connections
		"iptables -A INPUT -m conntrack --ctstate ESTABLISHED,RELATED -j ACCEPT",
		"iptables -A OUTPUT -m conntrack --ctstate ESTABLISHED -j ACCEPT",
		//Allow Anything on the localhost; aaye its the home afterall
		"iptables -A INPUT -i lo -j ACCEPT",
		"iptables -A OUTPUT -o lo -j ACCEPT",
		// Set default policies
		"iptables --policy INPUT DROP",
		"iptables --policy FORWARD DROP",
	}

	fmt.Println(hiwhite("Running following commands"))
	for _, i := range iptableRules {
		cmd = fmt.Sprintf("echo %v | sudo -S %v", ssh.client.SSHPass, i)
		fmt.Println(hiyellow(i))
		ssh.Cmd(cmd)
	}

	fmt.Println(hiwhite("Poking holes in the firewwall now"))
	//poking holes in for listening ports now
	for _, i := range ports {
		pokeHole := fmt.Sprintf("echo %v | sudo -S iptables -A INPUT -p tcp --dport %v -m state --state NEW -j ACCEPT", ssh.client.SSHPass, i)
		ssh.Cmd(pokeHole)
		if i == "22" || i == "21" {
			cmd := fmt.Sprintf("echo %v | sudo -S iptables -A INPUT -p tcp --dport %v -m state --state NEW -m recent --update --seconds 60 --hitcount 4 -j DROP", ssh.client.SShPort, i)
			ssh.Cmd(cmd)
			cmd = fmt.Sprintf("echo %v | sudo -S iptables -A INPUT -p tcp --dport %v -m state --state NEW -m recent --set", ssh.client.SShPort, i)
			ssh.Cmd(cmd)
		}
	}
	fmt.Println(higreen("[+] Firewall Sucessfully Deployed"))

}

func (ssh *sshSess) OnBoxDefence() {
	fmt.Println(hiwhite("Deploying On Box Defences"))
	cmd := "scp"
	_, err := exec.Command(cmd, []string{"-i", "id_rsa", "-r", "utils", ssh.client.Username + "@" + ssh.client.IP + ":/tmp/"}...).Output()
	if err != nil {
		panic(err)
	}

	fmt.Println(green("[+] Tranfered files on the box"))
	fmt.Println(green("[+] Installing AV on the box"))
	fmt.Println(green("[+] Checking for misconfigurations"))
	cmd = fmt.Sprintf("echo %v | sudo -S bash /tmp/utils/setDefences.sh", ssh.client.SSHPass)
	output, _ := ssh.Cmd(cmd)
	fmt.Println(output)
	fmt.Println(green("[+] All Seems Good :)"))
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
	if len(errorString) > 2 {
		err = fmt.Errorf("%s", string(errorString))
	} else {
		err = nil
	}

	return string(output), err
}
