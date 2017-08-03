# protemplates
This creates a simple project structure for different projects. ( similar to Cargo from rust )

---


Usage:

```
$ protemplates
Usage:
  protemplates [command]

Available Commands:
  create      Creates a project template for a given language.
  help        Help about any command

Flags:
  -h, --help   help for protemplates

Use "protemplates [command] --help" for more information about a command.
```


----
Example:

```
$ protemplates create python
Please enter a desired project name:
```

---
For convenience a binary is also committed in the github repo, under prebuilt directory, compiled using:

```
GOOS=linux go build -ldflags="-s -w" github.com/ansrivas/protemplates
```
