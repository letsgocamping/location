GOOS=linux go build -o main && zip main.zip main && aws lambda update-function-code --region us-east-1 --function-name lgc-midpoint --zip-file fileb://./main.zip
