#!/usr/bin/env bash
root=`dirname $(readlink -e $0)`

export SOURCE_ONLY=true
source "$root/generate.sh"
cd "$OUT_JAVA"
javac -cp "$ANTLR_JAR" -d . *.java

java -cp .:"$ANTLR_JAR" org.antlr.v4.gui.TestRig "${JAVA_PACKAGE}.Bitflow" script -tokens -gui

