#!/usr/bin/env bash
# from ~/.zshrc
# export CLASSPATH=".:/usr/local/lib/antlr-4.7.1-complete.jar:$CLASSPATH"
# alias antlr4='java -jar /usr/local/lib/antlr-4.7.1-complete.jar'
# alias grun='java org.antlr.v4.gui.TestRig'

rm -rf out
mkdir -p out
cp src/main/antlr/ch/chrisport/bitflow/* out
cd out
antlr4 Bitflow.g4
javac Bitflow*.java
cd ..