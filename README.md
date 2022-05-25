# golang-library

## Just a personal collection of golang libraries, algorithms, etc. A basic playground for testing (hopefully) useful code...

## Repository

git@github.com:mwiater/golang-library.git

## References

* Scaffolding/CLI
  * Cobra: https://github.com/spf13/cobra
* Logging
  * https://github.com/sirupsen/logrus
* SysInfo
  * https://github.com/elastic/go-sysinfo
* Data Structures and Algorithms 
  * https://www.golangprograms.com/data-structure-and-algorithms.html
  * https://github.com/JJ/pigo
  * https://golangbyexample.com/pi-value-golang/
* Http
  * https://medium.com/@javin.ambridge/scaffolding-a-new-golang-http-service-f88ab8466104
* Websockets
  * https://www.google.com/amp/s/www.freecodecamp.org/news/million-websockets-and-go-cc58418460bb/amp/
  * https://github.com/gobwas/ws
  * Gorilla: https://gist.github.com/tmichel/7390690
* Files
  * https://hisk.io/read-files-concurrently-in-go/
 
## Initial Setup

### Cobra

REF: https://ordina-jworks.github.io/development/2018/10/20/make-your-own-cli-with-golang-and-cobra.html

`go get github.com/spf13/cobra/cobra`

Issue: https://github.com/spf13/cobra/issues/1215

```
go get github.com/spf13/cobra/cobra
go: extracting github.com/spf13/cobra v1.0.0
go: finding github.com/spf13/cobra/cobra latest
go: extracting github.com/spf13/cobra/cobra v0.0.0-20200916152758-7f8e83d9366a
go get github.com/spf13/cobra/cobra: ambiguous import: found github.com/spf13/cobra/cobra in multiple modules:
        github.com/spf13/cobra v1.0.0 (/home/matt/go/pkg/mod/github.com/spf13/cobra@v1.0.0/cobra)
        github.com/spf13/cobra/cobra v0.0.0-20200916152758-7f8e83d9366a (/home/matt/go/pkg/mod/github.com/spf13/cobra/cobra@v0.0.0-20200916152758-7f8e83d9366a)

```

`go get -u github.com/spf13/cobra/cobra@v1.0.0`

by default, this should store the app in: `~/go`, e.g.: `/home/matt/go/bin/cobra`

Set the go path, assuming above:

```
export GOPATH="/home/matt/go"
export PATH=$PATH:$(go env GOPATH)/bin
```

Permanent:

`nano ~/.bashrc`

Add:

```
export GOPATH="/home/matt/go"
export PATH=$PATH:$(go env GOPATH)/bin
```

Initialize:

`source ~/.bashrc`

Init:

Assuming in base of repo, e.g.: `/home/matt/projects/golang-library/`

```
mkdir -p github.com/mwiater/golang-library
cd github.com/mwiater/golang-library
cobra init --pkg-name=github.com/mwiater/golang-library
```

#=>

```
Your Cobra applicaton is ready at
/home/matt/projects/golang-library/github.com/mwiater/golang-library
```

`ls -laF` #=>

```
...
drwxr-x--x 2 matt matt  4096 May 28 18:15 cmd/
-rw-rw-r-- 1 matt matt 11358 May 28 18:15 LICENSE
-rw-rw-r-- 1 matt matt   672 May 28 18:15 main.go
```

## Init mod: https://github.com/spf13/cobra/issues/910#issuecomment-512794416

`go mod init github.com/mwiater/golang-library`

This creates the necessary go.mod file needed in the repo, but not in the cobra docs for some reason.

## Commit and push:

```
git add -A
git commit -m "initial commit"
git push origin master 
```

Test: `go run main.go` #=>

```
go: finding github.com/spf13/viper v1.7.0
go: downloading github.com/spf13/viper v1.7.0
go: extracting github.com/spf13/viper v1.7.0
go: downloading gopkg.in/yaml.v2 v2.2.4
go: downloading github.com/subosito/gotenv v1.2.0
go: downloading gopkg.in/ini.v1 v1.51.0
go: downloading github.com/magiconair/properties v1.8.1
go: downloading golang.org/x/text v0.3.2
go: downloading golang.org/x/sys v0.0.0-20190624142023-c5567b49c5d0
go: extracting github.com/subosito/gotenv v1.2.0
go: extracting github.com/magiconair/properties v1.8.1
go: extracting gopkg.in/yaml.v2 v2.2.4
go: extracting gopkg.in/ini.v1 v1.51.0
go: extracting golang.org/x/sys v0.0.0-20190624142023-c5567b49c5d0
go: extracting golang.org/x/text v0.3.2
go: finding golang.org/x/sys v0.0.0-20190624142023-c5567b49c5d0
go: finding github.com/magiconair/properties v1.8.1
go: finding golang.org/x/text v0.3.2
go: finding github.com/subosito/gotenv v1.2.0
go: finding gopkg.in/ini.v1 v1.51.0
go: finding gopkg.in/yaml.v2 v2.2.4
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.
```

Modify the `root.go` file to print a message for testing:

Find:

```
  // Uncomment the following line if your bare application
  // has an action associated with it:
  //Run: func(cmd *cobra.Command, args []string) { },
```

Replace:

```
  // Uncomment the following line if your bare application
  // has an action associated with it:
  Run: func(cmd *cobra.Command, args []string) {fmt.Println("Hello CLI")},
```

Install app as CLI executable (Use full app path):

`go install github.com/mwiater/golang-library`

Should install here by default: `~/go/bin/golang-library`.

Since GOPATH has been added to PATH (at the beginning of this doc), the CLI command should now be available:

`golang-library` #=>

`Hello CLI`

## Adding commands:

NOTE: `go install github.com/mwiater/golang-library` must be run when updating cmd files...

cobra add bubblesort

### Creates: `cmd/bubblesort.go` scaffold...

Rebuild: `go install github.com/mwiater/golang-library`

Run: `golang-library bubblesort`

cobra add pi

### Creates: `cmd/pi.go` scaffold...

go install github.com/mwiater/golang-library

golang-library pi

cobra add json

### Creates: `cmd/json.go` scaffold...

go install github.com/mwiater/golang-library

golang-library json

# List Commands

`golang-library help` #=>

```
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  golang-library [flags]
  golang-library [command]

Available Commands:
  bubblesort  A bubblesort example.
  help        Help about any command
  pi          Calculate n-digits of Pi
  sysinfo     Get information about your host.

Flags:
      --config string   config file (default is $HOME/.golang-library.yaml)
  -h, --help            help for golang-library
  -t, --toggle          Help message for toggle
```

## Parallel: Monte Carlo (REF: https://www.soroushjp.com/2015/02/07/go-concurrency-is-not-parallelism-real-world-lessons-with-monte-carlo-simulations/amp/)

## Logging (REF: https://github.com/sirupsen/logrus)

cobra add logger

Rebuild: `go install github.com/mwiater/golang-library`

Run: `golang-library logger`