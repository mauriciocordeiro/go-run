package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	ics "github.com/arran4/golang-ical"
)

const (
	logFmtRun   = "%2d\t%2d Km\t%s\n"
	dateFormat  = "2006-01-02 15:04:05"
	eventFormat = "%d KM (%d/%d)"
	docUrl      = "https://drive.google.com/file/d/1wzPab2BlX4N_2vEJMdVu_alagE6pIlAt/view"
	banner      = `														  
________  ________                 ________  ___  ___  ________      
|\   ____\|\   __  \               |\   __  \|\  \|\  \|\   ___  \    
\ \  \___|\ \  \|\  \  ____________\ \  \|\  \ \  \\\  \ \  \\ \  \   
 \ \  \  __\ \  \\\  \|\____________\ \   _  _\ \  \\\  \ \  \\ \  \  
  \ \  \|\  \ \  \\\  \|____________|\ \  \\  \\ \  \\\  \ \  \\ \  \ 
   \ \_______\ \_______\              \ \__\\ _\\ \_______\ \__\\ \__\
    \|_______|\|_______|               \|__|\|__|\|_______|\|__| \|__|
	

`
)

var (
	now = time.Now()

	freq        = flag.Int("freq", 4, "Initial frequency of runs per week")
	maxFreq     = flag.Int("maxFreq", 7, "Max frequency of runs per week")
	longRun     = flag.Int("longRun", 10, "Long run distance (for saturday)")
	avgRun      = flag.Int("avgRun", 5, "Average run distance (for weekdays)")
	start       = flag.String("start", defStart(), "First race day [yyyy-MM-dd HH:mm:ss UTC]")
	stop        = flag.String("stop", defStop(), "Last race day [yyyy-MM-dd HH:mm:ss UTC]")
	addEachWeek = flag.Int("addEachWeek", 4, "Add a run day each x weeks")
	dir         = flag.String("dir", ".", "Directory for .ics file")
	verbose     = flag.Bool("v", false, "Verbose")
	showDoc     = flag.Bool("doc", false, "If true, shows 'Running Order of Operation' document")

	freqSchedule = map[int][]time.Weekday{
		2: {time.Tuesday, time.Thursday},
		3: {time.Monday, time.Wednesday, time.Friday},
		4: {time.Tuesday, time.Wednesday, time.Thursday, time.Saturday},
		5: {time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Saturday},
		6: {time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday},
		7: {time.Sunday, time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday},
	}
)

type Run struct {
	Day      time.Time
	Distance int
}

func main() {
	flag.Parse()
	fmt.Print(banner)

	if *showDoc {
		fmt.Printf("Running Order of Operation: %s\n", docUrl)
		return
	}

	plan, err := buildPlan()
	if err != nil {
		panic(err)
	}

	err = generateCalendar(plan)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Done! Go see it at %s\n", fileName())
}

func buildPlan() ([]Run, error) {
	var plan = []Run{}
	currentDate, err := parse(*start)
	if err != nil {
		return plan, err
	}

	finalDate, err := parse(*stop)
	if err != nil {
		return plan, err
	}

	finalDate = finalDate.AddDate(0, 0, 1)
	schedule := freqSchedule[*freq]

	distance := 0
	weekCount := 0

	for currentDate.Before(finalDate) {
		if currentDate.Weekday() == time.Saturday {
			distance = *longRun
			weekCount++
		} else {
			distance = *avgRun
		}

		if contains(schedule, currentDate.Weekday()) {
			plan = append(plan, Run{
				Day:      currentDate,
				Distance: distance,
			})
		}

		if weekCount == *addEachWeek {
			weekCount = 0
			if len(schedule) < *maxFreq {
				schedule = freqSchedule[len(schedule)+1]
			}
		}

		currentDate = currentDate.AddDate(0, 0, 1)
	}

	return plan, nil
}

func generateCalendar(plan []Run) error {
	calendar := ics.NewCalendar()
	calendar.SetMethod(ics.MethodRequest)

	for i, run := range plan {
		event := calendar.AddEvent(strconv.FormatInt(now.Unix()+int64(i), 10))
		event.SetSummary(fmt.Sprintf(eventFormat, run.Distance, i+1, len(plan)))
		event.SetCreatedTime(now)
		event.SetModifiedAt(now)
		event.SetStartAt(run.Day)

		log(i+1, run)
	}

	src := calendar.Serialize()

	file, err := os.Create(fileName())
	if err != nil {
		return err
	}

	file.WriteString(src)

	if err := file.Close(); err != nil {
		return err
	}

	return nil
}

func log(i int, run Run) {
	if *verbose {
		fmt.Printf(logFmtRun, i, run.Distance, run.Day.Format(dateFormat))
	}
}

func parse(src string) (time.Time, error) {
	date, err := time.Parse(dateFormat, src)
	if err != nil {
		return time.Time{}, err
	}

	return date, nil
}

func contains(week []time.Weekday, day time.Weekday) bool {
	for _, d := range week {
		if d == day {
			return true
		}
	}
	return false
}

func defStart() string {
	return now.Format(dateFormat)
}

func defStop() string {
	return now.AddDate(0, 3, 0).Format(dateFormat)
}

func fileName() string {
	return fmt.Sprintf("%s/go-run_-_plan.ics", *dir)
}
