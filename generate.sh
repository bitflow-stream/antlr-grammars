#!/usr/bin/env bash
root=`dirname $(readlink -e $0)`

function help() {
    echo "Possible commands: generate (default) | test script | gui script | gui query"
    exit 1
}

test $# -gt 2 && help
cmd="$1 $2"

ANTLR_JAR="$root/lib/antlr-4.7.1-complete.jar"
ANTLR_CLASS="org.antlr.v4.gui.TestRig"

SCRIPT_JAVA_PACKAGE="bitflow4j.script.generated"
SCRIPT_GO_PACKAGE="parser"
QUERY_JAVA_PACKAGE="bitflow4j.steps.query.generated"
QUERY_GO_PACKAGE="parser"

function generate() {
    grammar="$1"
    target="$2"
    shift 2

    rm -rf "$target"
    mkdir -p "$target"
    (
        cd $(dirname "$grammar")
        java -jar "$ANTLR_JAR" -o "$target" $@ $(basename "$grammar")
    )
}

function generate_java() {
    target=$(readlink -e $2)
    generate $@
    javac -cp "$ANTLR_JAR" $(find "$target" -name '*.java')
#-d "$target" 
}

case $cmd in
    " "|"generate ")
        # Bitflow Script
        script="$root/bitflow-script"
        generate_java "$script/Bitflow.g4" "$script/generated/java/bitflow4j/script/generated" -package $SCRIPT_JAVA_PACKAGE -visitor
        generate      "$script/Bitflow.g4" "$script/generated/go"   -package $SCRIPT_GO_PACKAGE -Dlanguage=Go -visitor

        # Bitflow Query Language
        query="$root/bitflow-query-language"
        generate_java "$query/BitflowQuery.g4" "$query/generated/java/bitflow4j/steps/query/generated" -package $QUERY_JAVA_PACKAGE -visitor
        generate      "$query/BitflowQuery.g4" "$query/generated/go"   -package $QUERY_GO_PACKAGE -Dlanguage=Go -visitor
        ;;
    "test script")
        cd "$root/bitflow-script/generated/java"
        for testFile in "$root/bitflow-script/tests"/*.txt; do
            echo -e "\nTesting: $testFile\n"
            cat "$testFile" | java -cp .:"$ANTLR_JAR" $ANTLR_CLASS "$SCRIPT_JAVA_PACKAGE.Bitflow" script
            echo "------------------------"
        done
        ;;
    "gui script")
        cd "$root/bitflow-script/generated/java"
        java -cp .:"$ANTLR_JAR" $ANTLR_CLASS "${SCRIPT_JAVA_PACKAGE}.Bitflow" script -tokens -gui
        ;;
    "gui query")
        cd "$root/bitflow-query-language/generated/java"
        java -cp .:"$ANTLR_JAR" $ANTLR_CLASS "${QUERY_JAVA_PACKAGE}.BitflowQuery" parse -tokens -gui
        ;;
    *)
        echo "Unknown command: '$cmd'"
        help
        ;;
esac

