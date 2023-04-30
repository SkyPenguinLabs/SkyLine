all:
	go build -o skyline main.go
	sudo mv skyline /usr/bin