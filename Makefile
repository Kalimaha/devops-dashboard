build:
	mkdir -p functions
	
	go get github.com/aws/aws-lambda-go/events
	go get github.com/aws/aws-lambda-go/lambda
	go get github.com/Kalimaha/ginkgo/reporter
	go get github.com/onsi/ginkgo
	go get github.com/onsi/gomega
	go get github.com/google/go-github/github
	go get golang.org/x/oauth2

	(cd ./functions && go test)
	go build -o functions/vm-github-pull-requests ./...