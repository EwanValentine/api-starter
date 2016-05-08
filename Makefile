BUILD=GOOS=linux ARCH=GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s" -a -installsuffix cgo -o 

default: build

build:
	go get
	$(BUILD) ./api .
	docker build -t "ewanvalentine:api-starter" .

run: 
	docker run -p 5000:5000 ewanvalentine:api-starter
