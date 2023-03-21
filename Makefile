
build:
	go build -o bin/tmux-music ./...

run: build
	./bin/tmux-music

install: uninstall build
	cp ./bin/tmux-music $${HOME}/bin

uninstall:
	rm $${HOME}/bin/tmux-music

clean:
	rm ./bin/tmux-music
