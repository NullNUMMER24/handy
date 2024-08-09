/*
Copyright © 2024 NullNUMMER24
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

type StandingTimer struct {
	Date    string   `json:"date"`
	Records []Record `json:"records"`
}

type Record struct {
	StartTime string `json:"start_time"`
	StopTime  string `json:"stop_time"`
}

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
		standingTimer, err := loadStandingTimer()
		if err != nil {
			log.Fatal(err)
		}
		newRecord := Record{
			StartTime: time.Now().Format(time.RFC3339),
		}
		standingTimer.addRecord(newRecord)
		err = saveStandingTimer(standingTimer)
		if err != nil {
			log.Fatal(err)
		}

	},
}

var StopStandingTimerCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops the timer",
	Run: func(cmd *cobra.Command, args []string) {
		standingTimer, err := loadStandingTimer()
		if err != nil {
			log.Fatal(err)
		}
		latestRecord := standingTimer.getLatestRecord()
		if latestRecord != nil {
			latestRecord.StopTime = time.Now().Format(time.RFC3339)
			err = saveStandingTimer(standingTimer)
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(StandingTimerCmd)
	StandingTimerCmd.AddCommand(StartStandingTimerCmd)
	StandingTimerCmd.AddCommand(StopStandingTimerCmd)

}

func loadStandingTimer() (StandingTimer, error) {
	data, err := ioutil.ReadFile(fmt.Sprintf("%s/StandingTimer.json", ProjectFiles))
	if err != nil {
		return StandingTimer{}, err
	}
	var standingTimer StandingTimer
	err = json.Unmarshal(data, &standingTimer)
	return standingTimer, err
}

func (st *StandingTimer) addRecord(record Record) {
	st.Records = append(st.Records, record)
}

func (st *StandingTimer) getLatestRecord() *Record {
	for i := len(st.Records) - 1; i >= 0; i-- {
		if st.Records[i].StopTime == "" {
			return &st.Records[i]
		}
	}
	return nil
}

func saveStandingTimer(standingTimer StandingTimer) error {
	data, err := json.MarshalIndent(standingTimer, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fmt.Sprintf("%s/StandingTimer.json", ProjectFiles), data, 0644)
	return err
}
