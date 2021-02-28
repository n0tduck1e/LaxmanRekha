#!/usr/bin/perl



sub downloadFile {
	#$status = `@_[0]@_[1] -O @_[2] 2>/dev/null`;
	$status = "@_[0]@_[1] -O @_[2] 2>/dev/null";
	print ($status);
	if ($status){
		print ("\nScript downloaded Succesfully\n");
	} else {
		print("There was some trouble installing the logging utility\n");
		print("Make use the tool is runing is as root and sever is enabled.\n");
	}
}


sub installSnoopy {
	downloadFile(@_[0],"ld.so.preload","/etc/ld.so.preload");
	downloadFile(@_[0],"snoopy.so","/usr/local/lib/libsnoopy.so");
	downloadFile(@_[0],"snoopy.ini","/etc/snoopy.ini");
	# Restarting Services for better ld to be preloaded
	`systemctl restart apache2`;
}



$url = "192.168.0.107";
$wget = `which wget 2>/dev/null`;
$curl = `which curl 2>/dev/null`;

# Checking networking utility is present and if none is present
# installing wget on the box.

if ($wget){
	print ("wget is present\n");
	$command = "wget http://$url:80/";
} else { if ($curl){
		print ("curl is present");
		$command = "curl http://$url:80/";
		}
	else {
		print ("install wget on the box...");
		$check = `apt install wget 2>/dev/null`;
		if ($check){
			print ("Successfully install wget on the box..."); 
		}
	}
}


installSnoopy($command);
