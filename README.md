# eventing-api

## start up

### terminal

_note: postgres database has to be running_

```
# run the commands:
go mod tidy
go run main.go

# press `control+c` to shutdown (mac)
```

### vscode debugger

_note: postgres database has to be running_

There is a `.vscode/launch.json` file. Therefore the api can be started/stopped via `vscode` debugger.

### docker-compose

```
# start services
docker compose -f "docker-compose.yml" up -d --build

# drop services
docker compose -f "docker-compose.yml" down
```
