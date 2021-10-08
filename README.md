# Golang Concurrency Curriculum

All the program start form main file with parameter as `program` and `inner-program` if any.  

### How to run each program

1) To run first program
```sh
$ go run main.go -program=1
```
output:
```sh
^CGoodbye
CPU consumed by this process 18.145971% [busy: 10216.000000, total: 56299.000000]
```

2) To run Second program
```sh
$ go run main.go -program=2
```
output:
```sh
^CGoodbye 0
^CGoodbye 1
^CGoodbye 2
^CGoodbye 3
^CGoodbye 4
^CGoodbye 5
^CGoodbye 6
^CGoodbye 7
^CGoodbye 8
^CGoodbye 9
^CGoodbye 10
^CGoodbye 11
^CGoodbye 12
^CGoodbye 13
^CGoodbye 14
^CGoodbye 15
CPU usage is 19.812806% [busy: 6922.000000, total: 34937.000000]
```

3) To run Third program
```sh
$ go run main.go -program=3
```
output:
```sh
Started Four
Started One
Started Two
Started Three
All done.
```

4) To run Fourth program
```sh
$ go run main.go -program=4
```
output:
```sh
four
one
two
three
All done.
```

5) To run Fifth program
```sh
$ go run main.go -program=5
```
output:
```sh
50.033347798%  
```


6) To run Sixth program
```sh
$ go run main.go -program=6
```
output:
```sh
^CCaught interrupt
All done 

Goodbye! timeout
All done 
```


7.1) To run Seventh program
```sh
$ go run main.go -program=7 --inner-program=1
```
output:
```sh
1
7
7
9
1
25.015411993
map[1:8 7:56 9:72]
```

7.2) To run Seventh program
```sh
$ go run main.go -program=7 --inner-program=2
```
output:
```sh
25.019348224
map[1:8 7:56 9:72]
```

7.3) To run Seventh program
```sh
$ go run main.go -program=7 --inner-program=3
```
output:
```sh
25.018396326
[7 1]
```

