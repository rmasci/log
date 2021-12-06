# Log
This is the log from the standard go install. All it does is allow the format of the  day month year. 

```
	lgOut := log.New(lf, "", log.LstdFlags)
	lgOut.Format = "20060102 15:04:05 MST"
```
Results in: 
```
20211206-14:35:18.477 EST
```
