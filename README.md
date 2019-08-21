# POC Serverless-Storage

Project that aims to demonstrate how to achieve persistence within serverless technology in a simple way.
<br/>
The application is based on a digital wallet in which we are storing coins. It consists of 3 parts:  
* Interface and frontend: developed with **Golang**.
* Backend and services: developed using serverless technology deployed in **Amazon Web Services Lambda** using **Node.js**.
* Database: simple storage using JSON format in **Amazon Web Services S3**.

## Getting Started

Let's get started with POCServerless! To start using the project and be able to play with it you have to clone the repository under the *src* folder of the Go workspace of your computer. 
<br/>
You should also replace the false AWS values in the [credentials.go](credentials.go) file with your own.


Once the previous steps have been completed, it is now possible to execute it: 
```
go run main.go aws.go constants.go credentials.go functions.go 
```
Or build an executable and run it later from wherever we want:
```
go install POCServerless
```
Typically this executable is located under the */bin* folder of our workspace. Remember that it is very useful to include the */bin* folder in our own path. So we have access to the executable from anywhere:
```
POCServerless
```

## Prerequisites

### AWS prerequisites

* Create an AWS account and obtain your credentials (*id* and *secret*).

### Code prerequisites

* Setting up the Go environment [Go Environment](https://golang.org/doc/code.html#Organization).

* Include [AWS](https://github.com/aws/aws-sdk-go) dependency for Go.

```
go get github.com/aws/aws-sdk-go
```

## AWS Deployment

1. **Create a private S3 bucket.** It's important to keep their name and region name. We'll need to configure it later in the credentials. 

2. **Include the storage file in S3.** Load the file [data.json](aws/data.json) into the bucket.

3. **Create an execution role in IAM service.** This role is for the Lambda execution. The services to be allowed are S3 and CloudWatch Logs (debug only). 
 
4. **Create the Lambda functions.** Create the functions [saveWallet.js](aws/saveWallet.js) and [getWallet.js](aws/getWallet.js) using the previously created role. It is important to replace the existing bucket name in the functions with the previously created one.

5. **Replace credentials.** Besides our *idkey* and *secretkey* in [credentials.go](credentials.go) we also need to replace the arn values of our Lambda functions and the region name where they are hosted.

## Running
```
MBP-Dani:~ daniortiz$ POCServerless 

*****************************
**Welcome to POC Serverless**
*****************************

Choose one of the following options:
1. View the current status of the wallet
2. Save coins
3. Exit
> 
```
```
> 1
The wallet has 0.00 coins

Choose one of the following options:
1. View the current status of the wallet
2. Save coins
3. Exit
> 
```
```
> 2
Coins to save:
5
Saving 5 coins...
5 coins saved in your wallet!

Choose one of the following options:
1. View the current status of the wallet
2. Save coins
3. Exit
> 
```
```
> 2
Coins to save:
8
Saving 8 coins...
8 coins saved in your wallet!

Choose one of the following options:
1. View the current status of the wallet
2. Save coins
3. Exit
> 
```

```
> 1
The wallet has 13.00 coins

Choose one of the following options:
1. View the current status of the wallet
2. Save coins
3. Exit
> 3
See you soon!
```


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

