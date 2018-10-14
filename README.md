# gitrepos

a simple cli that gets a list of git repositories for an owner from various sites (github, bitbucket, gitlab)

## install

`go get github.com/moustafab/gitrepos`

## usage

```For example:
 
 gitrepos <sitename> -o moustafab
 gitrepos <sitename> --owner moustafab
 
 Usage:
   gitrepos [command]
 
 Available Commands:
   github      list owner's repos at github.com
   help        Help about any command
 
 Flags:
       --config string   config file (default is $HOME/.gitrepos.yaml)
   -h, --help            help for gitrepos
   -o, --owner string    owner to query of the repositories
   -t, --toggle          Help message for toggle
 
 Use "gitrepos [command] --help" for more information about a command.
 ```
 
## supported clients
 
* github
 
## coming soon

1. bitbucket
2. gitlab
3. enterprise


