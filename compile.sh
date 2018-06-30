#!/usr/bin/env bash

rm -rf out
mkdir -p out
cp src/main/antlr/ch/chrisport/bitflow/* out
cd out
java -jar /usr/local/lib/antlr-4.7.1-complete.jar Bitflow.g4
javac Bitflow*.java
cd ..