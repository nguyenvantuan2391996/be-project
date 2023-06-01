package usecase

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nguyenvantuan2391996/be-project/internal/domain/model"
	"github.com/nguyenvantuan2391996/be-project/internal/domain/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestStatisticalDomain_UpsertStatistical_RecordNotFound_OK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	statisticalRepo := repository.NewMockIStatisticalRepositoryInterface(ctrl)
	statisticalRepo.EXPECT().GetByQueries(gomock.Any()).Return(nil, gorm.ErrRecordNotFound)
	statisticalRepo.EXPECT().Create(gomock.Any()).Return(nil)

	statisticalDomain := NewStatisticalDomain(statisticalRepo)
	err := statisticalDomain.UpsertStatistical([]*model.ChartInfo{
		{
			Day:   "19-08-2022",
			Value: 2.3,
		},
	})

	assert.Nil(t, err)
}

func TestStatisticalDomain_UpsertStatistical_OK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	statisticalRepo := repository.NewMockIStatisticalRepositoryInterface(ctrl)
	statisticalRepo.EXPECT().GetByQueries(gomock.Any()).Return(&model.Statistical{Metadata: "[{\"day\":\"19-08-2022\",\"value\":2.3}]"}, nil)
	statisticalRepo.EXPECT().UpdateWithMap(gomock.Any(), gomock.Any()).Return(nil)

	statisticalDomain := NewStatisticalDomain(statisticalRepo)
	err := statisticalDomain.UpsertStatistical([]*model.ChartInfo{
		{
			Day:   "19-08-2022",
			Value: 2.3,
		},
	})

	assert.Nil(t, err)
}

func TestStatisticalDomain_UpsertStatistical_Fail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	statisticalRepo := repository.NewMockIStatisticalRepositoryInterface(ctrl)
	statisticalRepo.EXPECT().GetByQueries(gomock.Any()).Return(nil, errors.New("test"))

	statisticalDomain := NewStatisticalDomain(statisticalRepo)
	err := statisticalDomain.UpsertStatistical([]*model.ChartInfo{
		{
			Day:   "19-08-2022",
			Value: 2.3,
		},
	})

	assert.NotNil(t, err)
}

func TestStatisticalDomain_UpsertStatistical_Fail_Unmarshal(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	statisticalRepo := repository.NewMockIStatisticalRepositoryInterface(ctrl)
	statisticalRepo.EXPECT().GetByQueries(gomock.Any()).Return(&model.Statistical{Metadata: "19-08-2022"}, nil)

	statisticalDomain := NewStatisticalDomain(statisticalRepo)
	err := statisticalDomain.UpsertStatistical([]*model.ChartInfo{
		{
			Day:   "19-08-2022",
			Value: 2.3,
		},
	})

	assert.NotNil(t, err)
}

func TestStatisticalDomain_GetBase64StringChart_OK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	statisticalRepo := repository.NewMockIStatisticalRepositoryInterface(ctrl)
	statisticalRepo.EXPECT().GetByQueries(gomock.Any()).Return(&model.Statistical{
		Metadata: "[{\"day\": \"19-08-2022\", \"value\": 2.3}, {\"day\": \"20-08-2022\", \"value\": 3.06}, {\"day\": \"21-08-2022\", \"value\": 3.22}, {\"day\": \"22-08-2022\", \"value\": 2.19}, {\"day\": \"23-08-2022\", \"value\": 2.84}, {\"day\": \"26-08-2022\", \"value\": 2.14}, {\"day\": \"27-08-2022\", \"value\": 2.27}, {\"day\": \"28-08-2022\", \"value\": 3.01}, {\"day\": \"29-08-2022\", \"value\": 2.29}, {\"day\": \"30-08-2022\", \"value\": 0}]",
	}, nil)

	statisticalDomain := StatisticalDomain{
		statisticalRepo: statisticalRepo,
	}

	_, sumKilometers, err := statisticalDomain.GetBase64StringChart(map[string]interface{}{})

	assert.Nil(t, err)
	assert.Equal(t, 23.32, sumKilometers)
}

func TestStatisticalDomain_GetBase64StringChart_Fail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	statisticalRepo := repository.NewMockIStatisticalRepositoryInterface(ctrl)
	statisticalRepo.EXPECT().GetByQueries(gomock.Any()).Return(nil, errors.New("test"))

	statisticalDomain := StatisticalDomain{
		statisticalRepo: statisticalRepo,
	}

	_, _, err := statisticalDomain.GetBase64StringChart(map[string]interface{}{})

	assert.NotNil(t, err)
}

func TestStatisticalDomain_GetBase64StringChart_Fail_Render(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	statisticalRepo := repository.NewMockIStatisticalRepositoryInterface(ctrl)
	statisticalRepo.EXPECT().GetByQueries(gomock.Any()).Return(&model.Statistical{
		Metadata: "[{\"day\":\"19-08-2022\",\"value\":2.3}]",
	}, nil)

	statisticalDomain := StatisticalDomain{
		statisticalRepo: statisticalRepo,
	}

	_, sumKilometers, err := statisticalDomain.GetBase64StringChart(map[string]interface{}{})

	assert.NotNil(t, err)
	assert.Equal(t, float64(0), sumKilometers)
}

func TestStatisticalDomain_GetBase64StringChart_Fail_Unmarshal(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	statisticalRepo := repository.NewMockIStatisticalRepositoryInterface(ctrl)
	statisticalRepo.EXPECT().GetByQueries(gomock.Any()).Return(&model.Statistical{
		Metadata: "19-08-2022",
	}, nil)

	statisticalDomain := StatisticalDomain{
		statisticalRepo: statisticalRepo,
	}

	_, sumKilometers, err := statisticalDomain.GetBase64StringChart(map[string]interface{}{})

	assert.NotNil(t, err)
	assert.Equal(t, float64(0), sumKilometers)
}
