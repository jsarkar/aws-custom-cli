package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/service/autoscaling"
)

func main() {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	svc := autoscaling.New(sess)

	input := &autoscaling.DescribeAutoScalingGroupsInput{}

	result, err := svc.DescribeAutoScalingGroups(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case autoscaling.ErrCodeInvalidNextToken:
				fmt.Println(autoscaling.ErrCodeInvalidNextToken, aerr.Error())
			case autoscaling.ErrCodeResourceContentionFault:
				fmt.Println(autoscaling.ErrCodeResourceContentionFault, aerr.Error())
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
