# appboot

A cross-platform language-agnostic bootstrapper which keeps your application up to date. Available for CLI and GUI applications.

## TODO

- [ ] Implement CLI version with [schollz/progressbar](https://github.com/schollz/progressbar)
- [ ] Implement GUI version. [wails](https://github.com/wailsapp/wails)? [fyne](https://github.com/fyne-io/fyne)? [astilectron](https://github.com/asticode/go-astilectron)? [Qt](https://github.com/therecipe/qt)?
- [ ] Is it possible to update appboot itself?

## Installation

## Application structure

```
<Program name>.(exe|sh)
app/appboot.json
```

## Manifest
A manifest can be generated with `appboot manifest` and is always named `appboot.json`. A manifest is a JSON file containing an array of objects with the following keys:

| Manifest key | Description |
| --- | --- |
| `file` | Path to the file. |
| `checksum` | A SHA-256 hash of the contents of the file. |
| `url` | The URL where the file can be downloaded from. |
| `size` | Size of the file in bytes. |

## Launcher configuration

| Configuration name | Description |
| --- | --- |
| `manifest_url` | URL to a manifest containing a list of files and their latest version. |
| `command` | The command to execute to start your program. |

## Lifecycle

1. Read configuration from env/files with [viper](https://github.com/spf13/viper)
2. Create launcher
3. Download remote master file
4. Determine if upate is needed, if none is needed continue to 6
5. Track download progress
6. Execute command