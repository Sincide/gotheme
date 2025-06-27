# GoTheme

GoTheme is a prototype theming manager for Hyprland on Arch Linux written entirely in Go. It provides a GUI to select a wallpaper, extract a color palette, and query a local Ollama LLM for theme suggestions.

## Features

- **GUI Wallpaper Selection** using Fyne
- **Color Palette Extraction** with a pure Go algorithm
- **Wallpaper Setting** via `swww`
- **LLM Integration** by calling a local Ollama instance

## Build

```bash
go build
```

## Usage

Run the compiled binary and choose `File -> Select Wallpaper` to pick an image. GoTheme will set the wallpaper, extract its palette, display swatches, and send the colors to `http://localhost:11434/api/generate`.

## Extending

Future work includes generating configuration files from templates and managing multiple theme profiles.
