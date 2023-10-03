package acc

import (
	"fmt"
	"time"
)

type AccData struct {
	AccVersion               string
	SessionType              AccSessionType
	Track                    string
	CarModel                 string
	SessionLength            time.Duration
	SessionTime              time.Duration
	LapTime                  time.Duration
	ProgressWithFuel         float32
	SessionLaps              int
	SessionLiter             int
	RaceProgressPercentage   float32
	FuelLevel                float32
	FuelPerLap               float32
	CompletedLaps            int
	LapsWithFuel             float32
	LapsDone                 int
	BoxLap                   int
	LapsToGo                 float32
	RefuelLevel              float32
	PitWindowStartTime       time.Duration
	PitWindowCloseTime       time.Duration
	PitWindowStartPercentage float32
	PitWindowEndPercentage   float32
}

type AccUpdater interface {
	update() (AccData, error)
}

// StartUpdater start ticker to read acc shm data in updateIntervalSeconds
// and send data into accChan
func StartUpdater(updateIntervalSeconds int, accChan chan<- AccData) {
	fmt.Printf("starting acc update ticker ...\n")
	ticker := time.NewTicker(time.Second * time.Duration(updateIntervalSeconds))
	for range ticker.C {
		data, err := updateAccShm()
		if err != nil {
			fmt.Printf("error: %e\n", err)
		}
		accChan <- data
	}
}
