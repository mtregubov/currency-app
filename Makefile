build-srv:
	rm -rf http/build
	cd curr-ui && npm run build
	mv curr-ui/build http/build
	go build -o cursrv ./http/server.go

build-cli:
	go build -o curcli ./cmd/main.go

build: build-srv build-cli

build-docker:
	docker build -t currapp .

build-all: build-srv build-cli build-docker

run-cli:
	CURR_YEAR=2019 ./curcli

run-srv:
	./cursrv

run-docker-cli:
	docker run --rm -it -e CURR_YEAR=2019 -v `pwd`/data:/app/data currapp /app/curcli

run-docker-srv:
	docker rm -f cursrv
	docker run --name=cursrv -d -p 8080:8080 -v `pwd`/data:/app/data currapp /app/cursrv

stop-docker-srv:
	docker rm -f cursrv

clean:
	rm -f cursrv curcli
	rm -rf http/build
	rm -rf data/data.db
	docker rm -f cursrv
	docker rmi -f currapp
