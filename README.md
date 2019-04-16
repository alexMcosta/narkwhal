# Narkwhal

![alt-text](https://i.pinimg.com/originals/74/68/f1/7468f1d665e551fad8eac0c9f97977e3.jpg)

### Summary
Narkwhal goes through your AWS account using the AWS SDK looking for available EBS volumes in default region and deletes them them

### Requirements
- Narkwhal makes use of the AWS credentials folder. More information on this at the following link, [AWS Credentials folder](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#creating-the-credentials-file)

### Feature Ideas

#### Turn application into CLI Application
- Add flags for choosing which account in `.aws/credentials` folder.
[X] Add flag for choosing region (Including all regions)
- Add flag for time stamp date
- Add flag for either choosing specific id's or all EBS volumes

#### Add notifications (Afterall this would not be NARKwhal without it)

