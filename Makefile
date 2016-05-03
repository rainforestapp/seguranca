.PHONY: all build clean seguranca

all: build

build: run

clean:
	rm -f seguranca

run:
	RF_SECURITY_GROUP_REGION="us-east-1" RF_SECURITY_GROUP="sg-b0574ba7" go run main.go