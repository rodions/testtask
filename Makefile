all: build
	docker run --name ws -d -p 1323:1323 webservice

build: webservice
	docker build --tag=webservice .

webservice: dependency
	GOARCH=386 GOOS=linux go build -ldflags='-s -w' -o notes

dependency:
	go get github.com/satori/go.uuid  github.com/boltdb/bolt/... 

clean:
	docker stop ws ; docker rm ws ; docker rmi webservice ; rm notes
