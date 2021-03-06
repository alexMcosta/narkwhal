# Narkwhal

![alt-text](https://i.pinimg.com/originals/74/68/f1/7468f1d665e551fad8eac0c9f97977e3.jpg)

Narkwhal is a CLI that goes through an AWS account and removes parked EBS volumes based on user specified criteria.

## Getting Started

### Requirements
- Narkwhal makes use of the AWS credentials folder. More information on this at the following link, [AWS Credentials folder](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#creating-the-credentials-file)

- Go installed. This has been only tested and confirmed to work on 1.11 and newer

### Installing on Mac

Go get Narkwhal
```
go get github.com/alexMcosta/narkwhal
```

Then from `github.com/alexMcosta/narkwhal` do
```
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
  -regions string
        Lets you select which regions you would like to run Narkwhal on (default "us-east-1")
  -time string
        Lets you select the amount of time a volume has been available based on MS, seconds, and Hours (default "0s")
```

### Flags

#### [Account Flag](https://github.com/alexMcosta/narkwhal/blob/master/documentation/account.md)
The `-account` flag takes the account name between the square brackets from the `.aws/credentials` file and uses it with the AWS SDK.

#### [Regions Flag](https://github.com/alexMcosta/narkwhal/blob/master/documentation/regions.md)
The `-regions` flag lets you choose which region or regions you would like to select.

#### [Time Flag](https://github.com/alexMcosta/narkwhal/blob/master/documentation/time.md)
The `-time` flag uses Cloudwatch to check if there was any activity from the EBS volumes since the time specified and it does this by checking both the `Read Ops` and `Write Ops` metrics.

### Feature Roadmap

#### Minor features
- Add the ability to choose which volumes to delete.
- Add ability to filter by EBS type

#### Major features
- Make binaries to install on multiple platforms so it does not require go.
- Notifications.
- Create a Lambda version

## Authors

* **Alex Costa** 

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details