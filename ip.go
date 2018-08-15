package main

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func getIP(args []string) {

	var instanceIds []*string

	for _, inst := range args {
		if strings.HasPrefix(inst, "i-") {
			instanceIds = append(instanceIds, aws.String(inst))
		} else {
			var id []string
			id = listByName(inst)

			if len(id) != 0 {
				instanceIds = append(instanceIds, aws.String(trimQuotes(id[0])))
			}
		}
	}

	if len(instanceIds) != 0 {
		input := &ec2.DescribeInstancesInput{
			InstanceIds: instanceIds,
		}

		result, err := ec2Svc.DescribeInstances(input)
		if err != nil {
			fmt.Println("Error", err)
		}
		fmt.Println("instance ID | IP | hostname")
		for _, reservation := range result.Reservations {
			for _, i := range reservation.Instances {
				id := i.InstanceId
				publicDNS := i.PublicDnsName
				publicIP := i.PublicIpAddress
				privateDNS := i.PrivateDnsName
				privateIP := i.PrivateIpAddress

				fmt.Printf("%s | %s | %s \n", *id, *publicDNS, *publicIP)
				fmt.Printf("%s | %s | %s \n", *id, *privateDNS, *privateIP)

			}
		}
	} else {
		fmt.Println("No instances found")
	}

}
