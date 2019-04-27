The time flag uses Cloudwatch to check if there was any activity from the EBS volumes since the time specified and it does this by checking the `Read Ops` metric. 
Time does not check for the last time the volume was attached but rather the last time the volume showed any kind of read activity. That being said, the volume 
could have been attached to an EC2 instance that was then stopped for a month. If that EC2 instance was then terminated, without being activated again, the volume 
would be up for deletion if a time flag of `48h` was passed since it would then be available and also show no signs of activity for longer then 48 hours.

The input is based off of the Go function `time.Duration()`.

# Time Flag

The time flag uses Cloudwatch to check if there was any activity from the EBS volumes since the time specified and it does this by checking the `Read Ops` metric.

## How Time Flag Works

Time does not check for the last time the volume was attached but rather the last time the volume showed any kind of read activity via the `Read Ops` metric. That being said, the volume could have been attached to an EC2 instance that was then stopped for a month. If that EC2 instance was then terminated, without being activated again, the volume would be up for deletion if a time flag of `48h` was passed since it would then be available and also show no signs of activity via `Read Ops` for longer then 48 hours. At this time there is a feature in the road-map to include checking Cloudwatch's `Write Ops` metric to increase the likeliness that the volume is inactive.

## Values

The input is based off what can be passed through the `Duration()` function of the Go library `time`. That being said, the duration function takes a string that is of a numeric value followed by characters representing the time measurement value such as `4ms` (4 miliseconds), `25s` (25 seconds), `10m` (10 minutes), and `48h` (48 hours). At this time the highest time measurement accepted is hours and I have it in the feature-road-map to include functionality to allow the time measurements of days and weeks.

## Example

```
$ narkwhal -time 30h
```

The above would return all available volumes in the default region, in us-east-1, that have had no metrics reporting from Cloudwatch's `Read Ops`.

[GO BACK TO README](https://github.com/alexMcosta/narkwhal#flags)
