CGO_ENABLED := 0

all: jc jk

jc:
	CGO_ENABLED=$(CGO_ENABLED) go build -v -o bin/jc jc.go

jk:
	CGO_ENABLED=$(CGO_ENABLED) go build -v -o bin/jk jk.go

docker: all
	docker build -t jc .
