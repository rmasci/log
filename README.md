# Log
This is the log from the standard go install. All it does is allow the format of the date string. 

```
	lgOut := log.New(lf, "", log.LstdFlags)
	lgOut.Format = "20060102 15:04:05 MST"
```
Results in: 
```
20211206-14:35:18.477 EST
```
if for some reason you don't want the date to print use a log.Format string of "-".

This is good when you want to use other packages that wrap log like lumberjack which is what I wrote this for.
