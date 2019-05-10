package cloud

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

// FilterVolumesByTime Uses Cloudwatch to get back any volumes that have not had read ops data report at a specified time.
func NonActiveVolumes(acctData map[string]map[string][]string, accounts []string, timeFlag string) map[string]map[string][]string {

	// If the time value given is 0s we can just leave here with the ID's passed through
	if timeFlag == "0s" {
		return acctData
	}

	// Define the values for the Metric Data Query
	volumeID := "thisIsNeededAndUselessForNow"
	endTime := time.Now()
	duration, _ := time.ParseDuration("-" + timeFlag)
	startTime := endTime.Add(duration)
	nameSpace := "AWS/EBS"
	metricName1 := "VolumeReadOps"
	metricName2 := "VolumeWriteOps"
	period := int64(60)
	stat := "Average"
	metricDimensionName := "VolumeId"

	for acct, regions := range acctData {
		// Loop through each region then child loops through volumes of region to see if they meet the criteria
		for region, sliceOfIDs := range regions {
			svc := createCloudwatchSession(acct, region)

			var sliceVol []string

			for _, volID := range sliceOfIDs {

				query := &cloudwatch.MetricDataQuery{
					Id: &volumeID,
					MetricStat: &cloudwatch.MetricStat{
						Metric: &cloudwatch.Metric{
							Namespace:  &nameSpace,
							MetricName: &metricName1,
							Dimensions: []*cloudwatch.Dimension{
								&cloudwatch.Dimension{
									Name:  &metricDimensionName,
									Value: aws.String(volID),
								},
							},
						},
						Period: &period,
						Stat:   &stat,
					},
				}

				respRead, err := svc.GetMetricData(&cloudwatch.GetMetricDataInput{
					EndTime:           &endTime,
					StartTime:         &startTime,
					MetricDataQueries: []*cloudwatch.MetricDataQuery{query},
				})
				if err != nil {
					fmt.Println("There was an error filtering available volumes in specified time")
					fmt.Println(err.Error())
					os.Exit(1)
				}

				// Replace the metric name with the write metric
				query.MetricStat.Metric.MetricName = &metricName2

				respWrite, err := svc.GetMetricData(&cloudwatch.GetMetricDataInput{
					EndTime:           &endTime,
					StartTime:         &startTime,
					MetricDataQueries: []*cloudwatch.MetricDataQuery{query},
				})
				if err != nil {
					fmt.Println("There was an error filtering available volumes in specified time")
					fmt.Println(err.Error())
					os.Exit(1)
				}

				// If the ID has metric data showing in CloudWatch then skip adding it to the filtered slices
				// If a metric is detected then we do not need to continue with the loop
				// Believe this can be refactored but
				for _, m := range respRead.MetricDataResults {
					if m.Timestamps != nil {
						continue
					}
					sliceVol = append(sliceVol, volID)
					break
				}

				for _, m := range respWrite.MetricDataResults {
					if m.Timestamps != nil {
						continue
					}
					sliceVol = append(sliceVol, volID)
					break
				}

				acctData[acct][region] = sliceVol

				// Do not send back empty values
				uReg := unique(acctData[acct][region])
				if len(uReg) != 0 {
					acctData[acct][region] = uReg
				}

			}
		}
	}
	return acctData
}

func unique(volSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range volSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list
}
