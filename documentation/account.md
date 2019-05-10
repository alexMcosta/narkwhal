# Account Flag

The account flag takes the account name between the square brackets from the `.aws/credentials` file and uses it with the AWS SDK.

## Values

The account flag takes whatever value is past through and searches the credentials file for said account. Like `-regions`, `-accounts` can take multiple account inputs seperated by a comma .

If we passed through the following:

```
$ narkwhal -account funTimes
```

Then the following account from the `.aws/credentials` file would be the one accessed

```
[funTimes]
aws_access_key_id = IAMSUCHAVERYCOOLACCESSKEY
aws_secret_access_key = 53cRE74CcEs5M30Wc47
```

## Examples

### Multiple Accounts

**Input**
```
narkwhal -account default,Narkwhal -regions us-west-2 
```

**Output**
```
account: default,Narkwhal, regions: us-west-2, Not used within: 0s
~~~~~~~~~~~
Account: Narkwhal --- Region: us-west-2
~~~~~~~~~~~
vol-0c5eb6b1532332b1d
vol-0d946de72ceffbad7
vol-028ea780f92c40852
---------------------
Would you like to remove the above EBS Volumes? (y/n):
y
~~~~~~~~~~~
Account: default --- Region: us-west-2
~~~~~~~~~~~
vol-0459d076c9be85932
vol-08562da015f48741a
---------------------
Would you like to remove the above EBS Volumes? (y/n):
y
Successfully removed vol-0c5eb6b1532332b1d
Successfully removed vol-0d946de72ceffbad7
Successfully removed vol-028ea780f92c40852
Successfully removed vol-0459d076c9be85932
Successfully removed vol-08562da015f48741a
```

**NOTE:** If you do not have an account in your file associated as `default` then you will get an error if no account is selected.

[GO BACK TO README](https://github.com/alexMcosta/narkwhal#flags)
