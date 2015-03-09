package schedule

import (
	"testing"
	"time"

	"github.com/intelsdilabs/pulse/core"
	. "github.com/smartystreets/goconvey/convey"
)

func TestWorkerManager(t *testing.T) {
	Convey("WorkerManager", t, func() {
		manager := new(managesWork)
		Convey("Work", func() {
			metricTypes := []core.MetricType{
				&MockMetricType{
					namespace:          []string{"foo", "bar"},
					version:            1,
					lastAdvertisedTime: time.Now(),
				},
			}
			job := NewCollectorJob(metricTypes)
			job = manager.Work(job)
			So(job.Errors(), ShouldBeNil)
			So(job.(CollectorJob).Metrics(), ShouldBeNil)
		})
	})
}
