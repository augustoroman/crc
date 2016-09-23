# CRC

CRC is a simple command-line utility to compute various CRC values for one or
more files.

It supports three CRC algorithms (selected via `--mode=<alg>`):
* `crc64-ecma` (*default*) uses the [ECMA polynomial](https://godoc.org/hash/crc64#pkg-constants) to compute a 64-bit CRC.
* `crc64-iso` uses the [ISO polynomial](https://godoc.org/hash/crc64#pkg-constants) to compute a 64-bit CRC.
* `crc32` uses the [IEEE polynomial](https://godoc.org/hash/crc32#pkg-constants) to compute a 32-bit CRC.  This is equivalent to the `crc32` linux binary.

If only a single file is specified on the command line, then only the hexadecimal hash is printed.  If more than one file is specified, the output will be
```
<hash><tab><filename>
```
for each file.

## Installation

```bash
go get github.com/augustoroman/crc
```

## Usage

```bash
# Compute the default hash (crc64 w/ ECMA polynomial) of one or more files.
crc <file> [<file> ...]
# Compute the crc64 hash with ISO polynomial.
crc --mode=crc64-iso <file> [<file> ...]
# Compute the crc32 hash (with the IEEE polynomial).
crc --mode=crc32 <file> [<file> ...]
```

## License
This is released under the MIT license
