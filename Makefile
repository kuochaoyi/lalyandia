AUTH_NAME = auth_server
GAME_NAME = game_server

BUILD_AUTH = make build_auth
BUILD_GAME = make build_game

.PHONY: build
build: build_proto build_auth build_game

.PHONY: build_proto
build_proto:
	protoc -I=. -I=${GOPATH}/src/ --gogoslick_out=plugins=grpc:. messages/external/*.proto

.PHONY: build_auth
build_auth:
	go build -o ${AUTH_NAME} ./auth/main.go

.PHONY: build_game
build_game:
	go build -o ${GAME_NAME} ./game/main.go

.PHONY: auth
auth:
	go run ./auth/main.go

.PHONY: game
game:
	go run ./game/main.go

.PHONY: watch_auth
watch_auth: build_proto
	reflex -r '\.(go|json|toml)$$' -R '^vendor/' -s -- sh -c '$(BUILD_AUTH) && ./${AUTH_NAME}'

.PHONY: watch_game
watch_game: build_proto
	reflex -r '\.(go|json|toml)$$' -R '^vendor/' -s -- sh -c '$(BUILD_GAME) && ./${GAME_NAME}'
