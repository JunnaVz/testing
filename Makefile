all: build-service build-repository build-cmd

build-service:
	cd internal\services && go build -o service && mv services ..\..

build-repository:
	cd internal\repository && go build -o repository && mv repository ..\..

build-cmd:
	cd cmd && go build -o cmd && mv cmd ..

.PHONY: all build-service build-repository build-cmd

#Ход компиляции программы целиком:
#go build main.go
#main.exe


#Ход компиляции модулей:
#go build -o service
#go build -o repository
#go build -o cmd
#lab3.exe