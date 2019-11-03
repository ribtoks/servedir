# servedir

[![Build Status](https://travis-ci.org/ribtoks/servedir.svg?branch=master)](https://travis-ci.org/ribtoks/servedir)
![license](https://img.shields.io/badge/license-MIT-blue.svg)
![copyright](https://img.shields.io/badge/%C2%A9-Taras_Kushnir-blue.svg)
![language](https://img.shields.io/badge/language-go-blue.svg)


Go has wonderful builtin web server and file server but there's no easy way to use it except of writing own code. This mini package solves this problem.

## Install

`go get -u github.com/ribtoks/servedir`

## Usage

Just `servedir` will serve local directory on localhost on port `8080`. Other options:

    Usage of servedir:
      -d string
        	Path to directory to serve (default "./")
      -p int
        	Port to use (default 8080)
      -r string
        	Http path to serve (default "/")
