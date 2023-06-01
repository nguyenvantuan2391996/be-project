package usecase

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nguyenvantuan2391996/be-project/handler/model"
	modelDomain "github.com/nguyenvantuan2391996/be-project/internal/domain/model"
	"github.com/nguyenvantuan2391996/be-project/internal/domain/repository"
	"github.com/stretchr/testify/assert"
)

func TestScoreRatingDomain_BulkCreateScoreRating(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	scoreRatingRepo := repository.NewMockScoreRatingRepositoryInterface(ctrl)
	scoreRatingRepo.EXPECT().BulkCreateScoreRating(gomock.Any(), gomock.Any()).Return(nil)

	type fields struct {
		scoreRatingRepo repository.IScoreRatingRepositoryInterface
	}
	type args struct {
		ctx            context.Context
		scoreRatingReq []*model.ScoreRatingRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: modelDomain.HappyCase,
			fields: fields{
				scoreRatingRepo: scoreRatingRepo,
			},
			args: args{
				ctx: context.Background(),
				scoreRatingReq: []*model.ScoreRatingRequest{
					{
						UserId:   "123",
						Metadata: "456",
					},
				},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewScoreRatingDomain(tt.fields.scoreRatingRepo)
			tt.wantErr(t, s.BulkCreateScoreRating(tt.args.ctx, tt.args.scoreRatingReq), fmt.Sprintf("BulkCreateScoreRating(%v, %v)", tt.args.ctx, tt.args.scoreRatingReq))
		})
	}
}

func TestScoreRatingDomain_GetListScoreRating(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	scoreRatingRepo := repository.NewMockScoreRatingRepositoryInterface(ctrl)
	scoreRatingRepo.EXPECT().GetScoreRatingByListQueries(gomock.Any(), gomock.Any(), gomock.Any()).Return([]*modelDomain.ScoreRating{}, nil)

	type fields struct {
		scoreRatingRepo repository.IScoreRatingRepositoryInterface
	}
	type args struct {
		ctx         context.Context
		scoreRating *model.ScoreRating
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*modelDomain.ScoreRating
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:   modelDomain.HappyCase,
			fields: fields{scoreRatingRepo: scoreRatingRepo},
			args: args{
				ctx: context.Background(),
				scoreRating: &model.ScoreRating{
					UserId: "test",
				},
			},
			want:    []*modelDomain.ScoreRating{},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ScoreRatingDomain{
				scoreRatingRepo: tt.fields.scoreRatingRepo,
			}
			got, err := s.GetListScoreRating(tt.args.ctx, tt.args.scoreRating)
			if !tt.wantErr(t, err, fmt.Sprintf("GetListScoreRating(%v, %v)", tt.args.ctx, tt.args.scoreRating)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetListScoreRating(%v, %v)", tt.args.ctx, tt.args.scoreRating)
		})
	}
}

func TestScoreRatingDomain_UpdateScoreRating(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	scoreRatingRepo := repository.NewMockScoreRatingRepositoryInterface(ctrl)
	scoreRatingRepo.EXPECT().GetScoreRatingByListQueries(gomock.Any(), gomock.Any(), gomock.Any()).Return([]*modelDomain.ScoreRating{
		{
			UserId: "test",
		},
	}, nil)
	scoreRatingRepo.EXPECT().UpdateScoreRatingWithMap(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

	scoreRatingFailRepo := repository.NewMockScoreRatingRepositoryInterface(ctrl)
	scoreRatingFailRepo.EXPECT().GetScoreRatingByListQueries(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("test"))

	type fields struct {
		scoreRatingRepo repository.IScoreRatingRepositoryInterface
	}
	type args struct {
		ctx         context.Context
		scoreRating *model.ScoreRating
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: modelDomain.HappyCase,
			fields: fields{
				scoreRatingRepo: scoreRatingRepo,
			},
			args: args{
				ctx:         context.Background(),
				scoreRating: &model.ScoreRating{},
			},
			wantErr: assert.NoError,
		},
		{
			name: modelDomain.GetFailCase,
			fields: fields{
				scoreRatingRepo: scoreRatingFailRepo,
			},
			args: args{
				ctx:         context.Background(),
				scoreRating: &model.ScoreRating{},
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ScoreRatingDomain{
				scoreRatingRepo: tt.fields.scoreRatingRepo,
			}
			tt.wantErr(t, s.UpdateScoreRating(tt.args.ctx, tt.args.scoreRating), fmt.Sprintf("UpdateScoreRating(%v, %v)", tt.args.ctx, tt.args.scoreRating))
		})
	}
}

func TestScoreRatingDomain_DeleteScoreRating(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	scoreRatingRepo := repository.NewMockScoreRatingRepositoryInterface(ctrl)
	scoreRatingRepo.EXPECT().DeleteScoreRatingByQueries(gomock.Any(), gomock.Any()).Return(nil)

	type fields struct {
		scoreRatingRepo repository.IScoreRatingRepositoryInterface
	}
	type args struct {
		ctx         context.Context
		scoreRating *model.ScoreRating
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:   modelDomain.HappyCase,
			fields: fields{scoreRatingRepo: scoreRatingRepo},
			args: args{
				ctx:         context.Background(),
				scoreRating: &model.ScoreRating{UserId: "test"},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ScoreRatingDomain{
				scoreRatingRepo: tt.fields.scoreRatingRepo,
			}
			tt.wantErr(t, s.DeleteScoreRating(tt.args.ctx, tt.args.scoreRating), fmt.Sprintf("DeleteScoreRating(%v, %v)", tt.args.ctx, tt.args.scoreRating))
		})
	}
}
