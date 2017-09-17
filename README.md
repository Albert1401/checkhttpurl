# checkhttpurl
```bash
$ go build checkurl.go
$ ./checkurl
Usage:
	checkurl [-t timeout] [-i interval] url
Examples:
	checkurl -t 1m http://ya.ru
	checkurl -t 45s -i 2s http://mail.ru
	checkurl https://www.jetbrains.com/go/
```
