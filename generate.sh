#!/usr/bin/env bash
root=`dirname $(readlink -e $0)`

JAVA_PACKAGE="bitflow4j.script.antlr.generated"
ANTLR_JAR="$root/lib/antlr-4.7.1-complete.jar"
OUT="$root/out"
OUT_JAVA="$OUT/java"
OUT_GO="$OUT/go"

if [ -z "$SOURCE_ONLY" ]; then
    rm -rf "$OUT"
    mkdir -p "$OUT"
    cp "$root/Bitflow.g4" "$OUT"

    # Work in subshells due to changing working directory
    (
    rm -rf "$OUT_JAVA" && mkdir -p "$OUT_JAVA" && cd "$OUT_JAVA"
    cp "$root/Bitflow.g4" .
    java -jar "$ANTLR_JAR" -package "$JAVA_PACKAGE" Bitflow.g4
    rm *.g4

    # Compile
    javac -cp "$ANTLR_JAR" -d . *.java
    )

    (
    rm -rf "$OUT_GO" && mkdir -p "$OUT_GO" && cd "$OUT_GO"
    cp "$root/Bitflow.g4" .
    java -jar "$ANTLR_JAR" -Dlanguage=Go -visitor Bitflow.g4 Bitflow.g4
    rm *.g4
    )
fi

