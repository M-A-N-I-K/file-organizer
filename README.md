## Organize CLI

Organize is a lightweight Go-based CLI tool that organizes files in a directory based on their extensions.

The tool scans a folder and moves files into subdirectories named after their file extensions, helping keep the file system clean and structured.

```bash
$ organize <path>
```

---

## Features

* Organize files by extension
* Create folders automatically
* Move files into extension-based directories
* Preview changes using `--dry-run`
* Show move logs in table format using `--log`
* Simple CLI interface
* Fast and lightweight
* Built using Go

---

## How it works

If a folder contains:

```
/downloads
  file1.txt
  file2.jpg
  file3.png
  notes.txt
```

After running:

```bash
$ organize ./downloads
```

Result:

```
/downloads
  /txts
    file1.txt
    notes.txt
  /jpgs
    file2.jpg
  /pngs
    file3.png
```

---

## Usage

```bash
$ organize <folder-path>
```

Example:

```bash
$ organize .
$ organize ./downloads
$ organize /Users/manik/Desktop/files
$ organize ./downloads --dry-run
$ organize ./downloads --log
```

### Dry run mode

Use `--dry-run` to preview file moves without changing anything on disk.

```bash
$ organize ./downloads --dry-run
[DRY RUN] The following changes will be made:
ACTION   SOURCE                      DESTINATION
MOVE     file1.txt                   txt/file1.txt
MOVE     image.jpg                   jpg/image.jpg
MOVE     data.csv                    csv/data.csv
```

### Logging mode

Use `--log` to print move operations in the same table UI.

```bash
$ organize ./downloads --log
ACTION   SOURCE                      DESTINATION
MOVE     file1.txt                   txt/file1.txt
MOVE     image.jpg                   jpg/image.jpg
MOVE     data.csv                    csv/data.csv
```

---

## Installation

Build the CLI:

```bash
go build -o organize
```

Run:

```bash
./organize <path>
```

---

## Tech Stack

* Go (Golang)
* os package for file handling
* filepath package for path operations
* Cobra for CLI arguments

---

## Current Limitations

* Does not handle duplicate file names
* Does not support ignore rules

---

## TODO (Future Features)

Planned improvements for this project:

* [ ] Ignore specific folders (node_modules, .git, etc.)
* [ ] Custom rules (group images, videos, docs, etc.)
* [ ] Config file support (JSON / YAML)
* [ ] Colored CLI output
* [ ] Progress indicator for large folders
* [ ] Move files based on size / date / type

---

## Conclusion

Organize CLI is a simple Go project built to practice file handling, CLI development, and system utilities.
It can be extended with many advanced features and is a good beginner-to-intermediate Go project.

