package main

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func getInstanceVolumes(instances []*string) *ec2.DescribeInstancesOutput {
	input := &ec2.DescribeInstancesInput{
		InstanceIds: instances,
	}

	res, err := ec2Svc.DescribeInstances(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
	}

	return res

}

func backup(args []string) {
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
		inst := getInstanceVolumes(instanceIds)
		for i := range inst.Reservations {
			for _, ins := range inst.Reservations[i].Instances {
				for v := range ins.BlockDeviceMappings {
					volID := *ins.BlockDeviceMappings[v].Ebs.VolumeId

					//fmt.Println("Starting snapshot for volume ", volId),

					// Snapshot information
					input := &ec2.CreateSnapshotInput{
						Description: aws.String("Backup"),
						VolumeId:    aws.String(volID),
						TagSpecifications: []*ec2.TagSpecification{
							{
								ResourceType: aws.String("snapshot"),
								Tags: []*ec2.Tag{
									{
										Key:   aws.String("instance"),
										Value: ins.InstanceId,
									},
								},
							},
						},
					}

					// Request snapshot
					result, err := ec2Svc.CreateSnapshot(input)
					if err != nil {
						if aerr, ok := err.(awserr.Error); ok {
							switch aerr.Code() {
							default:
								fmt.Println(aerr.Error())
							}
						} else {
							fmt.Println(err.Error())
						}
						return
					}

					fmt.Println(result)
				}
			}
		}
	}

}
