module github.com/aws/aws-sdk-go-v2/service/ec2

go 1.15

require (
	github.com/aws/aws-sdk-go-v2 v0.26.1-0.20201013182106-0fbdffd558c0
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v0.0.0-20201013182106-0fbdffd558c0
	github.com/awslabs/smithy-go v0.1.2-0.20201012175301-b4d8737f29d1
)

replace github.com/aws/aws-sdk-go-v2 => ../../

replace github.com/aws/aws-sdk-go-v2/service/internal/presigned-url => ../../service/internal/presigned-url/
