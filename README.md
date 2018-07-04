# Bitflow Script ANTLR

## Getting Started

For installation of ANTLR4, see [Installation page](https://github.com/antlr/antlr4/blob/master/doc/getting-started.md)

compile.sh expects ANTLR4 to be located under /usr/local/lib/antlr-4.7.1-complete.jar, test scripts are located in `src/test/resources`.   

## Commands
Test all test scripts against the grammar: `./test.sh`   
Visualize the tokenization of one test script: `./gui.sh <test-number>` (e.g. `./gui.sh 0` to visualize 0_simple.txt)

## Grammar Features
Please see the test scripts for examples.

Done:
- Scripts containing at least 1 pipeline
- Pipelines can be named, (`mypipeline: input() -> output()`), this is useful especially for forks (=> named subpipelines)!
- A Pipeline consists of Transforms and Forks
- A Transform consists of: `transform_name(transform_parameters)[execution_config]`
- Fork consists of: `name_ending_with_fork(fork_parameters)[execution_config]{ <one or multiple pipelines> }`
- configs and parameters are lists of key-value pairs (`some_key="some_value"`)
- Comments are ignord, following JAVA commenting syntax (`/* comment anywhere or */ // comment until end of line`)
- Multiple inputs wrapped in { } brackets are possible
- file, console and port have a "hardcoded" abbreviation, for example for http: `:7000` instead of `http(port=7000)`
