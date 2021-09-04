BINARY_NAME = archvium 

build:
	go build -v

run: build
	./$(BINARY_NAME) 

install: build
	sudo cp ./$(BINARY_NAME) /usr/bin/

update_mod:
	go build -v -mod=mod

# (build but with a smaller binary)
dist:
	go build -ldflags="-w -s" -gcflags=all=-l -v

# (even smaller binary)
pack: dist
	upx ./$(BINARY_NAME)

lint:
	revive -formatter friendly -config revive.toml ./... 

# fiber dev (with hot reload)
dev:
	fiber dev

# open inside a screen
screen:
	screen -dmS archvium_back make run
