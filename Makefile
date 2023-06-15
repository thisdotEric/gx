
PROGRAM_NAME := gitx

# Directory where the executable will be installed
INSTALL_DIR := /usr/local/bin

.PHONY: build
build:
	go build -o $(PROGRAM_NAME)

.PHONY: install
install: build
	sudo mv $(PROGRAM_NAME) $(INSTALL_DIR)

.PHONY: clean
clean:
	rm -f $(PROGRAM_NAME)