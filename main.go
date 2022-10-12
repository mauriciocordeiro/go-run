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
)

var (
	now = time.Now()

	level       = flag.Int("level", 2, "Runner level")
	freq        = flag.Int("freq", 4, "Initial frequency of runs per week")
	maxFreq     = flag.Int("maxFreq", 7, "Max frequency of runs per week")
	longRun     = flag.Int("longRun", 10, "Long run distance (for saturday)")
	avgRun      = flag.Int("avgRun", 5, "Average run distance (for weekdays)")
	start       = flag.String("start", now.String(), "First race day [yyyy-MM-dd HH:mm:ss UTC]")
	stop        = flag.String("stop", now.AddDate(0, 6, 0).String(), "Last race day [yyyy-MM-dd HH:mm:ss UTC]")
	dir         = flag.String("dir", ".", "Directory for .ics file")
	addEachWeek = flag.Int("addEachWeek", 4, "Add a run day each x weeks")

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

	// TODO: validate flags

	plan := buildPlan()
	generateCalendar(plan)
}

func buildPlan() []Run {
	var plan = []Run{}
	currentDate := parse(*start)
	finalDate := parse(*stop).AddDate(0, 0, 1)
	schedule := freqSchedule[*freq]

	distance := 0
	weekCount := 0

	for currentDate.Before(finalDate) {
		switch *level {
		case 1:
			fmt.Printf("Level %d not imlemented.\n", *level)
		case 2:
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
		case 3:
			fmt.Printf("Level %d not imlemented.\n", *level)
		case 4:
			fmt.Printf("Level %d not imlemented.\n", *level)
		case 5:
			fmt.Printf("Level %d not imlemented.\n", *level)
		case 6:
			fmt.Printf("Level %d not imlemented.\n", *level)
		}

		currentDate = currentDate.AddDate(0, 0, 1)
	}

	return plan
}

func generateCalendar(plan []Run) {
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

	file, err := os.Create(fmt.Sprintf("%s/go-run %s.ics", *dir, now.Format(dateFormat)))
	if err != nil {
		panic(fmt.Sprintf("Error creating .ics file\n%s", err.Error()))
	}
	defer file.Close()

	file.WriteString(src)
}

func log(i int, run Run) {
	fmt.Printf(logFmtRun, i, run.Distance, run.Day.Format(dateFormat))
}

func parse(src string) time.Time {
	date, err := time.Parse(dateFormat, src)
	if err != nil {
		panic(fmt.Sprintf("Error parsing date param. %s\n%s", src, err.Error()))
	}

	return date
}

func contains(week []time.Weekday, day time.Weekday) bool {
	for _, d := range week {
		if d == day {
			return true
		}
	}
	return false
}
