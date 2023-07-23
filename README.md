# Go EPUB Cover Extractor

Simple program to extract covers from EPUB to use with lf

## Program Overview

The main function in this program is `extractCover(ctx *cli.Context) (bool, error)`. This function takes a CLI context as input and extracts the cover image from the EPUB file specified in the context. If the cover image cannot be extracted for any reason, the function returns an error.

The program uses the `urfave/cli` package to handle command-line arguments and flags.

## Usage

To use this program, run the executable with the path to the EPUB file and the output file as arguments:

```bash
./epub-cover input.epub output.jpg
```

This will extract the cover image from input.epub and save it as output.jpg.

## Building

To build the program, use the go build command:

```bash
go build -o epub-cover
```

This will create an executable named epub-cover.

## Testing

This repository currently does not include any tests. Contributions to add tests are welcome.

## Contributing

Contributions to this repository are welcome. Please feel free to open an issue or submit a pull request.
