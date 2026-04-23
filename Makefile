BIN_DIR = ${HOME}/.local/bin
SYSTEMD_DIR = ${HOME}/.config/systemd/user


.PHONY: build install uninstall clean

build:
	go build -o ./bin/file-classifier ./cmd/classifier/main.go

install: build
	mkdir -p $(SYSTEMD_DIR)
	mkdir -p $(BIN_DIR)
	cp ./bin/file-classifier $(BIN_DIR)
	sed -e "s|{{HOME_DIR}}|${HOME}|g" ./systemd/file-classifier.service > $(SYSTEMD_DIR)/file-classifier.service
	sed -e "s|{{HOME_DIR}}|${HOME}|g" ./systemd/file-classifier.path > $(SYSTEMD_DIR)/file-classifier.path

uninstall:
	rm $(BIN_DIR)/file-classifier
	rm $(SYSTEMD_DIR)/file-classifier.service
	rm $(SYSTEMD_DIR)/file-classifier.path

clean:
	rm -rf ./bin
