package main

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/service/ec2"
)

// Start Comment
func start(args []string) {

	var instanceIds []*string

	for _, inst := range args[2:] {
		if strings.HasPrefix(inst, "i-") {
			instanceIds = append(instanceIds, aws.String(inst))
		} else {
			id := listByName(inst)
			if len(id) > 1 {
				fmt.Println("Start all these instances?")
				for _, v := range id {
					fmt.Println(trimQuotes(v))
				}
				fmt.Println("Type yes to confirm")
				if confirmation() {
					for _, v := range id {
						instanceIds = append(instanceIds, aws.String(trimQuotes(v)))
					}
				}
			} else {
				instanceIds = append(instanceIds, aws.String(trimQuotes(id[0])))
			}
		}
	}

	input := &ec2.StartInstancesInput{
		InstanceIds: instanceIds,
		DryRun:      aws.Bool(false),
	}

	result, err := ec2Svc.StartInstances(input)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Success", result)
	}

}
