#!/bin/bash


installSnoopy() {

	printf $Y"[+] "$GREEN"Installing Snoopy\n"$NC
	mv '/tmp/utils/ld.so.preload' '/etc/ld.so.preload'
	mv '/tmp/utils/libsnoopy.so' '/usr/local/lib/libsnoopy.so'
	mv '/tmp/utils/libsnoopy.so.0' '/usr/local/lib/libsnoopy.so.0'
	mv '/tmp/utils/libsnoopy.so.0.0.0' '/usr/local/lib/libsnoopy.so.0.0.0'
	mv '/tmp/utils/snoopy.ini' '/etc/snoopy.ini'
	
	# Restarting Services for better ld to be preloaded
	
	printf "Restarting Apache2 Webserver...\n"

	systemctl restart apache2 2>/dev/null

	printf "Success! logging utility was sucessfuly installed.\n"
	printf "You can view your logs at /var/log/auth.log\n"
}

setUputils() {
	
	printf $Y"[+] "$GREEN"Setting Up LaxmanRekha\n"$NC
	mkdir -p /opt/laxmanRekha/
	mv /tmp/utils/scanner /opt/laxmanRekha/scanner
	bash -c /opt/laxmanRekha/scanner
	printf "Laxman Rekha SucessFully Installed on the box\n"
	printf "Scanning the box\n"
	bash -c "/opt/laxmanRekha/scanner scan"
}

checkCronJobs() {
  cronjobs=`ls -la /etc/cron* 2>/dev/null`
	if [ "$cronjobs" ]; then
  		echo -e "\e[00;31m[-] Cron jobs:\e[00m\n$cronjobs" 
 		echo -e "\n"
	fi
	nonRoot=`echo $cronjobs | grep -v root`
	if [ "$nonRoot"]; then
		echo -e "[00;31m[-] You have non root cronjobs running."
		echo -e "[00;31m[-] Please Make Sure files used to cronfiles are world writeable."
	fi
  
}

loggedonusrs() {
	loggedonusrs=`w 2>/dev/null`
	if [ "$loggedonusrs" ]; then
  		echo -e "\e[00;31m[-] Who else is logged on:\e[00m\n$loggedonusrs" 
  		echo -e "\n"
	fi

}

checkForAllRoots(){
	superman=`grep -v -E "^#" /etc/passwd 2>/dev/null| awk -F: '$3 == 0 { print $1}' 2>/dev/null`
	if [ "$superman" ]; then
  		echo -e "\e[00;31m[-] Super user account(s):\e[00m\n$superman"
  		echo -e "\n"
	fi
	nonRoot=`echo $superman | grep -v root`
	if [ "$nonRoot"]; then
		echo -e "\e[00;31m[-] There seems to be more than one root accounts." 
		echo -e "\e[00;31m[-] Please Scan the system for any potential backdoors or compromise indicator." 
	fi
}

installClamav() {
	which clamscan || apt install clamav -y
}


loggedonusrs
checkCronJobs
checkForAllRoots
installSnoopy
installClamav
setUputils
