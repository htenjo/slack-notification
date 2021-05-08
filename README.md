# Slack PR Notification

This project retrieves open PRs from Github repositories and send them to slack channel through a webhook, just to keep an eye on them.

- Tools to use:
  - Language: Golang
  - Infrastructure:
    - AWS CLI V2 
    - AWS Lambda
  - Github Account with repositories

- AWS CLI v2:
    - Installing from terminal:
    ```
    $ curl "https://awscli.amazonaws.com/AWSCLIV2.pkg" -o "AWSCLIV2.pkg"
    sudo installer -pkg AWSCLIV2.pkg -target /
    ```
    - Checking version:   
    ```
    $ which aws
    --> /usr/local/bin/aws 
    $ aws --version
    --> aws-cli/2.2.3 Python/3.8.8 Darwin/20.3.0 exe/x86_64 prompt/off
    ```
    - Additional info on [AWS Docs](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2-mac.html)