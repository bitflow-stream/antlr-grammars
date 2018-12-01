# Bitflow Script and Bitflow Query Language

This repository contains the ANTLR grammars for two doman specific languages:
- *Bitflow Script* is implemented in the `bitflow-script` directory. The language can define a graph of interconnected, parameterized data processing steps.
- *Bitflow Query Language* (`bitflow-query-language`) is a SQL-inspired language for transforming a data stream, which can be used as one step within a pipeline defined through the Bitflow Script.

The 4.7.1 version of the ANTLR library is located in the `lib` folder and used by the `generate.sh` script to generate and compile the parser code for the two DSLs.
Currently, Java and Go are used as target languages.
The `generate.sh` script can also be started with `gui script` or `gui query` parameters, which can be used to visualize and debug code of the respective DSL (the code is piped into the script over the standard input). Examples:

- `echo 'a -> b()' | ./generate.sh gui script`
- `cat bitflow-script/tests/0_simple.txt | ./generate.sh gui script`
- `echo "select * where foo = 'bar'" | ./generate.sh gui query`
