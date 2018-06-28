#!/usr/bin/env bash

source ./compile.sh
cd out
cat ../src/test/resources/$1.txt | grun Bitflow pipeline -gui
cd ..