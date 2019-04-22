# Narkwhal

![alt-text](https://i.pinimg.com/originals/74/68/f1/7468f1d665e551fad8eac0c9f97977e3.jpg)

Narkwhal goes through your AWS account using the AWS SDK looking for available EBS volumes in the specified region and deletes them.

## Getting Started

### Requirements
- Narkwhal makes use of the AWS credentials folder. More information on this at the following link, [AWS Credentials folder](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#creating-the-credentials-file)

- Go. This has been only tested and confirmed to work on 1.11 and newer

### Installing on Mac

Go get Narkwhal
```
go get github.com/alexMcosta/narkwhal
```

Then from `github.com/alexMcosta/narkwhal` do
```
$ go build
$ go install
```

Now run Narkwhal!:
```
narkwhal -h
```

If you see the following then it is working and tells you the commands
```         
Usage of narkwhal:
  -account string
        Lets you select witch AWS account you would like to make changes to (default "default")
  -region string
        Lets you select which region you would like to run Narkwhal on (default "us-east-1")
  -time string
        Lets you select the amount of time a volume has been available based on MS, seconds, and Hours (default "0s")
```

Example of EBS volumes found and successful removal:
```
$ ./narkwhal -region ap-northeast-1                                                          
account: default, region: ap-northeast-1
---------------------
vol-0466fbaa999d132a6
---------------------
Would you like to remove the above EBS Volumes? (y/n):
y
Successfully removed vol-0466fbaa999d132a6
```

Example of no available EBS volumes in the region:
```
$ ./narkwhal -region ap-northeast-1
account: default, region: ap-northeast-1
---------------------
~~~~~~~~~~~~~~~~()~~~~~~
EXITING: There are no available EBS volumes in the ap-northeast-1 region to remove
~~~~~~~~~~~~~~~~()~~~~~~
---------------------
```

### More about the flags

#### -account
The account flag takes the account name between the square brackets from the `.aws/credentials` file. For example, to get the following: 

```
[Narkwhal]
aws_access_key_id = IAMSUCHAVERYCOOLACCESSKEY
aws_secret_access_key = 53cRE74CcEs5M30Wc47
```

The above credentials would be used if you passed `Narkwhal` through the account flag.

#### -region
The region flag lets you specify the region to use in that standard format like `us-west-1`. The default is `us-east-1`.

#### -time
The time flag uses Cloudwatch to check if there was any activity from the EBS volumes since the time specified. It does this by checking the `Read Ops` metric.

The time is based off of the Go function time.Duration.

### Feature Roadmap

#### Minor features
- [X] Add flag for choosing which account in `.aws/credentials` folder.
- [X] Add flag for choosing region
- Add the ability to scour all regions
- Add the ability to scour all accounts
- Add flag for either choosing specific id's or all EBS volumes
- Cache session to cut on calls
- Have Narkwhal take a config file and have it run as a cronjob.

#### Major features
- [X] Add flag to filter based on time
- Make binaries to install on multiple platforms so it does not require go
- Notifications 

## Authors

* **Alex Costa** 

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details