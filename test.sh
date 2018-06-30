#!/usr/bin/env bash

source ./compile.sh
cd out

for testFile in ../src/test/resources/*.txt; do
    echo "testing $testFile"
    cat ${testFile} | java org.antlr.v4.gui.TestRig Bitflow script
    echo "------------------------"
done

cd ..