# KMA Score Scanner

KMA Score Scanner is a command-line tool written in Go that scans scores and students' data from HTML and exports the
data to various files type.

This tool is a upgraded version of [KMA Score Extractor](https://github.com/KMA-Score/KMA-Score-Extractor)

## Features

- Scans scores and students' data from HTML.
- Exports the scanned data to:
    - TSV file.
    - JSON file.
    - SQL file.

## Usage

Get the latest version of the tool from the [release page](https://github.com/KMA-Score/kma_score_scanner/releases) or
clone and build the tool yourself.

### Step 1: Convert PDF to HTML

You need a account from Aspose Cloud or [Self-hosted Aspose Total for Cloud](https://purchase.aspose.cloud/self-hosting/) to convert PDF to HTML. You can get a free
account from [here](https://dashboard.aspose.cloud/#/apps).
Then you need to create a file named `config.json` by copying the content of `config.example.json` and fill in the
required information.

After that, you can use the `convert` command followed by the input path. The output directory path for the HTML files
can be specified using the `-o` flag.

```bash
kma_score_scanner tools pdf2html ./input -o ./output
```

Input can be a file or a directory. If input is a directory, all PDF files in that directory will be converted to HTML.

### Step 2: Scan scores and students' data

The tool can be used with the `scan` command followed by the input path. The output directory path for the TSV file can
be specified using the `-o` flag.

Example:

```bash
kma_score_scanner scan ./input -o ./output
```

In the above command, `./input` is the directory containing the HTML files to be scanned and `./output` is the directory
where the TSV files will be saved.

### Step 3: Export data to SQL command

The tool can be used with the `export` command followed by the input path. The output directory path for the JSON file
can be specified using the `-o` flag.

```bash
kma_score_scanner tools tsv2sql ./input -o ./output
```

