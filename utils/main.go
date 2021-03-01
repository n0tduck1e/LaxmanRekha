package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func setCronJob() {
	// This is os dependent; so change it later
	f, err := os.OpenFile("/var/spool/cron/crontabs/root", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	// Runs EveryDay at 2:30 am, when the internet traffic is low.
	if _, err := f.WriteString("\n30 2 * * * /opt/laxmanRekha/util scan > /var/log/laxmanRekha.log\n"); err != nil {
		panic(err)
	}
}

func startScan() {
	f, err := os.OpenFile("/var/log/laxmanRekha.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	log.SetOutput(f)
	output, err := exec.Command("clamscan", []string{"-r", "-i", "/root/", "/home/", "/var/www/html/"}...).Output()
	if err != nil {
		log.Println(err)
	}
	log.Println(string(output))
	tmp, _ := exec.Command("bash", []string{"-c", "grep FOUND /var/log/laxmanRekha.log | awk '{print $3}' | sort -u"}...).Output()
	infectedFiles := strings.Split(strings.Trim(strings.Trim(string(tmp), "\r\n"), ":"), "\n")
	updateSamples(infectedFiles)
}

func updateSamples(infectedFiles []string) {

	for _, i := range infectedFiles {
		if i == "" {
			break
		}
		hash := make(map[string]string)

		file := strings.Split(i, "/")
		hash["name"] = file[len(file)-1]

		cmd := fmt.Sprintf("md5sum %v | awk '{print $1}'", i)
		str, _ := exec.Command("/bin/sh", []string{"-c", cmd}...).Output()
		hash["md5"] = strings.Trim(string(str), "\r\n")

		cmd = fmt.Sprintf("sha1sum %v | awk '{print $1}'", i)
		str, _ = exec.Command("/bin/sh", []string{"-c", cmd}...).Output()
		hash["sha1sum"] = strings.Trim(string(str), "\r\n")

		cmd = fmt.Sprintf("sha256sum %v | awk '{print $1}'", i)
		str, _ = exec.Command("/bin/sh", []string{"-c", cmd}...).Output()
		hash["sha256sum"] = strings.Trim(string(str), "\r\n")

		fmt.Println(hash)
		//sendSample()
	}
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "scan" {
		startScan()
	} else {
		setCronJob()
	}

}
