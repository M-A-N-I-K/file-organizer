## Organize CLI

### Goal

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

* Not recursive (only organizes current folder)
* Does not handle duplicate file names
* Does not support ignore rules
* Does not support dry-run mode

---

## TODO (Future Features)

Planned improvements for this project:

* [ ] Recursive folder organization
* [ ] Ignore specific folders (node_modules, .git, etc.)
* [ ] Dry run mode (preview changes)
* [ ] Custom rules (group images, videos, docs, etc.)
* [ ] Config file support (JSON / YAML)
* [ ] Logging support
* [ ] Colored CLI output
* [ ] Progress indicator for large folders
* [ ] Move files based on size / date / type

---

## Conclusion

Organize CLI is a simple Go project built to practice file handling, CLI development, and system utilities.
It can be extended with many advanced features and is a good beginner-to-intermediate Go project.

