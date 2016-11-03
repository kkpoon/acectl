
mkdist:
	mkdir -p dist

linux64: mkdist
	GOOS=linux GOARCH=amd64 go build -o dist/acectl-linux-amd64 main.go

linux32: mkdist
	GOOS=linux GOARCH=386 go build -o dist/acectl-linux-386 main.go

macOS: mkdist
	GOOS=darwin GOARCH=amd64 go build -o dist/acectl-macos main.go

windows: mkdist
	GOOS=windows GOARCH=386 go build -o dist/acectl-win32.exe main.go

all: mkdist linux64 linux32 macOS windows

clean:
	rm -rf dist
