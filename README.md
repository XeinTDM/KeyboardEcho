# KeyboardEcho

KeyboardEcho is a prank program that randomly echoes the last key pressed after a few repetitions. It's harmless and lightweight.

## Features
- Randomly echoes keyboard input for added mischief.
- Compact executable size: only 1484 KB with UPX and Go build flags.

## Usage
1. Clone the repository:
   ```bash
   git clone https://github.com/XeinTDM/KeyboardEcho.git
   cd KeyboardEcho
   ```

2. Install dependencies:
   ```bash
   go get github.com/go-vgo/robotgo
   go get github.com/robotn/gohook
   ```

3. Run the program:
   ```bash
   go run main.go
   ```

## Configuration
- **MinThreshold**: The minimum random threshold for key echoing (default: `10`).
- **MaxThreshold**: The maximum random threshold for key echoing (default: `50`).

## License
This project is licensed under [The Unlicense](LICENSE), meaning it is free and open for public use.
