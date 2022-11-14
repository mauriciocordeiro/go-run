# go-run
An CLI .ics calendar generator for running plans.

## Build

```sh
go build main.go
```

## Usage

```sh
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

### Example

```sh
go-run -level 2 -start '2023-01-02 09:00:00' -stop '2023-06-02 09:00:00' -maxFreq 5 -avgRun 5 -longRun 10 -addEachWeek 4 -freq 3
```

This command will generate events from start to stop date, 3 runs per week (two 5k runs (`avgRun`) and one 10k (`longRun`)) and will increase an `avgRun` each 4 weeks until a top 5 runs per week.


### Run Schedule

| freq | Mon      | Tue      | Wed      | Thu      | Fri      | Sat      | Sun      |
|------|----------|----------|----------|----------|----------|----------|----------|
| 2    |          | :runner: |          | :runner: |          |          |          |
| 3    | :runner: |          | :runner: |          | :runner: |          |          |
| 4    |          | :runner: | :runner: | :runner: |          | :runner: |          |
| 5    | :runner: | :runner: | :runner: | :runner: |          | :runner: |          |
| 6    | :runner: | :runner: | :runner: | :runner: | :runner: | :runner: |          |
| 7    | :runner: | :runner: | :runner: | :runner: | :runner: | :runner: | :runner: |

This is how the runs are scheduled by the frequency defined. 

## Running Order of Operations

The calendar is generate slitly based on [this doument](https://drive.google.com/file/d/1wzPab2BlX4N_2vEJMdVu_alagE6pIlAt/view).
