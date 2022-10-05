# Dasar Pemrograman Backend Mentoring Week-1

## Basic Rules
- **This is not a demo!**
- ### **The challenges are for reference, can be used as is, customized or completely use your own challenges. Keep in mind the Learning Objectives!**
- Mentor will be the facilitator, guiding the mentee to solve the tasks.
  - Share screen method
    - One person will share their screen, can be the mentor and one of the mentee.
    - Mentee will take turn in giving the correct syntax to be typed by person sharing his/her screen.
  - Using tools such as replit https://replit.com/
    - Everyone can login using their own account.
    - Mentor can have a tab for each mentee to monitor their work.
    - Details over at Discord: https://discord.com/channels/996138478817513632/996141512704401579/1019923694254043136
  
## Learning Objectives

- Able to init Golang project
- Able to run main.go
- Understand what variables are for
- Able to do basic I/O
- Able to use if statement
- Looping
  - Understand the logic of looping
  - Able to write basic for loop
  - Able to write nested for loop
- Function
  - Able to write a function
  - Able to call a function

## Pre-requisites

- Code editor (e.g. VS Code).
- Golang installed in local computer

## Basic

### Greeting 1

- Concepts:
  - init project
  - variable
  - if statement
  - I/O
- Input:
  - name & gender
- Output:
  - Hi {title} {name}
- Directions:
  - Create empty go file
  - Create function main()
  - Use fmt.Print* for output
  - Use fmt.Scan* for input
  - Title Rule:
    - If male: Mr
    - If female: Mrs

### Greeting 2

- Concepts:
  - data type (string & int)
  - nested if
- Input:
  - name, gender, and age
- Output:
  - Hi {title} {name}
- Direction:
  - Create file .go and main()
  - Use fmt for input and output
  - Title Rule:
    - If male and >= 17 y.o.: Mr
    - If male and < 17 y.o: Boy
    - If female and >= 17 y.o.: Ms
    - If female and < 17 y.o.: Girl

### Accummulator

- Concepts:
  - I/O read with counter
- Input:
  - `n`: the number of inputs
  - `n` `numbers`
- Output:
  - Sum of the numbers
- Directions:
  - Setup Golang file and function
  - Use fmt for input and output
  - Will receive multiple inputs from the terminal, the first the the number of inputs, the second is the actual numbers separated by comma
  - Calculate the numbers by adding together
  - Output the total sum to terminal

### Odd Counter

- Concepts:
  - simple loop
  - modulo
  - if else
  - counter variable
- Input:
  - `n`: the number of inputs
  - `n` `numbers`
- Output:
  - The number of odd numbers
- Directions:
  - Setup Golang file and function
  - Use fmt for I/O
  - Will receive multiple inputs from the terminal, the first is the number of inputs, followed by a series of numbers
  - Find the odd number
  - Output odd numbers to terminal

### With Function
- Concepts: functions
- Input:
  - Using any of the above challenges
- Output:
  - Convert to function

## Challenging

### Square Root Inference

- Concepts:
  - loop
  - math
  - updating variable
- Input:
  - number `a`
- Output:
  - square root of `a`
- Directions:
  - Setup Go file and function
  - Use fmt for I/O
  - Loop n times
  - Apply this formula `x = x - ((x*x - a) / (2*x))`
  - Print the final value of `x`
  - See subfolder square-root for answer
