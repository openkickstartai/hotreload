# hotreload

Universal file watcher. Like nodemon for any language.

## Install

```bash
git clone https://github.com/openkickstartai/hotreload.git
cd hotreload && go build -o hotreload .
```

## Usage

```bash
hotreload --watch '*.go' --run 'go build'
hotreload  # reads .hotreload.yml
```

## Config (.hotreload.yml)

```yaml
watch: ['*.go','*.html']
ignore: [vendor/]
command: go run .
restart: true
debounce: 200
```

## Testing

```bash
go test -v ./...
```
