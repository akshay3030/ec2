package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func list() (response []string) {
	// Call to get detailed information on each instance
	result, err := ec2Svc.DescribeInstances(nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	var instances []string
	for i := range result.Reservations {
		res, _ := json.Marshal(result.Reservations[i].Instances[0].InstanceId)
		instances = append(instances, string(res))
	}

	return instances

}

func listByName(name string) (response []string) {
	input := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag:Name"),
				Values: []*string{
					aws.String(strings.Join([]string{"*", name, "*"}, "")),
				},
			},
		},
	}

	result, err := ec2Svc.DescribeInstances(input)
	if err != nil {
		fmt.Println("Error", err)
	}
	var instances []string
	for i := range result.Reservations {
		res, _ := json.Marshal(result.Reservations[i].Instances[0].InstanceId)
		instances = append(instances, string(res))
	}

	return instances

}
