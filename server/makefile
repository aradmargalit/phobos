run:
	go fmt ./... && go build && ./server
	
test:
	richgo test -v ./...

coverage:
	richgo test -v ./... -coverprofile=cover.out && go tool cover -html=cover.out 

format:
	go fmt ./...

check:
	${GOPATH}/bin/golint ./... && ${GOPATH}/bin/staticcheck ./...
