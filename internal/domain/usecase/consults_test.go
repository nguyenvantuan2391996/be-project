package usecase

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nguyenvantuan2391996/be-project/internal/domain/model"
	"github.com/nguyenvantuan2391996/be-project/internal/domain/repository"
	"github.com/stretchr/testify/assert"
)

func TestConsultDomain_Consult(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	standardRepo := repository.NewMockIStandardRepositoryInterface(ctrl)
	standardRepo.EXPECT().GetStandardByQueries(gomock.Any(), gomock.Any()).Return([]*model.Standard{
		{
			StandardName: "Chiều cao",
			Weight:       6,
			Type:         MaxMax,
		},
		{
			StandardName: "Tuổi",
			Weight:       8,
			Type:         MinMax,
		},
	}, nil)

	scoreRatingRepo := repository.NewMockScoreRatingRepositoryInterface(ctrl)
	scoreRatingRepo.EXPECT().GetScoreRatingByListQueries(gomock.Any(), gomock.Any(), gomock.Any()).Return([]*model.ScoreRating{
		{
			Metadata: "[{\"name\":\"Nguyen Van A\",\"score\":5,\"standard_name\":\"Chiều cao\"},{\"name\":\"Nguyen Van A\",\"score\":6,\"standard_name\":\"Tuổi\"}]",
		},
		{
			Metadata: "[{\"name\": \"Nguyen Van B\", \"score\": 6, \"standard_name\": \"Chiều cao\"}, {\"name\": \"Nguyen Van B\", \"score\": 8, \"standard_name\": \"Tuổi\"}]",
		},
	}, nil)

	type fields struct {
		userRepo        repository.IUserRepositoryInterface
		standardRepo    repository.IStandardRepositoryInterface
		scoreRatingRepo repository.IScoreRatingRepositoryInterface
	}
	type args struct {
		ctx    context.Context
		userId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.ConsultResult
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: model.HappyCase,
			fields: fields{
				userRepo:        nil,
				standardRepo:    standardRepo,
				scoreRatingRepo: scoreRatingRepo,
			},
			args: args{
				ctx:    context.Background(),
				userId: "test-123",
			},
			want: []*model.ConsultResult{
				{
					Name:       "Nguyen Van A",
					Similarity: 0.675612542537418,
				},
				{
					Name:       "Nguyen Van B",
					Similarity: 0.324387457462582,
				},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewConsultDomain(tt.fields.userRepo, tt.fields.standardRepo, tt.fields.scoreRatingRepo)
			got, err := c.Consult(tt.args.ctx, tt.args.userId)
			if !tt.wantErr(t, err, fmt.Sprintf("Consult(%v, %v)", tt.args.ctx, tt.args.userId)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Consult(%v, %v)", tt.args.ctx, tt.args.userId)
		})
	}
}
