# GoFindGit
Inspired by https://github.com/internetwache/GitTools/blob/master/Finder/gitfinder.py, I decided to write this in Go.

## Note
Still a work in progress. Not the ideal version yet.


# Usage
Build the binary
```go build -o gofindgit main.go```

Run the binary
```./gofindgit <domains_text_file_name>```

Modes:
- git 
- env

# ASCII ARTTTTT
```
	 _____        _____      _    _____ _ _   
	/ ____|      / ____|    | |  / ____(_) |  
       | |  __  ___ | |  __  ___| |_| |  __ _| |_ 
       | | |_ |/ _ \| | |_ |/ _ \ __| | |_ | | __|
       | |__| | (_) | |__| |  __/ |_| |__| | | |_ 
	\_____|\___/ \_____|\___|\__|\_____|_|\__|
											  
	Get Git and other juicy low hanging fruits...
	Support/Contribute: https://github.com/ahboon/GoFindGit
Usage:
	./gofindgit <mode> <domains_text_file_name>

	Modes:
	- git 
		Scans for git repo exposed in web root. http://example.com/.git/ OR https://example.com/.git/
	- env
		Scans for .env exposed in web root. http://example.com/.env/ OR https://example.com/.env/
```

## Things to do
To do:
- [ ] Add limit to number of go routines
- [ ] Add functionality to dump git files
- [ ] Add additional support for other requests (just a random thought)


# Support?
I am also open to contributions, suggestions and feedback! 