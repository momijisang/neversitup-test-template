package persistence

import (
	"neversitup-test-template/internal/pkg/config"
	"neversitup-test-template/internal/pkg/models"
	"sync"
	"time"

	"github.com/go-co-op/gocron"
)

type ScheduleRepository struct{}

var scheduleRepository *ScheduleRepository

func Schedule() *ScheduleRepository {
	if scheduleRepository == nil {
		scheduleRepository = &ScheduleRepository{}
	}
	return scheduleRepository
}

func (r *ScheduleRepository) SetSchedule() {
	conf := config.GetConfig()
	if conf.App.Env == "worker" {
		s := gocron.NewScheduler(time.UTC)
		s.Cron("* * * * *").Do(r.Schedule1, "")
		s.Cron("0 */6 * * *").Do(r.Schedule2)
		s.StartAsync()
	}
}

func (r *ScheduleRepository) Schedule1(param string) {
	listCustomer, err := models.Customer{}.FindAll()
	if err != nil {
		return
	}
	ch := make(chan models.Customer, 3)
	var wg sync.WaitGroup

	for i, video := range listCustomer {
		wg.Add(i)
		go r.DoSomething(ch, &wg, video)

		// Limit the number of concurrent Goroutines to 3
		if len(ch) == 3 {
			// Wait for one Goroutine to finish before starting the next one
			wg.Wait()
			// Remove one value from the channel to allow the next Goroutine to start
			<-ch
		}
	}

	// Wait for the remaining Goroutines to finish
	wg.Wait()

	// Close the channel
	close(ch)
}

func (r *ScheduleRepository) DoSomething(ch chan models.Customer, wg *sync.WaitGroup, customer models.Customer) {
	if wg != nil {
		defer wg.Done()
	}

	//Do something
	//...
	//...

	if ch != nil {
		ch <- customer
	}
}

func (r *ScheduleRepository) Schedule2() {

}
