# ⚖️ Body Mass Index (BMI) Calculator

<p align="center">
  <img src="https://img.shields.io/badge/Language-Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go">
  <img src="https://img.shields.io/badge/Type-CLI_Tool-181717?style=for-the-badge&logo=windows-terminal&logoColor=white" alt="CLI Tool">
  <img src="https://img.shields.io/badge/Status-Completed-2EA44F?style=for-the-badge" alt="Project Status">
</p>

<p align="center">
  <b>A simple and lightweight BMI calculator written in Go.</b>
  <br>
  Calculate your Body Mass Index directly from your terminal.
</p>

---

## 📌 About

**Body Mass Index (BMI) Calculator** is a command-line application written in **Go (Golang)** that calculates Body Mass Index based on a user's height and weight.

The program processes user input, calculates the BMI value, and displays the corresponding category in the terminal.

Designed with simplicity and readability in mind, this project is a small example of using Go for interactive console applications.

## ✨ Features

- ⚖️ **BMI Calculation** — calculates BMI using height and weight.
- 📊 **Result Classification** — displays a category based on the calculated value.
- ⌨️ **Interactive Input** — prompts the user to enter their height and weight.
- 🛡️ **Input Validation** — rejects non-positive values.
- 🖥️ **Terminal Interface** — lightweight and easy to run.
- 🧩 **Pure Go** — uses Go's standard library without external dependencies.

## 🧮 Formula

The calculator uses the standard BMI formula:

\[
\text{BMI} = \frac{\text{Weight (kg)}}{\text{Height (m)}^2}
\]

The height is entered in centimeters and converted to meters during the calculation.

## 🛠️ Built With

- **[Go](https://go.dev/)** — programming language.
- **fmt** — user input and terminal output.
- **math** — mathematical calculations.

## 🚀 Getting Started

### Prerequisites

Make sure [Go](https://go.dev/dl/) is installed on your computer.

### Installation

**1. Clone the repository**

```bash
git clone https://github.com/w94e/Body-Mass-Index-BMI-.git
```

**2. Open the project directory**

```bash
cd Body-Mass-Index-BMI-
```

**3. Run the application**

```bash
go run IMT.go
```

## 💻 Example

```text
《 Калькулятор индекса массы тела (ИМТ) 》

✦› Введите ваш рост (в сантиметрах): 175
✦› Введите ваш вес (в килограммах): 70

✦› Ваш индекс массы тела (ИМТ): 22.9
✦› Категория: Нормальный вес
```

*Example output is shown in Russian, matching the current program interface.*

## 📂 Project Structure

```text
Body-Mass-Index-BMI-/
└── IMT.go    # Main application source code
```

## 🎯 Project Goals

This project demonstrates the practical use of Go fundamentals, including:

- Variables and data types
- Functions and return values
- User input and terminal output
- Conditional statements (`switch`)
- Mathematical operations
- Basic input validation

## ⚠️ Disclaimer

> This calculator is intended for educational and informational purposes only. BMI is a general screening measure and does not provide a complete assessment of an individual's health.

## ⚠️ Project Disclaimer

> This project is **strictly educational** and was created with limited knowledge of the Go programming language.

> The code contains significant bugs, design flaws, and architectural shortcomings. It was developed as an experiment to explore Go's complexity, capabilities, and potential for further learning and development.

> **This project is not intended for production use.** The code is provided for educational and experimental purposes only and should not be considered production-ready.

> The project may be revisited, refactored, and improved as my understanding of Go evolves.

> *Learning by building, experimenting, and making mistakes.*

## 👨‍💻 Author

**[w94e](https://github.com/w94e)**

- GitHub: [@w94e](https://github.com/w94e)
- Repository: [Body-Mass-Index-BMI-](https://github.com/w94e/Body-Mass-Index-BMI-)

---

<p align="center">
  ⭐ If you find this project interesting, consider giving it a star!
  <br><br>
  <i>Built with Go 💙</i>
</p>
