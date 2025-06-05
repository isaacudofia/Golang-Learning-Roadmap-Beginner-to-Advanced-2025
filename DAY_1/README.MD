# Day 1: Getting Started with Go (Golang)

---

Welcome to Day 1 of your Go learning journey! Today, you'll set up your development environment and write your very first Go program. This foundational step ensures you're ready to dive deeper into the language.

## 🚀 Today's Goals

- Understand what Go is and why it's a powerful language.
- Successfully install Go on your machine.
- Write and execute your first "Hello, Go!" program.

## 🛠️ Setup Your Go Development Environment

### 1. Install Go

The quickest way to get Go up and running is by using the official installers.

- **Download:** Visit the official Go download page: [https://go.dev/dl/](https://go.dev/dl/)
- **Select your OS:** Choose the appropriate installer for your operating system (Windows, macOS, Linux).
- **Run Installer:** Follow the on-screen instructions. The installer usually handles setting up environment variables for you.

### 2. Verify Installation

Open your terminal or command prompt and run:

```bash
go version
```

You should see output similar to `go version go1.22.4 linux/amd64` (the version number may vary). If you do, Go is installed correctly!

### 3. Choose a Code Editor/IDE

While any text editor works, a dedicated code editor or IDE will make your Go development much smoother.

- **VS Code (Recommended for beginners):**
  - Download: [https://code.visualstudio.com/](https://code.visualstudio.com/)
  - Install the official **"Go" extension** by Google from the Extensions Marketplace within VS Code.
- **GoLand (Paid, powerful IDE):**
  - Download: [https://www.jetbrains.com/go/](https://www.jetbrains.com/go/)

## 💻 Project for Today: "Hello, Go!"

The classic first program! This project confirms your setup is working and introduces you to the basic structure of a Go application.

### Instructions:

1.  **Create a project directory:**
    Open your terminal and navigate to where you want to store your Go projects (e.g., `~/go/src/`). Then, create a new directory for today's project:

    ```bash
    mkdir hello_world
    cd hello_world
    ```

2.  **Create `main.go`:**
    Inside your `hello_world` directory, create a new file named `main.go`.

3.  **Write the code:**
    Open `main.go` in your chosen code editor and paste the following:

    ```go
    package main

    import "fmt"

    func main() {
        fmt.Println("Hello, Go!")
    }
    ```

    - `package main`: Every executable Go program starts with the `main` package.
    - `import "fmt"`: This imports the "fmt" (format) package, which provides functions for input/output.
    - `func main()`: This is the entry point of your program; execution begins here.
    - `fmt.Println("Hello, Go!")`: This line calls the `Println` function from the `fmt` package to print the string "Hello, Go!" to your console.

4.  **Run your program:**
    Save `main.go`. In your terminal, ensure you are in the `hello_world` directory, then run:

    ```bash
    go run main.go
    ```

    You should see:

    ```
    Hello, Go!
    ```

5.  **Build an executable (Optional):**
    You can compile your Go code into a standalone executable:

    ```bash
    go build main.go
    ```

    This creates an executable file (`main` or `main.exe`) in your current directory, which you can run directly.

---

## ✅ Day 1 Complete!

You've successfully taken the first step into the world of Go! You now have a working environment and understand the very basics of Go program structure.

Feel free to experiment with changing the text in `fmt.Println` to something else!

---

**Next Up: Day 2 will cover Variables, Constants, and Basic Data Types.**
