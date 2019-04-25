# Regions Flag

The `-regions` flag lets you choose which region or regions you would like to select. 

## Values

### One Region

The `-regions` flag can take one region in the tradition AWS region format of `us-west-1`. For example, If I entered the below:

```
$ narkwhal -regions us-west-1
```

Narkwhal would search in `us-west-1` for any available volumes in the default region.

Example of an empty return:

```
$ narkwhal -regions us-west-1
account: default, regions: us-west-1, Not used within: 0s
~~~~~~~~~~~~~~~~~~~~~~
No Available Volumes in us-west-1
~~~~~~~~~~~~~~~~~~~~~~
```

Example of a value return:

```
account: default, regions: us-west-1, Not used within: 0s
~~~~~~~~~~~
us-west-1
~~~~~~~~~~~
vol-0d2b7ec6350a1b67d
---------------------
```

### Region Grouping

The `-regions` flag can take special region grouping inputs that will process all regions in that group. The four groups and their regions are as follows:

| AM           | AP              | EU            | ALL  |
|--------------|-----------------|---------------|------|
| us-east-1    | ap-south-1      | eu-central-1  | Name |
| us-east-2    | ap-northeast-1  | eu-west-1     | Says |
| us-west-1    | ap-northeast-2  | eu-west-1     | It   |
| us-west-2    | ap-southeast-1  | eu-west-1     | All  |
| ca-central-1 | ap-southeast-2  | eu-north-1    |      |
| sa-east-1    |                 |               |      |

For example, if the following is entered:

```
$ narkwhal -regions AM
```

The following would be outputed:

```     
account: default, regions: AM, Not used within: 0s
~~~~~~~~~~~~~~~~~~~~~~
No Available Volumes in ca-central-1
~~~~~~~~~~~~~~~~~~~~~~
~~~~~~~~~~~~~~~~~~~~~~
No Available Volumes in sa-east-1
~~~~~~~~~~~~~~~~~~~~~~
~~~~~~~~~~~~~~~~~~~~~~
No Available Volumes in us-east-1
~~~~~~~~~~~~~~~~~~~~~~
~~~~~~~~~~~~~~~~~~~~~~
No Available Volumes in us-east-2
~~~~~~~~~~~~~~~~~~~~~~
~~~~~~~~~~~
us-west-1
~~~~~~~~~~~
vol-0d2b8eb6300a1b67d
---------------------
Would you like to remove the above EBS Volumes? (y/n):
```

### Multiple Inputs

The `-regions` flag can take multiple region and grouping inputs that will process all regions in the list. Comma separated with no whitespace. For instance if I input the following:

```
$ narkwhal -regions AP,us-west-1 
```

I would get the following output

```
account: default, regions: AP,us-west-1, Not used within: 0s
~~~~~~~~~~~~~~~~~~~~~~
No Available Volumes in ap-northeast-2
~~~~~~~~~~~~~~~~~~~~~~
~~~~~~~~~~~~~~~~~~~~~~
No Available Volumes in ap-southeast-1
~~~~~~~~~~~~~~~~~~~~~~
~~~~~~~~~~~~~~~~~~~~~~
No Available Volumes in ap-southeast-2
~~~~~~~~~~~~~~~~~~~~~~
~~~~~~~~~~~
us-west-1
~~~~~~~~~~~
vol-0d2b8eb6300a1b67d
---------------------
Would you like to remove the above EBS Volumes? (y/n): 
```
