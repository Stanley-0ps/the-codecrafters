## 🧰 String Transformer (Go)

A simple command-line text transformer built with Go. It allows users to manipulate strings in multiple ways using an interactive, command-based interface. Ideal for quick text formatting and fun text effects.

---

## ✨ Features

✍️ Convert text to **UPPERCASE** (`upper`)
🔡 Convert text to **lowercase** (`lower`)
📜 Convert text to **Title Case** (`title`)
🔠 Capitalize the **first letter of each word** (`cap`)
🐍 Convert text to **snake_case** (`snake`)
🔄 **Reverse letters** in each word (`reverse`)
💻 Interactive CLI with command-style input (`<command> <text>`)
🚫 Handles empty input gracefully (pressing Enter shows prompt again)

---

## 🚀 Getting Started

### Prerequisites

* Install Go (version 1.18 or later recommended)

### Run the program

```bash
go run main.go
```

---

## 📌 Usage

Enter commands in the format:

```bash
<command> <text>
```

### Examples:

```bash
> upper hello world
result: HELLO WORLD

> cap go is awesome
result: Go Is Awesome

> snake Hello World!
result: hello_world

> reverse Hello World!
result: olleH dlroW
```

To exit:

```bash
> exit
```

---

## ⚠️ Error Handling

This program is designed to handle common user mistakes:

**Unknown command**

```bash
> foobar test
✗ Unknown command: "foobar"
Valid commands: upper, lower, cap, title, snake, reverse, exit
```

**Missing text**

```bash
> upper
✗ No text provided. Example Usage: upper <text>
```

**Empty input**

```bash
> 
# Just shows the prompt again
```

---

## 🧠 Project Structure

```
string-transformer/
├── main.go        # Entry point and CLI logic
├── README.md      # Project documentation
```

---

## 📄 License

This project is open-source and free to use.

---
