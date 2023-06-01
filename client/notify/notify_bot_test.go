package notify

import (
	"context"
	"testing"

	"github.com/nguyenvantuan2391996/be-project/config"
	"github.com/nguyenvantuan2391996/be-project/internal/domain/usecase"
)

func TestBotNotify_ProcessNotifyDailyLeetCodingChallenge(t *testing.T) {
	t.Skip()
	type fields struct {
		cfg               *config.Config
		statisticalDomain *usecase.StatisticalDomain
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "happy_case",
			fields: fields{
				cfg: &config.Config{
					WebhookSlackLeetCode: "",
					TagsSlackLeetCode:    "<@tuan.nguyen25>",
				},
				statisticalDomain: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BotNotify{
				cfg:               tt.fields.cfg,
				statisticalDomain: tt.fields.statisticalDomain,
			}
			if err := b.ProcessNotifyDailyLeetCodingChallenge(context.Background()); (err != nil) != tt.wantErr {
				t.Errorf("ProcessNotifyDailyLeetCodingChallenge() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
