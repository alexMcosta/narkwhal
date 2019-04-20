package ingest

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

func FilterOldVolumes(accountFlag string, regionFlag string) {
	svc := createCloudwatchSession(accountFlag, regionFlag)

	volumeID1 := "vol000577028616e3640"
	endTime := time.Now()
	duration, _ := time.ParseDuration("-6h")
	startTime := endTime.Add(duration)
	nameSpace := "AWS/EBS"
	metricName := "VolumeWriteBytes"
	period := int64(60)
	stat := "Average"
	metricDim1Name := "VolumeId"
	metricDim1Value := "vol-0c72f0ff014bcefaf"

	query := &cloudwatch.MetricDataQuery{
		Id: &volumeID1,
		MetricStat: &cloudwatch.MetricStat{
			Metric: &cloudwatch.Metric{
				Namespace:  &nameSpace,
				MetricName: &metricName,
				Dimensions: []*cloudwatch.Dimension{
					&cloudwatch.Dimension{
						Name:  &metricDim1Name,
						Value: &metricDim1Value,
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

	exists := "false"

	for _, metricdata := range resp.MetricDataResults {
		if metricdata.Timestamps != nil {
			exists = "true"
		}
	}
	fmt.Println(exists)
}
