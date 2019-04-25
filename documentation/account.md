# Account Flag

The account flag takes the account name between the square brackets from the `.aws/credentials` file and uses it with the AWS SDK.

## Values

The account flag takes whatever value is past through and searches the credentials file for said account.

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

**NOTE:** If you do not have an account in your file associated as `default` then you will get an error if no account is selected.

[GO BACK TO README](README.md)
