package ingest

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

// FilterOldVolumesByTime Uses Cloudwatch to get back any volumes that have not had read ops data report at a specified time.
func FilterOldVolumesByTime(accountFlag string, regionFlag string, timeFlag string) []string {

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
	var sliceOfVolumes []string

	availableVolumes := GrabAvailableVolumes(accountFlag, regionFlag)

	for _, value := range availableVolumes.Volumes {

		query := &cloudwatch.MetricDataQuery{
			Id: &volumeID,
			MetricStat: &cloudwatch.MetricStat{
				Metric: &cloudwatch.Metric{
					Namespace:  &nameSpace,
					MetricName: &metricName,
					Dimensions: []*cloudwatch.Dimension{
						&cloudwatch.Dimension{
							Name:  &metricDimensionName,
							Value: aws.String(*value.VolumeId),
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
				sliceOfVolumes = append(sliceOfVolumes, *value.VolumeId)
			}
		}

	}

	//TODO: Refactor: Repeat from another function
	if sliceOfVolumes == nil {
		fmt.Println("~~~~~~~~~~~~~~~~~~()~~~~")
		fmt.Printf("EXITING: There are no available EBS volumes in the %s region to remove\n", regionFlag)
		fmt.Println("~~~~~~~~~~~~~~~~~~()~~~~")
		fmt.Println("---------------------")
		os.Exit(1)
	}
	return sliceOfVolumes
}
