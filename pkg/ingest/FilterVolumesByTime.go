package ingest

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

// FilterVolumesByTime Uses Cloudwatch to get back any volumes that have not had read ops data report at a specified time.
func FilterVolumesByTime(regionData map[string][]string, accountFlag string, timeFlag string) map[string][]string {

	// If the time value given is 0s we can just leave here with the ID's passed through
	if timeFlag == "0s" {
		return regionData
	}

	// Define the values for the Metric Data Query
	volumeID := "thisIsNeededAndUselessForNow"
	endTime := time.Now()
	duration, _ := time.ParseDuration("-" + timeFlag)
	startTime := endTime.Add(duration)
	nameSpace := "AWS/EBS"
	metricName := "VolumeReadOps"
	period := int64(60)
	stat := "Average"
	metricDimensionName := "VolumeId"

	for region, sliceOfIDs := range regionData {
		svc := createCloudwatchSession(accountFlag, region)

		var filteredSliceOfVolumes []string

		for _, volID := range sliceOfIDs {

			query := &cloudwatch.MetricDataQuery{
				Id: &volumeID,
				MetricStat: &cloudwatch.MetricStat{
					Metric: &cloudwatch.Metric{
						Namespace:  &nameSpace,
						MetricName: &metricName,
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

			resp, err := svc.GetMetricData(&cloudwatch.GetMetricDataInput{
				EndTime:           &endTime,
				StartTime:         &startTime,
				MetricDataQueries: []*cloudwatch.MetricDataQuery{query},
			})
			if err != nil {
				fmt.Println("There was an error grabbing available volumes in specified time")
				fmt.Println(err.Error())
				os.Exit(1)
			}

			// TODO: Refactor: This feels very wrong
			for _, metricdata := range resp.MetricDataResults {
				if metricdata.Timestamps == nil {
					filteredSliceOfVolumes = append(filteredSliceOfVolumes, volID)
				}
			}

			regionData[region] = filteredSliceOfVolumes

		}
	}
	return regionData
}
