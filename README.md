
# 2048 TUI

Want a Brazilian portuguese readme?[click here!](README_pt.md)

A terminal-based implementation of the popular 2048 game, written in Go. This version features a colorful interface and sound effects for an enhanced gaming experience.

## Features

- Play 2048 directly from your terminal.
- Colorful terminal interface using the Bubble Tea framework.
- Sound effects for moves and game over (when compiled with audio support).
- Cross-platform compatibility (macOS, Linux, Windows).
- Responsive layout that adjusts to terminal size.
- Simple and intuitive TUI for an enjoyable experience.
- Written in Go for high performance and portability.

## Prerequisites

To build and run this game, you need to have Go installed on your system. The project uses Go modules for dependency management.

## Installation

To get started with `2048-tui`, clone the repository and install the necessary dependencies:

```bash
git clone https://github.com/lsferreira42/2048-tui
cd 2048-tui
```

Then, run the provided script to install dependencies:

```bash
./deps.sh
```

This script will install [go-bindata](https://github.com/go-bindata/go-bindata) to manage assets in your project.

## Building

To build the game with audio support, run:

```bash
make build
```

This will create a binary in the `build` directory.

## Running the Game

After building, you can run the game using:

```bash
./build/2048
```

Alternatively, to run the game directly without building:

```bash
make run
```

Or you can use:

```bash
go run -tags=audio .
```

## Playing the Game

- Use arrow keys to move the tiles.
- Combine tiles of the same number to create higher numbers.
- Press 'r' to restart the game.
- Press 'q' to quit.
- Try to reach the 2048 tile!

## Distribution

To create binaries for multiple platforms:

```bash
make dist
```

This will create binaries for macOS (x86_64 and ARM), Linux (x86, ARM, and x64), and Windows (64-bit) in the `dist` directory.

## Project Structure

- `2048.go`: Main game logic and UI.
- `audio.go`: Audio playback functionality.
- `deps.sh`: Script to install additional dependencies.
- `Makefile`: Build and run commands.
- `move_sound.mp3`: Sound effect for moves.
- `ending.mp3`: Sound effect for game over.

## Dependencies

- **Go**: The latest version is recommended. Go is required to compile and run the project. You can install it by following the instructions on the [official website](https://golang.org/doc/install).
- **go-bindata**: This tool is used to embed data into Go applications, making it easier to manage static assets. You can install it by running the following command:

  ```bash
  go install -a -v github.com/go-bindata/go-bindata/...@latest
  ```

- **GNU Make**: Used to build the project with the `Makefile`. Ensure that GNU Make is installed on your system. On most Linux systems, you can install it with:

  ```bash
  sudo apt-get install make
  ```

  On macOS, you can use:

  ```bash
  brew install make
  ```

- **Bubble Tea**: Terminal UI framework used for building the colorful interface.
- **Lip Gloss**: Style definitions for terminal UI.

## License

This project is open-source and available under the [MIT License](LICENSE).

## Contributing

Feel free to submit issues, suggestions, or pull requests to help improve the project. Contributions are always welcome!

## Contact

Maintainer: [Leandro Ferreira](https://github.com/lsferreira42)

Repository URL: [2048 TUI on GitHub](https://github.com/lsferreira42/2048-tui)

## Acknowledgments

- Inspired by the original 2048 game by Gabriele Cirulli.
- Thanks to the Charm team for their excellent terminal UI libraries.
