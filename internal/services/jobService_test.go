package services

import (
	"errors"
	"project/internal/model"
	"project/internal/repository"
	"reflect"
	"testing"

	gomock "go.uber.org/mock/gomock"
)

func TestServices_JobCreate(t *testing.T) {
	type args struct {
		nj model.CreateJob
		id uint64
	}
	tests := []struct {
		name             string
		args             args
		want             model.Job
		wantErr          bool
		mockRepoResponse func() (model.Job, error)
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{nj: model.CreateJob{
				JobTitle:  "gwyug",
				JobSalary: "514212",
			},
				id: 15,
			},
			want: model.Job{
				JobTitle:  "gwyug",
				JobSalary: "514212",
				Company: model.Company{
					CompanyName: "widh",
					Adress:      "wkhd",
					Domain:      "wkjhcd",
				},
				Uid: 15,
			},
			wantErr: false,
			mockRepoResponse: func() (model.Job, error) {
				return model.Job{
					JobTitle:  "gwyug",
					JobSalary: "514212",
					Company: model.Company{
						CompanyName: "widh",
						Adress:      "wkhd",
						Domain:      "wkjhcd",
					},
					Uid: 15,
				}, nil

			},
		},

		{
			name: "failure",
			args: args{
				nj: model.CreateJob{},
				id: 15,
			},
			want:    model.Job{},
			wantErr: true,
			mockRepoResponse: func() (model.Job, error) {
				return model.Job{}, errors.New("failed to crerate job")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)

			mockRepo := repository.NewMockAllInRepo(mc)

			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().CreateJob(gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()
			}

			s, _ := NewServices(mockRepo)
			//got, err := tt.s.CompanyCreate(tt.args.nc)
			got, err := s.JobCreate(tt.args.nj, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Services.JobCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Services.JobCreate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServices_GetJobsByCompanyId(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name             string
		args             args
		want             []model.Job
		wantErr          bool
		mockRepoResponse func() ([]model.Job, error)
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				id: 12,
			},
			want: []model.Job{
				{
					JobTitle:  "hcwc",
					JobSalary: "wjhc",
					Uid:       12,
				},
				{
					JobTitle:  "ygyudsgf",
					JobSalary: "buidshfh",
					Uid:       13,
				},
			},
			wantErr: false,
			mockRepoResponse: func() ([]model.Job, error) {
				return []model.Job{
					{
						JobTitle:  "hcwc",
						JobSalary: "wjhc",
						Uid:       12,
					},
					{
						JobTitle:  "ygyudsgf",
						JobSalary: "buidshfh",
						Uid:       13,
					},
				}, nil
			},
		},
		{
			name: "failure",
			args: args{
				id: 12,
			},
			want:    nil,
			wantErr: true,
			mockRepoResponse: func() ([]model.Job, error) {
				return nil, errors.New("job not found")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mc := gomock.NewController(t)

			mockRepo := repository.NewMockAllInRepo(mc)

			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().GetJobs(gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()
			}

			s, _ := NewServices(mockRepo)
			//got, err := tt.s.CompanyCreate(tt.args.nc)
			got, err := s.GetJobsByCompanyId(tt.args.id)

			if (err != nil) != tt.wantErr {
				t.Errorf("Services.GetJobs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Services.GetJobs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServices_FetchAllJobs(t *testing.T) {
	tests := []struct {
		name             string
		want             []model.Job
		wantErr          bool
		mockRepoResponse func() ([]model.Job, error)
	}{
		// TODO: Add test cases.
		{
			name: "sucesss",
			want: []model.Job{
				{
					JobTitle:  "hcwc",
					JobSalary: "wjhc",
					Uid:       12,
				},
				{
					JobTitle:  "ygyudsgf",
					JobSalary: "buidshfh",
					Uid:       13,
				},
			},
			wantErr: false,
			mockRepoResponse: func() ([]model.Job, error) {
				return []model.Job{
					{
						JobTitle:  "hcwc",
						JobSalary: "wjhc",
						Uid:       12,
					},
					{
						JobTitle:  "ygyudsgf",
						JobSalary: "buidshfh",
						Uid:       13,
					},
				}, nil
			},
		},
		{
			name:    "failure",
			want:    nil,
			wantErr: true,
			mockRepoResponse: func() ([]model.Job, error) {
				return nil, errors.New("fail to found jobs")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)

			mockRepo := repository.NewMockAllInRepo(mc)

			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().GetAllJobs().Return(tt.mockRepoResponse()).AnyTimes()
			}

			s, _ := NewServices(mockRepo)
			got, err := s.FetchAllJobs()
			if (err != nil) != tt.wantErr {
				t.Errorf("Services.FetchAllJobs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Services.FetchAllJobs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServices_Getjobid(t *testing.T) {
	type args struct {
		id uint64
	}
	tests := []struct {
		name             string
		args             args
		want             model.Job
		wantErr          bool
		mockRepoResponse func() (model.Job, error)
	}{
		// TODO: Add test cases.

		{
			name: "success",
			args: args{
				id: 15,
			},
			want: model.Job{

				JobTitle:  "eyhe",
				JobSalary: "162522",
				Company: model.Company{
					CompanyName: "bdggd",
					Adress:      "h2usdhd",
					Domain:      "2idh2",
				},
				Uid: 17,
			},
			wantErr: false,
			mockRepoResponse: func() (model.Job, error) {
				return model.Job{

					JobTitle:  "eyhe",
					JobSalary: "162522",
					Company: model.Company{
						CompanyName: "bdggd",
						Adress:      "h2usdhd",
						Domain:      "2idh2",
					},
					Uid: 17,
				}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)

			mockRepo := repository.NewMockAllInRepo(mc)

			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().GetJobId(gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()
			}

			s, _ := NewServices(mockRepo)
			got, err := s.Getjobid(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Services.Getjobid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Services.Getjobid() = %v, want %v", got, tt.want)
			}
		})
	}
}
