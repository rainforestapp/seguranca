.PHONY: all build clean seguranca

all: build

build: seguranca

clean:
	rm -f seguranca

seguranca: *.go
	GOOS=linux GOARCH=amd64 go build

release:  seguranca
	aws s3 cp seguranca s3://rnfrst-binaries/seguranca/seguranca --acl public-read;
	aws s3 cp seguranca s3://rnfrst-binaries/seguranca/seguranca-$(CIRCLE_SHA1) --acl public-read;

run:
	RF_SECURITY_GROUP_REGION="us-east-1" RF_SECURITY_GROUP="sg-b0574ba7" go run main.go