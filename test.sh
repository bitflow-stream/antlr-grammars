#!/usr/bin/env bash
root=`dirname $(readlink -e $0)`

export SOURCE_ONLY=true
source "$root/generate.sh"
cd "$OUT_JAVA"

for testFile in "$root/tests"/*.txt; do
    echo -e "\nTesting: $testFile\n"
    cat "$testFile" | java -cp .:"$ANTLR_JAR" org.antlr.v4.gui.TestRig "${JAVA_PACKAGE}.Bitflow" script
    echo "------------------------"
done

