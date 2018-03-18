package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func stop(instances string) {

	input := &ec2.StopInstancesInput{
		InstanceIds: []*string{
			aws.String(instances),
		},
		DryRun: aws.Bool(false),
	}

	result, err := ec2Svc.StopInstances(input)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Success", result)
	}

}
