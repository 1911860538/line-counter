# line-counter
&nbsp;&nbsp;&nbsp;&nbsp; 
The line-counter is a Go-based command line tool, 
that quickly analyzes directories and provides detailed statistics 
on file types, line counts, empty lines, and so on. 
It's a handy tool for developers and administrators 
seeking insights into their codebase. 

## Getting started
### Install
```shell
go install github.com/1911860538/line-counter@latest
```

### Usage
```text
$ line-counter --help
Usage of line-counter:
  -in string
        ignore directories or files from statistics.
        For example: a/b/,a/c/x.txt
  -it string
        ignore types from statistics.
        For example: json,xml
  -n string
        target directory or file for statistics. Default is the current directory.
  -t string
        file types for statistics.
        For example: txt,go,py
```

### Run
```text
$ line-counter -in=vendor -it=exe,mod
2024/04/12 11:03:47 The target is `/home/someone/code/line-counter`.

+---------------------------------------------------------------------------------------------------------------------------------------------+
|                                                                line-counter                                                                 |
+-----------+-------+----------+---------+---------+---------+-------+----------+----------+----------+-----------+--------------+------------+
| Extension | Count | SizeSum  | SizeMin | SizeMax | SizeAvg | Lines | LinesMin | LinesMax | LinesAvg | LinesCode | LinesComment | LinesBlank |
+-----------+-------+----------+---------+---------+---------+-------+----------+----------+----------+-----------+--------------+------------+
| .go       | 39    | 15.66 KB | 96 B    | 2.73 KB | 411 B   | 903   | 6        | 120      | 23       | 774       | 0            | 129        |
| .md       | 1     | 851 B    | 851 B   | 851 B   | 851 B   | 35    | 35       | 35       | 35       | 26        | 5            | 4          |
| .sum      | 1     | 905 B    | 905 B   | 905 B   | 905 B   | 10    | 10       | 10       | 10       | 10        | 0            | 0          |
+-----------+-------+----------+---------+---------+---------+-------+----------+----------+----------+-----------+--------------+------------+

2024/04/12 11:03:47 Done! Successfully calculated statistic in 8 milliseconds.
```

### Notice
&nbsp;&nbsp;&nbsp;&nbsp;
Files without suffixes and files or directories whose names start with  "." 
will be ignored by default,
unless they are the target files or directories.
