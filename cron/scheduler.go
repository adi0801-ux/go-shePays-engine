package cron

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"shepays/repositories"
	"shepays/utils"
	"time"
)

type Cron struct {
	Sc           *gocron.Scheduler
	Conf         *utils.Config
	HappayClient *repositories.NSDLClient
}

func CreateScheduler() *gocron.Scheduler {
	return gocron.NewScheduler(time.Local)
}

func timer() {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for {
			select {

			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()
}

func (cron *Cron) InitializeScheduler() {
	utils.Log.Info("initializing crons")
	cron.AuthApiJobs()

	//list more jobs
	cron.Sc.StartAsync()
}

func (cron *Cron) AuthApiJobs() {

	job := cron.Sc.Every(5).Minutes()

	_, err := job.Do(cron.CallAuthApi)
	if err != nil {
		return
	}

}

func (cron *Cron) CallAuthApi() {
	//	 fetch all active SIP's where date is current date
	//	deduct wallet balance --> call service

}
