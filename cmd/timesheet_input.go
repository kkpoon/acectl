package cmd

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	aceproject "github.com/kkpoon/go-aceproject"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// inputCmd represents the input command
var timesheetInputCmd = &cobra.Command{
	Use:   "input",
	Short: "timesheet input",
	Long:  `save input to ACE project timesheet`,
	Run: func(cmd *cobra.Command, args []string) {
		guid := viper.GetString("GUID")
		guidInfo := aceproject.GUIDInfo{GUID: guid}
		timesheetSvc := aceproject.NewTimesheetService(&http.Client{}, &guidInfo)

		hours := [7]float64{}
		hoursStr := strings.Split(viper.GetString("hours"), ",")
		if len(hoursStr) != 7 {
			fmt.Printf("invalid format in `hours` parameter")
			os.Exit(1)
		} else {
			for i, v := range hoursStr {
				if num, err := strconv.ParseFloat(v, 64); err == nil {
					hours[i] = num
				}
			}
		}

		weekStart := time.Now()
		if weekStartStr := viper.GetString("date"); weekStartStr != "" {
			inputDate, err := time.Parse("2006-01-02", weekStartStr)
			if err != nil {
				fmt.Printf("invalid format in `date` parameter")
				os.Exit(1)
			}
			weekStart = inputDate
		}

		timeTypeStr := strings.ToLower(viper.GetString("timetype"))
		var timeType int64 = 1
		switch {
		case timeTypeStr == "regular":
			timeType = 1
		case timeTypeStr == "training":
			timeType = 2
		case timeTypeStr == "overtime":
			timeType = 3
		default:
			timeType = 1
		}

		taskID := viper.GetInt64("taskid")
		if taskID == 0 {
			fmt.Printf("invalid `taskid` parameter")
			os.Exit(1)
		}

		comments := viper.GetString("comments")

		workItem := aceproject.SaveWorkItem{
			WeekStart:  weekStart.Format("2006-01-02"),
			TaskID:     taskID,
			TimeTypeID: timeType,
			HoursDay1:  hours[0],
			HoursDay2:  hours[1],
			HoursDay3:  hours[2],
			HoursDay4:  hours[3],
			HoursDay5:  hours[4],
			HoursDay6:  hours[5],
			HoursDay7:  hours[6],
			Comments:   &comments,
		}

		_, err := timesheetSvc.SaveWorkItem(&workItem)

		if err != nil {
			fmt.Printf("save work item error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("saved")
	},
}

func init() {
	timesheetCmd.AddCommand(timesheetInputCmd)
	timesheetInputCmd.Flags().StringP("date", "d", "", "a date of the week you want to log time, e.g. 2016-08-31, default now")
	timesheetInputCmd.Flags().Int64P("taskid", "t", 0, "the task ID")
	timesheetInputCmd.Flags().StringP("timetype", "", "regular", "the type of the time to log, regular/training/overtime")
	timesheetInputCmd.Flags().StringP("comments", "c", "", "comments of this log entry")
	timesheetInputCmd.Flags().StringP("hours", "", "0,0,0,0,0,0,0", "7-days hours represent in comma separated format")
	viper.BindPFlag("date", timesheetInputCmd.Flags().Lookup("date"))
	viper.BindPFlag("taskid", timesheetInputCmd.Flags().Lookup("taskid"))
	viper.BindPFlag("timetype", timesheetInputCmd.Flags().Lookup("timetype"))
	viper.BindPFlag("comments", timesheetInputCmd.Flags().Lookup("comments"))
	viper.BindPFlag("hours", timesheetInputCmd.Flags().Lookup("hours"))
}
