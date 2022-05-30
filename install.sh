#!/bin/bash

set -e

go build
sudo cp ./kulana /usr/bin/.