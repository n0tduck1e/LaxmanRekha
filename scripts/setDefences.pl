#!/usr/bin/perl

sub checkerr {
	if (@_[0]){
		print("We Ran into an error while doing")
	}
	exit();

}


sub installSnoopy {

	`mv /tmp/scripts/ld.so.preload /etc/ld.so.preload`;
	`mv /tmp/scripts/libsnoopy.so /usr/local/lib/libsnoopy.so`;
	`mv /tmp/scripts/snoopy.ini /etc/snoopy.ini`;
	
	# Restarting Services for better ld to be preloaded
	
	print("Restarting Apache2 Webserver...\n");

	#`systemctl restart apache2 2>/dev/null`;
	
	#checkerr($?);

	print("Success! logging utility was sucessfuly installed.");
	print("You can view your logs at /var/log/auth.log\n");
}

installSnoopy();
