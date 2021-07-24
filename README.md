# flugel-job-test

## How to use

### Pre-requisites

- golang >= 1.13 | [Official Link](https://www.terraform.io/downloads.html)
- terraform >= 0.12.20 | [Official Link](https://golang.org/dl/)
- aws-cli | [Official Link](https://docs.aws.amazon.com/pt_br/cli/latest/userguide/install-linux.html)
- git

### Configure your aws-cli

Follow the official documentation for configure your AWS-CLI tool 
https://docs.aws.amazon.com/pt_br/cli/latest/userguide/cli-chap-configure.html

### Download the repository 

**Important:** Donwload the code in your GOPATH

```bash
git clone --branch dev https://github.com/carreirorco/flugel-job-test.git
cd flugel-job-test
```

### Download the dependencies

```bash
 go mod init
 go mod tidy
```

### For execute the tests

```bash
go test -v test/s3_test.go
```

### For create this IaC

```bash
terraform init
terraform plan
terraform apply
```

### For destroy this IaC

```bash
terraform destroy
```