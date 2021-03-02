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