8) To run Eighth program
```sh
$ go run main.go -program=8
```
output:
```sh
job 1
Pool increased/decreased by 1. Current Pool size: 2
^Csignal received
job 3
Pool increased/decreased by 1. Current Pool size: 1
^Csignal received
job 4
Pool increased/decreased by 1. Current Pool size: 2
^Csignal received
job 6
Pool increased/decreased by 1. Current Pool size: 1
^Csignal received
job 7
Pool increased/decreased by 1. Current Pool size: 2
^Csignal received
job 9
Pool increased/decreased by 1. Current Pool size: 1
^Csignal received
job 10
Pool increased/decreased by 1. Current Pool size: 2
^Csignal received
job 12
Pool increased/decreased by 1. Current Pool size: 1
^Csignal received
job 13
Pool increased/decreased by 1. Current Pool size: 2
^Csignal received
job 15
Pool increased/decreased by 1. Current Pool size: 1
^Csignal received
job 16
Pool increased/decreased by 1. Current Pool size: 2
^Csignal received
job 18
Pool increased/decreased by 1. Current Pool size: 1
^Csignal received
job 19
Pool increased/decreased by 1. Current Pool size: 2
^Csignal received
job 21
Pool increased/decreased by 1. Current Pool size: 1
^Csignal received
job 22
Pool increased/decreased by 1. Current Pool size: 2
^Csignal received
job 24
Pool increased/decreased by 1. Current Pool size: 1
^Csignal received
job 25
Pool increased/decreased by 1. Current Pool size: 2
^Csignal received
job 27
Pool increased/decreased by 1. Current Pool size: 1
^Csignal received
job 28
Pool increased/decreased by 1. Current Pool size: 2
^Csignal received
job 30
Pool increased/decreased by 1. Current Pool size: 1
^Csignal received
job 31
Pool increased/decreased by 1. Current Pool size: 2
^Csignal received
job 33
Pool increased/decreased by 1. Current Pool size: 1
^Csignal received
job 34
Pool increased/decreased by 1. Current Pool size: 2
^Csignal received
job 36
Pool increased/decreased by 1. Current Pool size: 1
^Csignal received
job 37
Pool increased/decreased by 1. Current Pool size: 2
^Csignal received
job 39
Pool increased/decreased by 1. Current Pool size: 1
^Csignal received
job 40
Pool increased/decreased by 1. Current Pool size: 2
^Csignal received
job 42
Pool increased/decreased by 1. Current Pool size: 1
^Csignal received
job 43
Pool increased/decreased by 1. Current Pool size: 2
^Csignal received
job 45
Pool increased/decreased by 1. Current Pool size: 1
^Csignal received
job 46
Pool increased/decreased by 1. Current Pool size: 2
^Csignal received
job 48
Pool increased/decreased by 1. Current Pool size: 1
^Csignal received
job 49
Pool increased/decreased by 1. Current Pool size: 2
^Csignal received
job 51
Pool increased/decreased by 1. Current Pool size: 1
^Csignal received
job 52
Pool increased/decreased by 1. Current Pool size: 2
^Csignal received
job 54
Pool increased/decreased by 1. Current Pool size: 1
^Csignal received
job 55
Pool increased/decreased by 1. Current Pool size: 2
^Csignal received
job 57
Pool increased/decreased by 1. Current Pool size: 1
^Csignal received
job 58
Pool increased/decreased by 1. Current Pool size: 2
^Csignal received
job 60
Pool increased/decreased by 1. Current Pool size: 1
^Csignal received
job 61
Pool increased/decreased by 1. Current Pool size: 2
^Csignal received
job 63
Pool increased/decreased by 1. Current Pool size: 1
^Csignal received
job 64
Pool increased/decreased by 1. Current Pool size: 2
^Csignal received
job 66
Pool increased/decreased by 1. Current Pool size: 1
^Csignal received
job 67
Pool increased/decreased by 1. Current Pool size: 2
^Csignal received
job 69
Pool increased/decreased by 1. Current Pool size: 1
^Csignal received
job 70
Pool increased/decreased by 1. Current Pool size: 2
^Csignal received
job 72
Pool increased/decreased by 1. Current Pool size: 1
^Csignal received
job 73
Pool increased/decreased by 1. Current Pool size: 2
^Csignal received
job 75
Pool increased/decreased by 1. Current Pool size: 1
^Csignal received
job 76
Pool increased/decreased by 1. Current Pool size: 2
^Csignal received
job 78
Pool increased/decreased by 1. Current Pool size: 1
^Csignal received
job 79
Pool increased/decreased by 1. Current Pool size: 2
^Csignal received
job 81
Pool increased/decreased by 1. Current Pool size: 1
^Csignal received
job 82
Pool increased/decreased by 1. Current Pool size: 2
^Csignal received
job 84
Pool increased/decreased by 1. Current Pool size: 1
^Csignal received
job 85
Pool increased/decreased by 1. Current Pool size: 2
^Csignal received
job 87
Pool increased/decreased by 1. Current Pool size: 1
^Csignal received
job 88
Pool increased/decreased by 1. Current Pool size: 2
^Csignal received
job 90
Pool increased/decreased by 1. Current Pool size: 1
^Csignal received
job 91
Pool increased/decreased by 1. Current Pool size: 2
^Csignal received
job 93
Pool increased/decreased by 1. Current Pool size: 1
^Csignal received
job 94
Pool increased/decreased by 1. Current Pool size: 2
^Csignal received
job 96
Pool increased/decreased by 1. Current Pool size: 1
^Csignal received
job 97
Pool increased/decreased by 1. Current Pool size: 2
^Csignal received
job 99
Pool increased/decreased by 1. Current Pool size: 1
^Csignal received
32.60549968%         
```