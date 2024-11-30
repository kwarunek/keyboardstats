OUT := build/keyboardstats
SRC_FILES := $(wildcard ./cmd/*.go) $(wildcard ./internal/**/*.go)
VERSION := "1.1.1"
INSTALL_DIR := /opt/keyboardstats

codes:
	wget https://raw.githubusercontent.com/torvalds/linux/refs/heads/master/include/uapi/linux/input-event-codes.h -O /tmp/input-event-codes.h
	./tools/generate-codes.py /tmp/input-event-codes.h > internal/codes.go
	go fmt internal/codes.go


build:  # Make a production build (w/o debuginfo)
	go build -v -o ${OUT} -ldflags="-s -w -X main.version=${VERSION}" ./cmd

install: build
	sudo install -Dm755 ${OUT} ${INSTALL_DIR}/keyboardstats
	sudo install -Dm644 contrib/keyboardstats.service /etc/systemd/system/keyboardstats.service

	sudo systemctl enable keyboardstats
	sudo systemctl start keyboardstats

.PHONY: build codes install
