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

Example of EBS volumes found and successful removal:
```
$ narkwhal -account default -region ap-northeast-1 -time 24h                                                          
account: default, region: us-east-1, Not used within: 24h
---------------------
vol-0e4a10ca6d4e1fb09
---------------------
Would you like to remove the above EBS Volumes? (y/n): 
y
Successfully removed vol-0e4a10ca6d4e1fb09
```

Example of no available EBS volumes that meet the flags:
```
$ narkwhal -time 3h
account: default, region: us-east-1, Not used within: 3h
~~~~~~~~~~~~~~~~~~~~~~
~~~~~~~~~~~~~~~~~~~~~~
EXITING: There are no available EBS volumes to remove in the us-east-1 region of account default that have been non-active for at least 3h
~~~~~~~~~~~~~~~~~~~~~~
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

#### -regions
The `-regions` flag has multiple possible in puts for example:

- One Region
      - You can place the common format of that region such as `us-west-1` The default if the region flag is not used will be `us-east-1`.

      - EXAMPLE: `narkwhal -regions us-west-2`

- Region Groupings
      - You can select regions based on geographic locations such as America `AM`, Europe `EU`, and Asia Pacific `AP`.
      - EXAMPLE: `narkwhal -regions AM`
            
            - `AM`
                  - us-east-1
		      - us-east-2
		      - us-west-1
		      - us-west-2
                  - ca-central-1
                  - sa-east-1
            
            - `AP`
                  - ap-south-1
		      - ap-northeast-1
		      - ap-northeast-2
		      - ap-southeast-1
		      - ap-southeast-2
            
            - `EU`
                  - eu-central-1
		      - eu-west-1
		      - eu-west-2
		      - eu-west-3
		      - eu-north-1


- Multiple Regions
      - You can place multiple chosen regions separated by a comma.
      
      EXAMPLE: `narkwhal -regions us-west-1,us-east-2, eu-central-1`.

- All Regions
      * To select all available regions just pass the value `ALL`.
      
      * EXAMPLE: `narkwhal -regions ALL`.
      
      * List of regions scoured with the `ALL` value are as follows:

      - us-east-1
	- us-east-2
	- us-west-1
	- us-west-2
	- ap-south-1
	- ap-northeast-1
	- ap-northeast-2
	- ap-southeast-1
	- ap-southeast-2
	- ca-central-1
	- eu-central-1
	- eu-west-1
	- eu-west-2
	- eu-west-3
	- eu-north-1
	- sa-east-1
      
      * NOTE: China, Government, and ap-northeast-3 were left out due to limitations.
      

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