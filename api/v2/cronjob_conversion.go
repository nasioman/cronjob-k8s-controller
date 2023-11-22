package v2

import (
	v1 "example/user/hello/api/v1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"strings"
)

func (src *CronJob) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1.CronJob)

	sched := src.Spec.Schedule
	scheduleParts := []string{"*", "*", "*", "*", "*"}
	if sched.Minute != nil {
		scheduleParts[0] = string(*sched.Minute)
	}
	if sched.Hour != nil {
		scheduleParts[1] = string(*sched.Hour)
	}
	if sched.DayOfMonth != nil {
		scheduleParts[2] = string(*sched.DayOfMonth)
	}
	if sched.Month != nil {
		scheduleParts[3] = string(*sched.Month)
	}
	if sched.DayOfWeek != nil {
		scheduleParts[4] = string(*sched.DayOfWeek)
	}

	dst.Spec.Schedule = strings.Join(scheduleParts, " ")

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.StartingDeadlineSeconds = src.Spec.StartingDeadlineSeconds
	dst.Spec.ConcurrencyPolicy = v1.ConcurrencyPolicy(src.Spec.ConcurrencyPolicy)
	dst.Spec.Suspend = src.Spec.Suspend
	dst.Spec.JobTemplate = src.Spec.JobTemplate
	dst.Spec.SuccessfulJobsHistoryLimit = src.Spec.SuccessfulJobsHistoryLimit
	dst.Spec.FailedJobsHistoryLimit = src.Spec.FailedJobsHistoryLimit

	// Status
	dst.Status.Active = src.Status.Active
	dst.Status.LastScheduleTime = src.Status.LastScheduleTime

	return nil
}

// ConvertFrom converts from the Hub version (v1) to this version.
func (dst *CronJob) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1.CronJob)

	sched := src.Spec.Schedule
	schedParts := strings.Split(src.Spec.Schedule, " ")
	f := func(raw string) *CronField {
		if sched == "*" {
			return nil
		}
		part := CronField(raw)
		return &part
	}
	dst.Spec.Schedule.Minute = f(schedParts[0])
	dst.Spec.Schedule.Hour = f(schedParts[1])
	dst.Spec.Schedule.DayOfMonth = f(schedParts[2])
	dst.Spec.Schedule.Month = f(schedParts[3])
	dst.Spec.Schedule.DayOfWeek = f(schedParts[4])

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.StartingDeadlineSeconds = src.Spec.StartingDeadlineSeconds
	dst.Spec.ConcurrencyPolicy = ConcurrencyPolicy(src.Spec.ConcurrencyPolicy)
	dst.Spec.Suspend = src.Spec.Suspend
	dst.Spec.JobTemplate = src.Spec.JobTemplate
	dst.Spec.SuccessfulJobsHistoryLimit = src.Spec.SuccessfulJobsHistoryLimit
	dst.Spec.FailedJobsHistoryLimit = src.Spec.FailedJobsHistoryLimit

	// Status
	dst.Status.Active = src.Status.Active
	dst.Status.LastScheduleTime = src.Status.LastScheduleTime
	return nil

}
