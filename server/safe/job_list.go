package safe

import (
	"sync"
)

type CronJob struct {
	Name, IsRun, NextStartTime string
}

type SafeJobList struct {
	jobs []CronJob
	sync.RWMutex
}

func (j *SafeJobList) GetLen() int {
	j.RLock()
	defer j.RUnlock()
	return len(j.jobs)
}

func (j *SafeJobList) AddJob(c CronJob) {
	j.Lock()
	defer j.Unlock()
	j.jobs = append(j.jobs, c)
}

func (j *SafeJobList) Remove(name string) {
	j.Lock()
	defer j.Unlock()
	for k, c := range j.jobs {
		if c.Name == name {
			j.jobs = append(j.jobs[0:k], j.jobs[k:len(j.jobs)-1]...)
			break
		}
	}
}

func (j *SafeJobList) GetJobList() []CronJob {
	j.RLock()
	defer j.RUnlock()
	return j.jobs
}

func (j *SafeJobList) GetSortedJobList() []CronJob {
	j.SortJobList()
	return j.jobs
}

func (j *SafeJobList) SortJobList() {
	j.Lock()
	defer j.Unlock()

	num := len(j.jobs)
	if num == 0 {
		return
	}

	for i := 0; i < num; i++ {
		for t := 0; t < num-(i+1); t++ {
			if j.jobs[t].NextStartTime >= j.jobs[t+1].NextStartTime {
				tmp := j.jobs[t+1]
				j.jobs[t+1] = j.jobs[t]
				j.jobs[t] = tmp
			}
		}
	}
}

func (j *SafeJobList) UpdateCronByName(name, startTime, isRun string) {
	j.Lock()
	defer j.Unlock()
	for k, c := range j.jobs {
		if c.Name == name {
			j.jobs[k].IsRun = isRun
			j.jobs[k].NextStartTime = startTime
			break
		}
	}
}
