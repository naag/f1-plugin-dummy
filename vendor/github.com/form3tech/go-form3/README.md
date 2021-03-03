# Form 3 private Go client

**YOU SHOULD PROBABLY USE [form3tech-oss/go-form3](https://github.com/form3tech-oss/go-form3) UNLESS YOU NEED PRIVATE APIS.**

This project provides a client for Form3's APIs, including private endpoints. It is the sibling of 
[form3tech-oss/go-form3](https://github.com/form3tech-oss/go-form3) which provides only public APIs.

  [![Build Status](https://travis-ci.org/form3tech/go-form3.svg?branch=master)](https://travis-ci.org/form3tech/go-form3)

# Developing

## Build locally
* `git clone git@github.com:form3tech/go-form3.git`
* `cd go-form3`
* `make build`

# Swagger
The swagger file in this project is currently maintained and updated manually.

To generate the swagger model run: `swagger generate client -f ./swagger.yaml`
