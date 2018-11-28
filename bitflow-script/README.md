# Bitflow Script ANTLR

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

