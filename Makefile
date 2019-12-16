all: test bin

bin:
	go build -o spot-instance-advisor github.com/AliyunContainerService/spot-instance-advisor/cmd/spot-instance-advisor

# Run tests
test: fmt vet
	go test ./... -coverprofile cover.out

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

