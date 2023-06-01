package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/nguyenvantuan2391996/be-project/client/constants"
	"github.com/nguyenvantuan2391996/be-project/client/notify"
	"github.com/nguyenvantuan2391996/be-project/config"
	"github.com/nguyenvantuan2391996/be-project/internal/domain/usecase"
	"github.com/nguyenvantuan2391996/be-project/internal/infrastructure/repository"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	cfg, err := config.LoadConfig("build")
	if err != nil {
		logrus.Warnf("cannot load config: %v", err)
		return
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(cfg.DBSource), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal("failed to open database:", err)
		return
	}

	ctx := NewContext()

	statisticalRepo := repository.NewStatisticalRepository(db)
	statisticalDomain := usecase.NewStatisticalDomain(statisticalRepo)

	notifyBot := notify.NewBotNotify(cfg, statisticalDomain)
	logrus.Info("Create new cron")
	c := cron.New()

	// Notify run at 17:30 every day
	notifyRun, err := c.AddFunc(cfg.CronNotifyRun, func() {
		err = notifyBot.ProcessNotifyRun(ctx)
		if err != nil {
			logrus.Warnf(constants.TaskErrorMessage, "ProcessNotifyRun", err)
		}
	})

	// Notify summary 09:00 every day
	notifySummary, err := c.AddFunc(cfg.CronNotifySummary, func() {
		err = notifyBot.ProcessNotifySummary(ctx)
		if err != nil {
			logrus.Warnf(constants.TaskErrorMessage, "ProcessNotifySummary", err)
		}

		err = notifyBot.ProcessNotifyDailyLeetCodingChallenge(ctx)
		if err != nil {
			logrus.Warnf(constants.TaskErrorMessage, "ProcessNotifyDailyLeetCodingChallenge", err)
		}
	})

	// Notify statistical on day-of-month 1
	notifyStatistical, err := c.AddFunc(cfg.CronNotifyStatistical, func() {
		err = notifyBot.ProcessNotifyStatistical(ctx)
		if err != nil {
			logrus.Warnf(constants.TaskErrorMessage, "ProcessNotifyStatistical", err)
		}
	})
	if err != nil {
		c.Remove(notifyRun)
		c.Remove(notifySummary)
		c.Remove(notifyStatistical)
		return
	}

	// Start cron with one scheduled job
	logrus.Info("Start cron")
	c.Run()
}

func NewContext() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, constants.XRequestID, uuid.New().String())

	return ctx
}
