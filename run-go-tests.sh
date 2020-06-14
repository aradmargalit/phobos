#!/bin/sh

cd server
make test
make format
make check