# Bitflow Script ANTLR

## Getting Started

Version 4.7.1 of the Antlr JAR file is available in the `lib` subfolder and will be used by the provided scripts.

The `./generate.sh` script generates Parser and Lexer code for `Java` and `Go` into the `out` directory and compiles the Java version.
The `./test.sh` script parses all test cases in the `tests` directory.
The `./gui.sh` script can be used to visualize the tokenization of a test case (or any other Bitflow script) like so:

- `echo 'a -> b()' | ./gui.sh`
- `cat tests/0_simple.txt | ./gui.sh`

The `./gui.sh` and `./test.sh` scripts require `./generate.sh` to be run first.

## Grammar Features
Please see the test scripts for examples.

Done:
- Scripts containing at least 1 pipeline
- Pipelines can be named, (`mypipeline: input() -> output()`), this is useful especially for forks (=> named subpipelines)!
- A Pipeline consists of Transforms and Forks
- A Transform consists of: `transform_name(transform_parameters)[execution_config]`
- Fork consists of: `name_ending_with_fork(fork_parameters)[execution_config]{ <one or multiple pipelines> }`
- Configs and parameters are lists of key-value pairs (`some_key="some_value"`)
- Comments are ignord, following JAVA commenting syntax (`/* comment anywhere or */ // comment until end of line`)
- Multiple inputs wrapped in { } brackets are possible
- File, console and port have a "hardcoded" abbreviation, for example for http: `:7000` instead of `http(port=7000)`

