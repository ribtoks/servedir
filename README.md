# servedir

Go has wonderful builtin web server and file server but there's no easy way to use it except of writing own code. This mini package solves this problem.

## Install

`go get github.com/ribtoks/servedir`

## Usage

Just `servedir` will serve local directory on localhost on port `8080`. Other options:

    Usage of servedir:
      -d string
        	Path to directory to serve (default "./")
      -p int
        	Port to use (default 8080)
      -r string
        	Http path to serve (default "/")
