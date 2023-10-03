package main

import (
	"flag"
	"fmt"
	"goAccFuel/acc"
	"goAccFuel/fuel"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

const updateIntervalSeconds = 1

func main() {

	extraLap := flag.Int("l", 2, "add this extra laps to the fuel calculation")
	flag.Parse()
	fmt.Printf("extra laps: %v sec\n", *extraLap)

	accChan := make(chan acc.AccData)
	defer close(accChan)

	m := fuel.NewFuelModel()

	// start acc update go function

	go func() {
		go acc.StartUpdater(updateIntervalSeconds, accChan)
		for {
			d := <-accChan
			m.Update(fuel.UpdateAcc(d))
			m.View()
		}
	}()

	// start ui

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
