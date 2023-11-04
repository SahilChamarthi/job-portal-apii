package services

import (
	"errors"
	"project/internal/model"

	"github.com/rs/zerolog/log"
)

func (s *Services) JobCreate(nj model.CreateJob, id uint64) (model.Job, error) {
	job := model.Job{JobTitle: nj.JobTitle, JobSalary: nj.JobSalary, Uid: id}
	cu, err := s.r.CreateJob(job)
	if err != nil {
		log.Error().Err(err).Msg("couldnot create job")
		return model.Job{}, errors.New("job creation failed")
	}

	return cu, nil
}

func (s *Services) GetJobsByCompanyId(id int) ([]model.Job, error) {
	AllCompanies, err := s.r.GetJobs(id)
	if err != nil {
		return nil, errors.New("job retreval failed")
	}
	return AllCompanies, nil
}

func (s *Services) FetchAllJobs() ([]model.Job, error) {

	AllJobs, err := s.r.GetAllJobs()
	if err != nil {
		return nil, err
	}
	return AllJobs, nil

}

func (s *Services) Getjobid(id uint64) (model.Job, error) {

	jobData, err := s.r.GetJobId(id)
	if err != nil {
		return model.Job{}, err
	}
	return jobData, nil
}
