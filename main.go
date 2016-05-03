package main

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"log"
	"net/http"
	"os"
	"sync"
)

var wg sync.WaitGroup

type VmIps []string

func authSg(ip string, protocol string, svc *ec2.EC2) {

	// Call the DescribeInstances Operation
	params := &ec2.AuthorizeSecurityGroupIngressInput{
		GroupId:    aws.String(os.Getenv("RF_SECURITY_GROUP")),
		CidrIp:     aws.String(ip + "/32"),
		DryRun:     aws.Bool(false),
		FromPort:   aws.Int64(0),
		ToPort:     aws.Int64(65535),
		IpProtocol: aws.String(protocol),
	}

	_, err := svc.AuthorizeSecurityGroupIngress(params)

	wg.Add(-1)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		log.Println("Error: ", ip, err.Error())
		return
	}
}

func main() {
	// Create an EC2 service object in the "us-west-2" region
	// Note that you can also configure your region globally by
	// exporting the AWS_REGION environment variable
	svc := ec2.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})

	res, err := http.Get("https://app.rainforestqa.com/api/1/vm_stack")
	if err != nil {
		log.Panicln(err)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var data VmIps
	err = decoder.Decode(&data)
	if err != nil {
		log.Panicln(err)
	}

	for _, ip := range data {
		log.Println(ip)
		wg.Add(2)
		go authSg(ip, "tcp", svc)
		go authSg(ip, "udp", svc)
	}

	wg.Wait()
}
