# appboot

A cross-platform language-agnostic bootstrapper which keeps your application up to date. Available for CLI and GUI applications.

## TODO

- [ ] Implement CLI version with [schollz/progressbar](https://github.com/schollz/progressbar)
- [ ] Implement GUI version. [wails](https://github.com/wailsapp/wails)? [fyne](https://github.com/fyne-io/fyne)? [astilectron](https://github.com/asticode/go-astilectron)? [Qt](https://github.com/therecipe/qt)?
- [ ] Is it possible to update appboot itself?

## Application structure

```
<Program name>.(exe|sh)
appboot.yaml
app/
```

## Manifest
A manifest can be created with `appboot packer`. A manifest is a YAML file containing an array of:

| Manifest key | Description |
| --- | --- |
| `file` | Path to the file from `app/`. |
| `checksum` | A SHA-256 hash of the contents of the file. |
| `url` | The URL where the file can be downloaded from. |
| `size` | Size of the file in bytes. |

## Launcher configuration

| Configuration name | Description |
| --- | --- |
| `manifest_url` | URL to a manifest containing a list of files and their latest version. |
| `command_windows` | The command to execute to start your program on windows. |
| `command_linux` | The command to execute to start your program on linux. |
| `command_macos` | The command to execute to start your program on macOS. |

## Lifecycle

1. Read configuration from env/files with [viper](https://github.com/spf13/viper)
2. Create launcher
3. Download remote master file
4. Determine if upate is needed, if none is needed continue to 7
5. Track download progress
6. Track unpacking progress
7. Execute command