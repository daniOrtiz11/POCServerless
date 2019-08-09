package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

/*
Lambda is a function to call Lambda Services
*/
func Lambda(arn string, params string) string {
	resp := ""
	sess := getCredentials()
	svc := lambda.New(sess)
	var bytespayload []byte
	bytespayload, err := json.Marshal(params)
	input := &lambda.InvokeInput{
		FunctionName:   aws.String(arn),
		InvocationType: aws.String(responseParam),
		LogType:        aws.String(logParam),
		Payload:        bytespayload,
	}

	result, err := svc.Invoke(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			//check errors
			switch aerr.Code() {
			case lambda.ErrCodeServiceException:
				fmt.Println(lambda.ErrCodeServiceException, aerr.Error())
			case lambda.ErrCodeResourceNotFoundException:
				fmt.Println(lambda.ErrCodeResourceNotFoundException, aerr.Error())
			case lambda.ErrCodeInvalidRequestContentException:
				fmt.Println(lambda.ErrCodeInvalidRequestContentException, aerr.Error())
			case lambda.ErrCodeRequestTooLargeException:
				fmt.Println(lambda.ErrCodeRequestTooLargeException, aerr.Error())
			case lambda.ErrCodeUnsupportedMediaTypeException:
				fmt.Println(lambda.ErrCodeUnsupportedMediaTypeException, aerr.Error())
			case lambda.ErrCodeTooManyRequestsException:
				fmt.Println(lambda.ErrCodeTooManyRequestsException, aerr.Error())
			case lambda.ErrCodeInvalidParameterValueException:
				fmt.Println(lambda.ErrCodeInvalidParameterValueException, aerr.Error())
			case lambda.ErrCodeEC2UnexpectedException:
				fmt.Println(lambda.ErrCodeEC2UnexpectedException, aerr.Error())
			case lambda.ErrCodeSubnetIPAddressLimitReachedException:
				fmt.Println(lambda.ErrCodeSubnetIPAddressLimitReachedException, aerr.Error())
			case lambda.ErrCodeENILimitReachedException:
				fmt.Println(lambda.ErrCodeENILimitReachedException, aerr.Error())
			case lambda.ErrCodeEC2ThrottledException:
				fmt.Println(lambda.ErrCodeEC2ThrottledException, aerr.Error())
			case lambda.ErrCodeEC2AccessDeniedException:
				fmt.Println(lambda.ErrCodeEC2AccessDeniedException, aerr.Error())
			case lambda.ErrCodeInvalidSubnetIDException:
				fmt.Println(lambda.ErrCodeInvalidSubnetIDException, aerr.Error())
			case lambda.ErrCodeInvalidSecurityGroupIDException:
				fmt.Println(lambda.ErrCodeInvalidSecurityGroupIDException, aerr.Error())
			case lambda.ErrCodeInvalidZipFileException:
				fmt.Println(lambda.ErrCodeInvalidZipFileException, aerr.Error())
			case lambda.ErrCodeKMSDisabledException:
				fmt.Println(lambda.ErrCodeKMSDisabledException, aerr.Error())
			case lambda.ErrCodeKMSInvalidStateException:
				fmt.Println(lambda.ErrCodeKMSInvalidStateException, aerr.Error())
			case lambda.ErrCodeKMSAccessDeniedException:
				fmt.Println(lambda.ErrCodeKMSAccessDeniedException, aerr.Error())
			case lambda.ErrCodeKMSNotFoundException:
				fmt.Println(lambda.ErrCodeKMSNotFoundException, aerr.Error())
			case lambda.ErrCodeInvalidRuntimeException:
				fmt.Println(lambda.ErrCodeInvalidRuntimeException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
		return ""
	}
	//parser result from AWS
	s := string(result.Payload)
	mapBody := make(map[string]interface{})
	err2 := json.Unmarshal([]byte(s), &mapBody)
	if err2 != nil {
		log.Fatal(awsError)
	}
	var statusCode float64
	body := ""
	statusCode = 0
	for _, value := range mapBody {
		switch v := value.(type) {
		case float64:
			statusCode = value.(float64)
		case string:
			body = value.(string)
		default:
			println(v)
			log.Fatal(errorRespAws)
		}
	}
	if statusCode != 0 && body != "" {
		if statusCode == 200 || statusCode == 201 || statusCode == 202 {
			resp = body
		} else {
			resp = koC
			fmt.Println(awsError)
		}
	}

	return resp
}

func getCredentials() *session.Session {
	var AnonymousCredentials = credentials.NewStaticCredentials(idiam, secretiam, "")
	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: AnonymousCredentials,
		Region:      aws.String(regionaws),
	}))
	return sess
}
