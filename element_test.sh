#!/usr/bin/env bash

SAMPLE="""
{
    :7000;
    -;
    \"asd\";
}
"""

rm -rf out
mkdir -p out
cp src/main/antlr/ch/chrisport/bitflow/* out
cd out
java -jar /usr/local/lib/antlr-4.7.1-complete.jar ElementTester.g4
javac ElementTester*.java
echo ${SAMPLE} | java org.antlr.v4.gui.TestRig ElementTester test -gui
cd ..
