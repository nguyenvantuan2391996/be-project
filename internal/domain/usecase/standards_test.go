package usecase

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nguyenvantuan2391996/be-project/handler/model"
	modelDomain "github.com/nguyenvantuan2391996/be-project/internal/domain/model"
	"github.com/nguyenvantuan2391996/be-project/internal/domain/repository"
	"github.com/stretchr/testify/assert"
)

func TestStandardDomain_CreateStandard(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	standardRepo := repository.NewMockIStandardRepositoryInterface(ctrl)
	standardRepo.EXPECT().CreateStandard(gomock.Any(), gomock.Any()).Return(&modelDomain.Standard{
		StandardName: "test",
	}, nil)

	type fields struct {
		standardRepo repository.IStandardRepositoryInterface
	}
	type args struct {
		ctx             context.Context
		standardRequest *model.StandardRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *modelDomain.Standard
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: modelDomain.HappyCase,
			fields: fields{
				standardRepo: standardRepo,
			},
			args: args{
				ctx: context.Background(),
				standardRequest: &model.StandardRequest{
					StandardName: "test",
				},
			},
			want:    &modelDomain.Standard{StandardName: "test"},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStandardDomain(tt.fields.standardRepo)
			got, err := s.CreateStandard(tt.args.ctx, tt.args.standardRequest)
			if !tt.wantErr(t, err, fmt.Sprintf("CreateStandard(%v, %v)", tt.args.ctx, tt.args.standardRequest)) {
				return
			}
			assert.Equalf(t, tt.want, got, "CreateStandard(%v, %v)", tt.args.ctx, tt.args.standardRequest)
		})
	}
}

func TestStandardDomain_GetStandards(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	standardRepo := repository.NewMockIStandardRepositoryInterface(ctrl)
	standardRepo.EXPECT().GetStandardByQueries(gomock.Any(), gomock.Any()).Return([]*modelDomain.Standard{
		{
			UserId: "123",
		},
	}, nil)

	type fields struct {
		standardRepo repository.IStandardRepositoryInterface
	}
	type args struct {
		ctx      context.Context
		standard *model.Standard
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*modelDomain.Standard
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:   modelDomain.HappyCase,
			fields: fields{standardRepo: standardRepo},
			args: args{
				ctx:      context.Background(),
				standard: &model.Standard{},
			},
			want: []*modelDomain.Standard{
				{
					UserId: "123",
				},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StandardDomain{
				standardRepo: tt.fields.standardRepo,
			}
			got, err := s.GetStandards(tt.args.ctx, tt.args.standard)
			if !tt.wantErr(t, err, fmt.Sprintf("GetStandards(%v, %v)", tt.args.ctx, tt.args.standard)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetStandards(%v, %v)", tt.args.ctx, tt.args.standard)
		})
	}
}

func TestStandardDomain_DeleteStandard(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	standardRepo := repository.NewMockIStandardRepositoryInterface(ctrl)
	standardRepo.EXPECT().DeleteStandardByQueries(gomock.Any(), gomock.Any()).Return(nil)

	type fields struct {
		standardRepo repository.IStandardRepositoryInterface
	}
	type args struct {
		ctx      context.Context
		standard *model.Standard
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:   modelDomain.HappyCase,
			fields: fields{standardRepo: standardRepo},
			args: args{
				ctx:      context.Background(),
				standard: &model.Standard{},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StandardDomain{
				standardRepo: tt.fields.standardRepo,
			}
			tt.wantErr(t, s.DeleteStandard(tt.args.ctx, tt.args.standard), fmt.Sprintf("DeleteStandard(%v, %v)", tt.args.ctx, tt.args.standard))
		})
	}
}
