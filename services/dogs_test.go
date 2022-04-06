package services

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/JustSomeHack/go-api-sample/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Test_dogsService_Add(t *testing.T) {
	teardownTests := setupTests(t)
	defer teardownTests(t)

	dogID := uuid.New()

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx context.Context
		dog *models.Dog
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *uuid.UUID
		wantErr bool
	}{
		{
			name:   "Should add a dog to the database",
			fields: fields{db: db},
			args: args{
				ctx: context.Background(),
				dog: &models.Dog{
					ID:        dogID,
					Name:      "Snowball",
					Breed:     "Sniba Inu",
					Color:     "Cream",
					Birthdate: time.Now(),
					Weight:    22,
				},
			},
			want:    &dogID,
			wantErr: false,
		},
		{
			name:   "Should not be able to add empty dog",
			fields: fields{db: db},
			args: args{
				ctx: context.Background(),
				dog: &models.Dog{
					ID: uuid.New(),
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &dogsService{
				db: tt.fields.db,
			}
			got, err := s.Add(tt.args.ctx, tt.args.dog)
			if (err != nil) != tt.wantErr {
				t.Errorf("dogsService.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dogsService.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dogsService_Delete(t *testing.T) {
	teardownTests := setupTests(t)
	defer teardownTests(t)

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Should delete dog by ID",
			fields:  fields{db: db},
			args:    args{ctx: context.Background(), id: dogs[0].ID},
			wantErr: false,
		},
		{
			name:    "Should not delete an ID that does not exists",
			fields:  fields{db: db},
			args:    args{ctx: context.Background(), id: uuid.New()},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &dogsService{
				db: tt.fields.db,
			}
			if err := s.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("dogsService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_dogsService_Get(t *testing.T) {
	teardownTests := setupTests(t)
	defer teardownTests(t)

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx    context.Context
		filter interface{}
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantCount int
		wantErr   bool
	}{
		{
			name:      "Should get all the dogs",
			fields:    fields{db: db},
			args:      args{ctx: context.Background(), filter: nil},
			wantCount: 2,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &dogsService{
				db: tt.fields.db,
			}
			got, err := s.Get(tt.args.ctx, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("dogsService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.wantCount {
				t.Errorf("dogsService.Get() = %v, want %v", len(got), tt.wantCount)
			}
		})
	}
}

func Test_dogsService_GetOne(t *testing.T) {
	teardownTests := setupTests(t)
	defer teardownTests(t)

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Dog
		wantErr bool
	}{
		{
			name:    "Should get dog by ID",
			fields:  fields{db: db},
			args:    args{ctx: context.Background(), id: dogs[0].ID},
			want:    &dogs[0],
			wantErr: false,
		},
		{
			name:    "Should not get dog that does not exist",
			fields:  fields{db: db},
			args:    args{ctx: context.Background(), id: uuid.New()},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &dogsService{
				db: tt.fields.db,
			}
			got, err := s.GetOne(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("dogsService.GetOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dogsService.GetOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dogsService_Update(t *testing.T) {
	teardownTests := setupTests(t)
	defer teardownTests(t)

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx context.Context
		id  uuid.UUID
		dog *models.Dog
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "Should update a dog by ID",
			fields: fields{db: db},
			args: args{ctx: context.Background(), id: dogs[0].ID, dog: &models.Dog{
				Name:      "Snowball",
				Breed:     "Shiba Inu",
				Color:     "Cream",
				Birthdate: time.Date(2020, 6, 12, 0, 0, 0, 0, time.UTC),
				Weight:    22,
			}},
			wantErr: false,
		},
		{
			name:   "Should not update a dog with no valid ID",
			fields: fields{db: db},
			args: args{ctx: context.Background(), id: uuid.New(), dog: &models.Dog{
				Name:      "Snowball",
				Breed:     "Shiba Inu",
				Color:     "Cream",
				Birthdate: time.Date(2020, 6, 12, 0, 0, 0, 0, time.UTC),
				Weight:    22,
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &dogsService{
				db: tt.fields.db,
			}
			if err := s.Update(tt.args.ctx, tt.args.id, tt.args.dog); (err != nil) != tt.wantErr {
				t.Errorf("dogsService.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}