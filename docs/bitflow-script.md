# Bitflow Script

Bitflow Script is a domain specific language for defining data processing pipelines.

A Bitflow Script mostly ignores white space (spaces, tabs and newlines).
Comments start with `#`. There are no multi-line comments.

The primary element of a Bitflow Script is a *processing step*.
A processing step is defined through an identifier, followed by parameters in parentheses:

```
avg()
```

An *identifier* may contain letters, numbers and the following characters: `._:\/-`.
Identifiers may not end in a Hyphen. If identifiers that break these rules, they can also be written as *strings*.
Strings are enclosed by pairs of single-quotes, double-quotes or back-ticks: ``'  " ` ``.
There are three different quotes to assure flexibility when embedding different kinds of strings into each other:

```
"avg"()
```

```
'avg'()
```

```
`my "special" avg`()
```

Processing steps may be used with parameters. Parameters are key-value string pairs, separated through `=` und `,` characters:

```
tag(key=value)
```

*Parameter names* are identifiers and follow the same rules as processing step names.
*Parameter values* must be quoted strings or numbers.

```
store_pca_model(precision = 0.99, file = pca-model.bin)
```
```
'store_pca_model'("precision"=`0.99`, `file`="pca-model.bin")
```

Multiple processing steps can be chained together with the `->` operator. This is called a *pipeline*:

```
input-file.csv -> filter(host=mypc) -> avg() -> tag(type=pc)
```

The first step of a pipeline **must** define a *data source*.
A data source can be written as quoted string or an identifier, but without the parameters used for pipeline steps!
Data sinks (or outputs) can be inserted at any position of the pipeline. Data sinks are defined in the same manner as data sources.
If a pipeline contains only one string, it will be interpreted as a data source.
Further details on how to write data sources and sinks are found below, in the *Definition of data sources and sinks* section.
One example of a pipeline with both a data source and multiple processing steps and data sinks:

```
input-file.csv -> filter(host=mypc) -> "intermediate file.bin" -> avg() -> tag(type=pc) -> `tcp+bin://192.168.7.7:8888`
```

Multiple pipelines can be executed in parallel. This is done by separating them through ';' characters and surrounding them by curly braces.

```
{
	input1.csv -> avg();
	input2.csv -> filter(host=mypc) -> 192.168.0.55:5555
}
```

Parallel pipelines can also be used in place of a processing step and arbitrarily nested:

```
input.csv -> { avg() ; max() } -> output.csv
```
```
{
	input.csv -> { avg(); max() } -> 192.168.0.1:5555;
	:6666 -> output.csv
}
```

*NOTE:* in the Java version of Bitflow script, parallel pipelines may only be inserted in the middle of a pipeline, not at the beginning.
This means that there can be only one data source, such as a file or network connection.


Using parallel pipelines as a processing step is  called a *multiplex fork*.
very sample that reaches the opening curly brace will be copied and piped through all defined sub pipelines in parallel.
The step right after the closing curly brace will receive the results of all the sub pipelines.

There can be other *forks* with different functionality, which are written as a pipeline step, followed by parallel pipelines
enclosed in curly braces and separated by semicolons:

```
input.csv -> fork_tags(tag=data_type) { raw -> fft() ; * -> noop() } -> output.csv
```

The available *processing steps*, *forks* and *data sources/sinks* depend on the language implementing the bitflow script,
as well as all names and parameters. See the respective documentations for details.

### Definition of data sources and sinks

A data source or sink (called *endpoint* here) is denoted by a simple string without parameters.
The four main endpoint types are files, TCP listen sockets, active TCP connections, and the standard in/out streams.
The kind of source is detected automatically from the string:
- A *host:port* pair like `192.160.0.55:4444` is interpreted as an active TCP connection, meaning a connection to that endpoint will be established.
  Depending on whether this data store is used as a data source or data sink, the connection will be used to either send or receive data.
- A separate *:port* like `:5555` is interpreted as a TCP listening socket. A socket will be opened, and wait for incoming connections.
  Sockets can also be used to both receive and serve data.
- A single hyphen `-` is interpreted as standard input or output.
- Any other string is interpreted as a file for writing or reading.

If the automatic data store type is not desired, it can be explicitly specified with a URL-like notation:
* `tcp://xxx` will interpret the rest of the string as a TCP endpoint to actively connect to (can also be a single port without hostname, in which case it will connect to localhost)
* `listen://xxx` will interpret the string as a local TCP endpoint to listen on (can include an IP address or hostname for disambiguation)
* `file://xxx` will force read or write to a file
* `empty://xxx` will produce no input or output, the rest of the URL string is ignored (useful in rare cases)
* `std://-` will read or write from standard input or output, only valid with the `-` string

Further endpoint types may be available in the different language implementations, e.g. the `box` endpoint in `Go`, which displays
the data in a continuously updated table on the console.

For most endpoint types, the format of the data is important as well.
The data can be transported and stored in two main formats: CSV (`csv`) and binary (`bin`).
A third data format is the `text` format, which can only be used as output and is helpful to display data in a human-readable way on the console.

The different endpoint types have *default* formats:
* `file` uses `csv`, except if the file ending matches one of the other formats (e.g. `.bin` or `.wav`).
* `bin` is used for `tcp` and `listen`
* `csv` is used for `std`

Other special endpoint types (like `empty`) do not depend on the data format.
*Note:* the different language implementations can define additional formats, such as `wav` in `Java`.
*Note:* in Go, the data format of input data will be determined automatically, while Java currently requires explicit
definition of the expected data format.

The format can also be specified explicitly by specifying it in the same way.
If both a data format and a data store type are specified, they are separated through `+` signs:
```
listen+csv://10.0.0.1:4444
```

### TODO Missing explanations
* Scheduling hints
* Special case of parallel pipelines in the beginning of the top-level pipeline (currently only supported in Go)
* Multi-input definitions
