/*
Copyright © 2024 NullNUMMER24
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

type StandingTimer struct {
	Date         string        `json:"date"`
	LastStart    string        `json:"last_start"`
	TimeStanding time.Duration `json:"time_standing"`
	Breaks       int           `json:"breaks"`
}

var NewEntry StandingTimer

// StandingTimerCmd represents the StandingTimer command
var StandingTimerCmd = &cobra.Command{
	Use:   "StandingTimer",
	Short: "A stopwatch to count how long you are standing",
	Long:  `This is a function which should help you to avoid sitting to much in the office.`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := os.OpenFile(fmt.Sprintf("%s/StandingTimer.json", ProjectFiles), os.O_CREATE, 0755)
		if err != nil {
			log.Printf("Error creating file: %sStandingTimer.json", ProjectFiles)
		}
	},
}

var StartStandingTimerCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the timer",
	Run: func(cmd *cobra.Command, args []string) {
		existingData, err := loadStandingTimer()
		if err != nil {
			log.Fatal(err)
		}

		if existingData.Date == time.Now().Format("2006-01-02") {
			NewEntry = StandingTimer{
				Date:         time.Now().Format("2006-01-02"),
				LastStart:    time.Now().Format("15:04:03"),
				TimeStanding: existingData.TimeStanding,
				Breaks:       existingData.Breaks + 1,
			}
		} else {
			NewEntry = StandingTimer{
				Date:      time.Now().Format("2006-01-02"),
				LastStart: time.Now().Format("15:04:03"),
				Breaks:    0,
			}
		}

		newEntryJson, err := json.MarshalIndent(NewEntry, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		err = os.WriteFile(fmt.Sprintf("%s/StandingTimer.json", ProjectFiles), newEntryJson, 0644)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var StopStandingTimerCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops the timer",
	Run: func(cmd *cobra.Command, args []string) {
		existingData, err := loadStandingTimer()
		if err != nil {
			log.Fatal(err)
		}
		startTime, _ := time.Parse("2006-01-02 15:04", existingData.Date+" "+existingData.LastStart)
		endTime := time.Now()
		totalTime := endTime.Sub(startTime)
		if existingData.Date == time.Now().Format("2006-01-02") {
			NewEntry = StandingTimer{
				Date:         time.Now().Format("2006-01-02"),
				LastStart:    existingData.LastStart,
				TimeStanding: existingData.TimeStanding + totalTime,
				Breaks:       existingData.Breaks + 1,
			}
		} else {
			NewEntry = StandingTimer{
				Date:         time.Now().Format("2006-01-02"),
				LastStart:    time.Now().Format("15:04:03"),
				TimeStanding: totalTime,
				Breaks:       1,
			}
		}

		// Convert TimeStanding to seconds
		seconds := int(NewEntry.TimeStanding.Seconds())

		// Create a temporary struct to hold the data
		tempStruct := struct {
			Date         string `json:"date"`
			LastStart    string `json:"last_start"`
			TimeStanding int    `json:"time_standing"`
			Breaks       int    `json:"breaks"`
		}{
			Date:         NewEntry.Date,
			LastStart:    NewEntry.LastStart,
			TimeStanding: seconds,
			Breaks:       NewEntry.Breaks,
		}

		// Marshal the temporary struct to JSON
		newEntryJson, err := json.MarshalIndent(tempStruct, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		err = os.WriteFile(fmt.Sprintf("%s/StandingTimer.json", ProjectFiles), newEntryJson, 0644)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(StandingTimerCmd)
	StandingTimerCmd.AddCommand(StartStandingTimerCmd)
	// StandingTimerCmd.AddCommand(StopStandingTimerCmd)

}

func loadStandingTimer() (StandingTimer, error) {
	filePath := fmt.Sprintf("%s/StandingTimer.json", ProjectFiles)
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return StandingTimer{}, nil
		} else {
			return StandingTimer{}, err
		}
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return StandingTimer{}, err
	}

	var existingData StandingTimer
	err = json.Unmarshal(data, &existingData)
	if err != nil {
		return StandingTimer{}, err
	}

	return existingData, nil
}
