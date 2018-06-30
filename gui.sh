#!/usr/bin/env bash

source ./compile.sh
cd out
cat ../src/test/resources/$1*.txt | java org.antlr.v4.gui.TestRig Bitflow script -gui
cd ..