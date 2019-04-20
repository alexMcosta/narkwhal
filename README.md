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
go build
```

Now run Narkwhal!:
```
./narkwhal
```

If you see the following then it is working and tells you the commands
```
Usage of ./narkwhal:
  -account string
    	Lets you select witch AWS account you would like to make changes to (default "default")
  -region string
    	Lets you select which region you would like to run Narkwhal on (default "us-east-1")
```

Example of EBS volumes found and successful removal:
```
$ ./narkwhal -region ap-northeast-1                                                                          2 ↵
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

### Feature Roadmap

#### Minor features
- [X] Add flag for choosing which account in `.aws/credentials` folder.
- [X] Add flag for choosing region
- Add the ability to scour all regions
- Add the ability to scour all accounts
- Add flag for either choosing specific id's or all EBS volumes
- Cache session to cut on calls
- Make binaries to install on multiple platforms so it does not require go
- Notifications 
- Have Narkwhal take a config file and have it run as a cronjob.

#### Major features
* Add flag to filter based on time
  * This will be harder then initially thought since AWS does not seem to have a data set to find out when a volume was last added to an EC2 instance. I am assuming using tags will be how I have to go about this.

## Authors

* **Alex Costa** 

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details