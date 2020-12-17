# Simple S3 upload utility

Simple utility written in golang to upload files to AWS S3.

AWS S3 credentials are taken from environment variables `AWS_ACCESS_KEY` and `AWS_SECRET_KEY`.  AWS region defaults to `eu-north-1`.  For example:

````
$ export AWS_SECRET_KEY=xxxx
$ export AWS_ACCESS_KEY=zzzz
$ ./simple-s3 --bucket demo-bucket --name demofile --file testfile
$ aws s3 ls s3://demo-bucket/ --region eu-north-1
2020-12-17 15:25:12       1220 demofile
````

## Static binary

Static binary is compiled with the following options:

````
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o simple-s3 simple-s3.go
````
