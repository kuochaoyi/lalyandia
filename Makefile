.PHONY: build auth game

AUTH_NAME = auth_server
GAME_NAME = game_server

build:
	go build -o ${AUTH_NAME} ./auth/main.go
	go build -o ${GAME_NAME} ./game/main.go

auth:
	go run ./auth/main.go

game:
	go run ./game/main.go
