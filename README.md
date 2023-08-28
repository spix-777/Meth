# Mouse Mover with Termbox Display

This project provides a simple demonstration of utilizing the `termbox-go` library to display a string message in the center of a terminal while simultaneously using the `robotgo` library to move the mouse cursor by 1 pixel on the X-axis every second.

## Features

1. Displays a string message at the center of the terminal.
2. Moves the mouse cursor incrementally every second.
3. Exits the program when a key is pressed.

## Prerequisites

Ensure you have the following dependencies installed:

- [termbox-go](https://github.com/nsf/termbox-go): A library for creating cross-platform text-based interfaces.
  
  ```bash
  go get -u github.com/nsf/termbox-go
  ```

- [robotgo](https://github.com/go-vgo/robotgo): A cross-platform library for GUI automation.

  ```bash
  go get -u github.com/go-vgo/robotgo
  ```

## Usage

1. Clone the repository.
2. Navigate to the project directory.
3. Run the program:

```bash
go run main.go
```

Once executed, you'll see the message "YOu hIght on MeTHamPhEtamInEs" displayed at the center of your terminal. Simultaneously, the mouse cursor will move by 1 pixel on the X-axis every second.

Press any key to exit the program.

## Contribution

Feel free to fork the project, submit PRs, and report issues or suggestions. We always welcome contributions from the community!

## License

This project is open-sourced under the MIT License.