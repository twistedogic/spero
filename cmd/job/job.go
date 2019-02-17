package job

import "log"

type Job interface {
	Start() error
	Stop()
}
type Runner struct {
	jobs []Job
}

func New(jobs ...Job) *Runner {
	return &Runner{jobs}
}

func (r *Runner) Start() {
	for _, job := range r.jobs {
		go func(j Job) {
			if err := j.Start(); err != nil {
				log.Println(err)
			}
		}(job)
	}
}

func (r *Runner) Stop() {
	for _, job := range r.jobs {
		job.Stop()
	}
}
