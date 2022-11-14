# go-run
An .ics calendar generator for running plans

## Usage

```bash
$ go-run -help
  -addEachWeek int
        Add a run day each x weeks (default 4)
  -avgRun int
        Average run distance (for weekdays) (default 5)
  -dir string
        Directory for .ics file (default ".")
  -doc
        If true, shows 'Running Order of Operation' document
  -freq int
        Initial frequency of runs per week (default 4)
  -level int
        Runner level (default 2)
  -longRun int
        Long run distance (for saturday) (default 10)
  -maxFreq int
        Max frequency of runs per week (default 7)
  -start string
        First race day [yyyy-MM-dd HH:mm:ss UTC] (default "2022-11-14 10:46:16")
  -stop string
        Last race day [yyyy-MM-dd HH:mm:ss UTC] (default "2023-02-14 10:46:16")
  -v    Verbose
```

## Running Order of Operations

The calendar is generate based on [this doument](https://drive.google.com/file/d/1wzPab2BlX4N_2vEJMdVu_alagE6pIlAt/view).
