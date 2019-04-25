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

### Flags

#### [-account](https://github.com/alexMcosta/narkwhal/blob/master/documentation/account.md)

#### [-regions](https://github.com/alexMcosta/narkwhal/blob/master/documentation/regions.md)

#### -time
The time flag uses Cloudwatch to check if there was any activity from the EBS volumes since the time specified and it does this by checking the `Read Ops` metric. 
Time does not check for the last time the volume was attached but rather the last time the volume showed any kind of read activity. That being said, the volume 
could have been attached to an EC2 instance that was then stopped for a month. If that EC2 instance was then terminated, without being activated again, the volume 
would be up for deletion if a time flag of `48h` was passed since it would then be available and also show no signs of activity for longer then 48 hours.

The input is based off of the Go function `time.Duration()`.

### Feature Roadmap

#### Minor features
- Add the ability to scour some or all accounts
- Add flag for either choosing specific id's or all EBS volumes
- Cache session to cut on calls
- Have Narkwhal take a config file and have it run as a cronjob.
- Silent mode

#### Major features
- Make binaries to install on multiple platforms so it does not require go
- Notifications 

## Authors

* **Alex Costa** 

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details