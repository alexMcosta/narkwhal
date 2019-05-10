# Time Flag

The time flag uses Cloudwatch to check if there was any activity from the EBS volumes since the time specified and it does this by checking the `Read Ops` and `Write Ops` metric.

## How Time Flag Works

Time does not check for the last time the volume was attached but rather the last time the volume showed any kind of read activity via the `Read Ops` or `Write Ops` metric. That being said, the volume could have been attached to an EC2 instance that was then stopped for a month. If that EC2 instance was then terminated, without being activated again, the volume would be up for deletion if a time flag of `48h` was passed since it would then be available and also show no signs of activity via `Read Ops` or `Write Ops` for longer then 48 hours.

## Values

The input is based off what can be passed through the `Duration()` function of the Go library `time`. That being said, the duration function takes a string that is of a numeric value followed by characters representing the time measurement value such as `4ms` (4 miliseconds), `25s` (25 seconds), `15m` (15 minutes of fame), `8h` (8 hours), `48d` (48 days later), and `1w` (1 week since you looked at me).
## Example

```
$ narkwhal -time 30h
```

The above would return all available volumes in the default region, in us-east-1, that have had no metrics reporting from Cloudwatch's `Read Ops` or `Write Ops` for the past 30 hours.

[GO BACK TO README](https://github.com/alexMcosta/narkwhal#flags)
