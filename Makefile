all: ohs

deps:
	go get github.com/gorilla/mux
	touch deps

ohs: deps
	go build -o ohs ./src

clean:
	rm -rf deps ohs