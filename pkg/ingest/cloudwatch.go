package ingest

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

func FilterOldVolumesByTime(accountFlag string, regionFlag string, timeFlag string) []string {
	svc := createCloudwatchSession(accountFlag, regionFlag)

	volumeID1 := "thisIsNeeded"
	endTime := time.Now()
	duration, _ := time.ParseDuration("-" + timeFlag)
	startTime := endTime.Add(duration)
	nameSpace := "AWS/EBS"
	metricName := "VolumeReadOps"
	period := int64(60)
	stat := "Average"
	metricDim1Name := "VolumeId"
	var sliceOfVolumes []string

	availableVolumes := GrabAvailableVolumes(accountFlag, regionFlag)

	for _, value := range availableVolumes.Volumes {

		query := &cloudwatch.MetricDataQuery{
			Id: &volumeID1,
			MetricStat: &cloudwatch.MetricStat{
				Metric: &cloudwatch.Metric{
					Namespace:  &nameSpace,
					MetricName: &metricName,
					Dimensions: []*cloudwatch.Dimension{
						&cloudwatch.Dimension{
							Name:  &metricDim1Name,
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

		for _, metricdata := range resp.MetricDataResults {
			if metricdata.Timestamps == nil {
				sliceOfVolumes = append(sliceOfVolumes, *value.VolumeId)
			}
		}

	}
	if sliceOfVolumes == nil {
		fmt.Println("~~~~~~~~~~~~~~~~~~()~~~~")
		fmt.Printf("EXITING: There are no available EBS volumes in the %s region to remove\n", regionFlag)
		fmt.Println("~~~~~~~~~~~~~~~~~~()~~~~")
		fmt.Println("---------------------")
		os.Exit(1)
	}
	return sliceOfVolumes
}
