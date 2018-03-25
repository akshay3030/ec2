package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/fatih/color"
)

func list() {
	// Call to get detailed information on each instance
	result, err := ec2Svc.DescribeInstances(nil)
	if err != nil {
		fmt.Println("Error", err)
	}
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			id := *instance.InstanceId
			var name string
			status := *instance.State.Name
			for _, tag := range instance.Tags {
				if *tag.Key == "Name" {
					name = *tag.Value
				}
			}
			switch status {
			case "running":
				fmt.Printf("%s | %s | %s \n", id, color.GreenString(status), name)
			case "stopped":
				fmt.Printf("%s | %s | %s \n", id, color.RedString(status), name)
			case "pending":
				fmt.Printf("%s | %s | %s \n", id, color.YellowString(status), name)
			case "stopping":
				fmt.Printf("%s | %s | %s \n", id, color.YellowString(status), name)
			default:
				fmt.Printf("%s | %s | %s \n", id, status, name)
			}
		}
	}
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
