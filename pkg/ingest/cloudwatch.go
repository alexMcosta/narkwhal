package ingest

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

// FilterVolumesByTime Uses Cloudwatch to get back any volumes that have not had read ops data report at a specified time.
func FilterVolumesByTime(sliceOfVolumes []string, accountFlag string, regionFlag string, timeFlag string) []string {

	// If the time value given is 0s we can just leave here with the ID's passed through
	if timeFlag == "0s" {
		return sliceOfVolumes
	}

	svc := createCloudwatchSession(accountFlag, regionFlag)

	volumeID := "thisIsNeeded"
	endTime := time.Now()
	duration, _ := time.ParseDuration("-" + timeFlag)
	startTime := endTime.Add(duration)
	nameSpace := "AWS/EBS"
	metricName := "VolumeReadOps"
	period := int64(60)
	stat := "Average"
	metricDimensionName := "VolumeId"
	var filteredSliceOfVolumes []string

	for _, value := range sliceOfVolumes {

		query := &cloudwatch.MetricDataQuery{
			Id: &volumeID,
			MetricStat: &cloudwatch.MetricStat{
				Metric: &cloudwatch.Metric{
					Namespace:  &nameSpace,
					MetricName: &metricName,
					Dimensions: []*cloudwatch.Dimension{
						&cloudwatch.Dimension{
							Name:  &metricDimensionName,
							Value: aws.String(value),
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
			fmt.Println("There Was an error grabbing available volumes in specified time")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		// TODO: Refactor: This feels very wrong
		for _, metricdata := range resp.MetricDataResults {
			if metricdata.Timestamps == nil {
				filteredSliceOfVolumes = append(filteredSliceOfVolumes, value)
			}
		}

	}
	return filteredSliceOfVolumes
}
