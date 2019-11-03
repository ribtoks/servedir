# servedir

[![Build Status](https://travis-ci.org/ribtoks/servedir.svg?branch=master)](https://travis-ci.org/ribtoks/servedir)
![license](https://img.shields.io/badge/license-MIT-blue.svg)
![copyright](https://img.shields.io/badge/%C2%A9-Taras_Kushnir-blue.svg)
![language](https://img.shields.io/badge/language-go-blue.svg)

## About

`servedir` allows you to serve local directory over http with one command.

Go has wonderful builtin web server and file server but there's no easy way to use it except of writing own code. This mini package solves this problem.

## Install

`go get -u github.com/ribtoks/servedir`

## Usage

    Usage of servedir:
      -d string
        	Path to directory to serve (default "./")
      -p int
        	Port to use (default 8080)
      -r string
        	Http path to serve (default "/")

## Examples

Serve current directory on `http://locahost:8080/`:

    servedir
    
Serve directory `~/Photos/2019-11-03/` with path `photos` on port `1313` (`http://locahost:1313/photos/`):

    servedir -d ~/Photos/2019-11-03 -p 1313 -r photos
