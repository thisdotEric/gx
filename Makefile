PROGRAM_NAME := gx

# Directory where the executable will be installed
INSTALL_DIR := /usr/local/bin

# Set the target operating system and architecture for Windows
GOOS_WINDOWS := windows
GOARCH_WINDOWS := amd64

.PHONY: build
build:
	go build -o $(PROGRAM_NAME)

.PHONY: build-windows
build-windows:
	GOOS=$(GOOS_WINDOWS) GOARCH=$(GOARCH_WINDOWS) go build -o $(PROGRAM_NAME).exe

.PHONY: install
install: build
	sudo mv $(PROGRAM_NAME) $(INSTALL_DIR)

.PHONY: clean
clean:
	rm -f $(PROGRAM_NAME) $(PROGRAM_NAME).exe

.PHONY: release
release: build-windows
	zip $(PROGRAM_NAME)_$(GOOS_WINDOWS)_$(GOARCH_WINDOWS).zip $(PROGRAM_NAME).exe
