compile:
	# Creating the LaxmanScanner
	echo "Building the laxmanRekha"
	go build

	# Building the mother ship

	echo "Building the mothership"
	cd motherShip ; go build
	cd ..

	# Building the scanner

	echo "Building the Scanner"
	cd scanner; CGO_ENABLED=0 go build
	cd ..

	# moving things to place
	echo "Moving things to place"
	mv scanner/scanner utils/

clean:
	# Cleaning things Up
	echo "Cleaning the laxmanRekha"
	go clean
	rm id_rsa*

	# Cleaning the mother ship
	echo "Cleaning the mothership"
	cd motherShip ; go clean
	cd ..

	# Cleaning the scanner
	echo "Cleaning the Scanner"
	cd scanner; go clean
	cd ..

	echo "Cleaning Moved Things"
	rm utils/scanner
