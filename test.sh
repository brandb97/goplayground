#!/bin/bash

go import test.go > testtmp.go
cat testtmp.go > test.go
rm testtmp.go

go run test.go
