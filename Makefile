BIN=PongEbiten
32BIT=_i386
64BIT=_x86_64
ARM=arm64
WINDOWS_EXTENSION=.exe
STRIP=strip
UPX=upx -9

all:
	go build .

run:
	go run .

# to do:  build-linux32 build-linuxARM
release: build-windows64 build-windows32 build-windowsARM build-linux64
	$(UPX) $(BIN)$(64BIT)$(WINDOWS_EXTENSION)
	$(UPX) $(BIN)$(32BIT)$(WINDOWS_EXTENSION)
	$(UPX) $(BIN)$(ARM)$(WINDOWS_EXTENSION)
	$(UPX) $(BIN)$(64BIT)

build-windows64:
	GOOS=windows GOARCH=amd64 go build -o $(BIN)$(64BIT)$(WINDOWS_EXTENSION) .

build-windows32:
	GOOS=windows GOARCH=386 go build -o $(BIN)$(32BIT)$(WINDOWS_EXTENSION) .

build-windowsARM:
	GOOS=windows GOARCH=$(ARM) go build -o $(BIN)$(ARM)$(WINDOWS_EXTENSION) .

build-linux64:
	GOOS=linux GOARCH=amd64 go build -o $(BIN)$(64BIT) .
	$(STRIP) $(BIN)$(64BIT)

build-linux32:
	GOOS=linux GOARCH=386 go build -o $(BIN)$(32BIT) .
	$(STRIP) $(BIN)$(32BIT)

build-linuxARM:
	GOOS=linux GOARCH=$(ARM) go build -o $(BIN)$(ARM) .
	$(STRIP) $(BIN)$(ARM)

clean:
	rm -f $(BIN)$(64BIT)$(WINDOWS_EXTENSION)
	rm -f $(BIN)$(32BIT)$(WINDOWS_EXTENSION)
	rm -f $(BIN)$(ARM)$(WINDOWS_EXTENSION)
	rm -f $(BIN)$(64BIT)
	rm -f $(BIN)$(32BIT)
	rm -f $(BIN)$(ARM)
