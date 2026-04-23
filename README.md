# file-classifier-go

A systemd-triggered file classifier for Linux downloads, written in Go.

## What it does

Watches a directory and automatically moves files into categorised subdirectories — Images, Videos, Music, Documents, and Archives. Triggered by a systemd `.path` unit so it runs instantly when the directory changes. Files still being written to are skipped safely by checking open file descriptors in `/proc`.

## Requirements

- Linux
- Go 1.21+
- systemd

## Installation

Clone the repo and run:

```bash
make install
```

This builds the binary and copies it to `~/.local/bin`, and installs the systemd
unit files to `~/.config/systemd/user` with your home directory substituted in
automatically.

## Enabling the service

After installing, enable and start the path watcher:

```bash
systemctl --user daemon-reload
systemctl --user enable --now file-classifier.path
```

Verify it is running:

```bash
systemctl --user status file-classifier.path
```

## Uninstalling

```bash
make uninstall
```

Then run:

```bash
systemctl --user disable file-classifier.path
systemctl --user daemon-reload
```

## Configuration

The watched directory is set in `systemd/file-classifier.service` and `systemd/file-classifier.path`.
By default it watches `~/Downloads`. To change it, update the placeholder locations
in both unit files before running `make install`, or edit the installed unit files
directly in `~/.config/systemd/user`.

## Categories

| Category  | Extensions                                               |
| --------- | -------------------------------------------------------- |
| Images    | jpg, jpeg, png, gif, webp, tiff, tif, svg, ico, heic     |
| Videos    | mp4, mkv, avi, mov, webm                                 |
| Music     | mp3, flac, wav, aac, ogg, m4a, opus                      |
| Documents | pdf, doc, docx, xls, xlsx, ppt, pptx, txt, md, csv, epub |
| Archives  | zip, tar, gz, 7z, rar, tgz                               |

Unrecognised file types are left in place.
