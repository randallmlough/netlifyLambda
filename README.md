# Netlify + AWS GO Lambda (Locally!)

## Prerequisites
* [Docker](https://docs.docker.com/install)
* [AWS Command Line Interface](https://docs.aws.amazon.com/cli/latest/userguide/installing.html) - and configured 
* [AWS SAM Local](https://github.com/awslabs/aws-sam-local#windows-linux-macos-with-npm-recommended)

## Overview
This repo is an attempt to get AWS Lambda GO functions to work locally with Hugo with Netlify's protocol structure. This is a current **work in progress**.

## Runing Locally
:warning: Make sure to install all the [Prerequisites](#prerequisites). On Mac
OSX and Windows, ensure that the Docker VM is running.

## Installation & usage
```bash
# Install npm dependencies
npm install -g aws-sam-local

# Clone the skeleton
cd $GOPATH/src 
git clone https://github.com/randallmlough/netlifyLambda.git golambda

# Compile the sample lambdas
cd golambda/
make

# Invoke the "apigw" sample lambda locally at localhost:3000
make start

```

Once launched go to http://localhost:9000/index.html to test some of the functions. **NOTE:** http://localhost:9000/ only will result in an error, need to go to /index.html