# appboot

A cross-platform language-agnostic bootstrapper which keeps your application up to date. Available for CLI and GUI applications.

## Manifest file
The manifest file is a JSON file which has to be hosted somewhere. The manifest file holds information about your application such as what files are needed to run, where they can be downloaded from and what their SHA2 hash is. Appboot will use this information to keep your application up to date.

A manifest file can be generated with `appboot manifest`. The JSON file consists of an array of objects with the following keys:

| Manifest key | Description |
| --- | --- |
| `file` | Path to the file. |
| `checksum` | A SHA-256 hash of the contents of the file. |
| `url` | The URL where the file can be downloaded from. |
| `size` | Size of the file in bytes. |

## Launcher configuration
Place a compiled appboot together with a folder named `app`. An `appboot.json` file is required in this app folder with the following contents:

| Configuration name | Description |
| --- | --- |
| `manifest_url` | URL to a manifest containing a list of files and their latest version. |
| `(darwin|windows|linux)_command` | The command to execute to start your program. |

## TODO

- [X] Implement CLI version with [schollz/progressbar](https://github.com/schollz/progressbar)
- [ ] Implement GUI version. [wails](https://github.com/wailsapp/wails)? [fyne](https://github.com/fyne-io/fyne)? [astilectron](https://github.com/asticode/go-astilectron)? [Qt](https://github.com/therecipe/qt)?
- [ ] Is it possible to update appboot itself?
