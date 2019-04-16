# Narkwhal

![alt-text](https://i.pinimg.com/originals/74/68/f1/7468f1d665e551fad8eac0c9f97977e3.jpg)

Narkwhal goes through your AWS account using the AWS SDK looking for available EBS volumes in the specified region and deletes them.

## Getting Started

### Requirements
- Narkwhal makes use of the AWS credentials folder. More information on this at the following link, [AWS Credentials folder](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#creating-the-credentials-file)

### Installing on Mac

Install by doing the following:
```
wget https://github.com/alexMcosta/narkwhal/releases/download/0.1.0/narkwhal
```

Then from the file you installed it you can run
```
./narkwhal -h
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
~~~~~~~~~~~~~~~~~~~~~~~
EXITING: There are no available EBS volumes in the ap-northeast-1 region to remove
~~~~~~~~~~~~~~~~~~~~~~~
---------------------
```

### Feature Ideas

#### Turn application into CLI Application
- [Close need testing] Add flags for choosing which account in `.aws/credentials` folder.
- [X] Add flag for choosing region
- Make it so you can scour all regions
- Add flag for time stamp date
- Add flag for either choosing specific id's or all EBS volumes

## Authors

* **Alex Costa** 

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details