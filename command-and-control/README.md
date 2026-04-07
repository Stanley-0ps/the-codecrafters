# 🛰️ SENTINEL — Command & Control Console (Go)

A unified command-line console built with Go, combining a **calculator**, **base converter**, and **string transformer**. Perform calculations, convert numbers, and transform text all from **one interactive CLI**.

---

## ✨ Features

### 🧮 Calculator (`calc`)

* `add`, `sub`, `mul`, `div`, `mod`, `pow` operations
* `last` keyword tracks most recent numeric result
* Handles invalid input, missing arguments, and division/mod by zero gracefully

### 🔢 Base Converter (`base`)

* Convert decimal → binary & hex (`dec`)
* Convert hex → decimal (`hex`)
* Convert binary → decimal (`bin`)
* Validates input and rejects invalid numbers with clear feedback

### 🔤 String Transformer (`str`)

* `upper`, `lower`, `cap`, `title`, `snake`, `reverse`
* Palindrome detection and character/word counting
* Ignores extra spaces, preserves numbers & symbols, case-insensitive commands

### 🔗 Integration

* Pipe outputs between modules using `|`
* Shared `last` result across modules
* Unified session history (last 10 commands)
* `clear` resets history and last value with confirmation

---

## 🚀 Getting Started

### Prerequisites

* Install [Go](https://go.dev/) (version 1.18 or later recommended)

### Run the Program

```bash
go run main.go
```

### Usage

From the `C&C>` prompt:

```text
C&C> calc add 42 58
✦ Result: 100

C&C> base dec 255
✦ Binary: 11111111
✦ Hex: FF

C&C> str upper sentinel is watching
✦ SENTINEL IS WATCHING

C&C> calc add 200 55 | base dec
✦ Result: 255
✦ Binary: 11111111
✦ Hex: FF

C&C> history
1. calc add 42 58
2. base dec 255
3. str upper sentinel is watching
4. calc add 200 55 | base dec

C&C> clear
Type CONFIRM to clear the session: CONFIRM
✔ Session cleared.

C&C> exit
```

---

## ⚠️ Error Handling

* Invalid commands or missing arguments show a **clear message**
* Division/modulus by zero handled without crashing
* Invalid numbers in base conversion rejected
* Empty inputs ignored
* Case-insensitive commands for string module

---

## 🧠 Project Structure

```text
command-and-control/
├── main.go                  # Entry point & console loop (C&C prompt)
├── calculator.go           # Calculator module (add, sub, mul, div, mod, pow)
├── base-converter.go       # Base conversion logic (dec, bin, hex)
├── string-transformer.go   # String operations (upper, lower, snake, etc.)
├── pipe.go                 # Pipe system (command chaining with |)
├── helpers.go              # Shared utilities (parsing, formatting, last result)
├── help.go                 # Help menu and command descriptions
├── README.md               # Project documentation
```

---

## 📖 Example Function

```go
func Addition(a, b float64) float64 {
    return a + b
}

func Convert(value string, base string) {
    switch strings.ToLower(base) {
    case "hex":
        val, _ := strconv.ParseInt(value, 16, 64)
        fmt.Println("✦ Decimal:", val)
    }
}
```

---

## 🎯 Future Improvements

* Support more numeric bases (octal, base-36)
* Extend string module with more transformations
* Add CLI flags (`--from`, `--to`)
* Unit tests for all modules
* Build as a global CLI tool

---

## 📄 License

Open-source and free to use.

---
