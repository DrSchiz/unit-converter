### Unit Converter

This application is a go web project that allows the user to convert various quantities based on their units of measurement. Idea: https://roadmap.sh/projects/unit-converter.

### Installation

1.  Install Go from the official website: https://go.dev/doc/install.

2.  Create a local directory and clone the project code into it:
    ```bash
    git clone https://github.com/DrSchiz/unit-converter
    ```

3.  Navigate to the project directory:
    ```bash
    cd unit-converter
    ```

4.  Build the application using the Go command:
    ```bash
    go build -o unit-converter cmd/main.go
    ```

This will create an executable binary file named `unit-converter` in your current directory.

### Launching/Running

After you launch your executable binary file, the terminal will show you:

```bash
Server running on :8000
```

So, you need to open link in your browser: `http://localhost:8000/`.

After that you will see the unit measurements (`Length`, `Weight`, `Temperature`), click on them to start convert you quantities. 

There is three areas for enter your values: first is the value of unit measurement, second is the unit of measurement to convert from, and third is the unit of measurement to convert to.