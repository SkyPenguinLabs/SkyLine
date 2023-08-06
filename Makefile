
install:
	go build -o SkyLine main.go 
	sudo cp SkyLine /usr/bin
	sudo cp -r Modules /usr/share
uninstall:
	sudo rm /usr/bin/SkyLine
	sudo rm -rf /usr/share/Modules