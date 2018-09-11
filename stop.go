package main

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func ec2StopRequest(instanceIds []*string) {
	input := &ec2.StopInstancesInput{
		InstanceIds: instanceIds,
		DryRun:      aws.Bool(false),
	}

	result, err := ec2Svc.StopInstances(input)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Success", result)
	}
}

func stop(args []string) {

	var instanceIds []*string

	for _, inst := range args {
		if strings.HasPrefix(inst, "i-") {
			instanceIds = append(instanceIds, aws.String(inst))
		} else {
			var id []string
			id = listByName(inst)

			if len(id) != 0 {
				for _, i := range id {
					instanceIds = append(instanceIds, aws.String(trimQuotes(i)))
				}
			}
		}
	}

	switch len(instanceIds) {
	case 0:
		fmt.Println("No instances to found")
	case 1:
		ec2StopRequest(instanceIds)
	default:
		for _, v := range instanceIds {
			fmt.Printf("%s \n", *v)
		}
		fmt.Println("Type yes to confirm")
		if confirmation() {
			ec2StopRequest(instanceIds)
		}
	}
}
