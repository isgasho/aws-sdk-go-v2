module github.com/aws/aws-sdk-go-v2/service/s3/internal/configtesting

go 1.15

require (
	github.com/aws/aws-sdk-go-v2/config v0.2.0
	github.com/aws/aws-sdk-go-v2/service/s3 v0.27.0
	github.com/aws/aws-sdk-go-v2 v0.27.0
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v0.2.0
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v0.2.0
	github.com/aws/aws-sdk-go-v2/service/sts v0.27.0
	github.com/aws/aws-sdk-go-v2/credentials v0.1.2
	github.com/aws/aws-sdk-go-v2/ec2imds v0.1.2
)

replace (
	github.com/aws/aws-sdk-go-v2 => ../../../../
	github.com/aws/aws-sdk-go-v2/config => ../../../../config/
	github.com/aws/aws-sdk-go-v2/service/s3 => ../../
)

replace github.com/aws/aws-sdk-go-v2/credentials => ../../../../credentials/

replace github.com/aws/aws-sdk-go-v2/ec2imds => ../../../../ec2imds/

replace github.com/aws/aws-sdk-go-v2/service/internal/s3shared => ../../../../service/internal/s3shared/

replace github.com/aws/aws-sdk-go-v2/service/sts => ../../../../service/sts/

replace github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding => ../../../../service/internal/accept-encoding/
