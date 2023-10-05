# eventing-api

## start up

### docker-compose (recommended)

```
# start services
docker compose -f "docker-compose.yml" up -d --build

# drop services
docker compose -f "docker-compose.yml" down
```

### terminal

_note: postgres database has to be running_

```
# run the commands:
cd src
go mod tidy
go run main.go

# press `control+c` to shutdown (mac)
```

### vscode debugger

_note: postgres database has to be running_

There is a `.vscode/launch.json` file. Therefore the api can be started/stopped via `vscode` debugger.

## try out

Run this `curl` command via terminal.

```
curl --location 'http://localhost:4001/api/event' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Event name",
    "date": "2023-04-20T14:00:00Z",
    "languages": ["English"],
    "videoQuality": ["1080p"],
    "audioQuality": ["High"],
    "invitees": [ "john.doh@gmail.com"],
    "Description": "My description."
}'
```
