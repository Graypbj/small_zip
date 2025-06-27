# SmallZip
A lightweight CLI tool for compressing and decompressing files and folders using `.zip` and `.gz` formats, optimized for **low-memory environments** like a Raspberry Pi.

## Features

- 📁 Compress folders into `.zip` archives
- 📂 Extract `.zip` archives
- 🗜️ Compress single files using `.gz`
- 🔄 Decompress `.gz` files
- 📉 Low memory usage with **user-defined buffer size**
- 💡 Default buffer size: 32 KB

## 🚀 Getting Started

### Build
```bash
go build -o smallzip
```

### Usage
```bash
./smallzip -mode=<zip|unzip|gzip|gunzip> -src=<source_path> -dest=<destination_path> [-bufsize=<bytes>]
```

### Examples

#### Zip a folder (default 32KB buffer)

```bash
./smallzip -mode=zip -src=./myfolder -dest=archive.zip
```


#### Unzip with 64KB buffer
```bash
./smallzip -mode=unzip -src=archive.zip -dest=./extracted -bufsize=65536
```

#### Gzip a file

```bash
./smallzip -mode=gzip -src=log.txt -dest=log.txt.gz
```

#### Gunzip a file with small buffer

```bash
./smallzip -mode=gunzip -src=log.txt.gz -dest=log.txt -bufsize=16384
```

---


## Command Line Flags

| Flag       | Description                                             | Default        |
|------------|---------------------------------------------------------|----------------|
| `-mode`    | Operation to perform: `zip`, `unzip`, `gzip`, `gunzip` | _(required)_   |
| `-src`     | Source file or folder                                   | _(required)_   |
| `-dest`    | Output file or directory                                | _(required)_   |
| `-bufsize` | Max memory buffer in bytes (e.g. `32768`)               | `32768` (32KB) |

---

## How It Works

This utility:
- Uses `os.Open()` and `io.CopyBuffer()` to stream files.
- Only loads data into memory based on the buffer size (`[]byte` slice).
- Avoids full-file reads (`ioutil.ReadAll`, etc.).
- Uses Go's built-in `archive/zip` and `compress/gzip` packages.

> Efficient even for multi-gigabyte files on embedded systems.

---

## Designed for

- Raspberry Pi
- SBCs (single board computers)
- Containers with limited memory
- Portable USB tools

---

## License
GNU V3.0

## Author
Grayson Butcher
_Engineer • Developer • Maker_

---

## Feedback

Pull requests and feature suggestions welcome!

