# meta info
NAME := go-webapp-sample

# tasks
setup:
	go get github.com/Masterminds/glide

deps: setup
	glide install

update: setup
	glide update

test: deps
	go test &&(glide novendor)

run: deps
	go get github.com/codegangsta/gin
	gin -i
