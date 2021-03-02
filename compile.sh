# Creating the LaxmanScanner

echo "Building the laxmanRekha"
cd laxmanRekha/ ; go build
cd ..

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
mv scanner/scanner laxmanRekha/utils/
