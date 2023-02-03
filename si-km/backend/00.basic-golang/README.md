# Basic Golang Mentoring

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

## Basic

### Greeting 1

- Concepts: init project, variable, if statement, I/O
- Input: name & gender
- Output: Hi {title} {name}
- Title Rule:
  - If male: Mr
  - If female: Mrs

### Greeting 2

- Concepts: data type (string & int), nested if
- Input: name, gender, and age
- Output: Hi {title} {name}
- Title Rule:
  - If male and >= 17 y.o.: Mr
  - If male and < 17 y.o: Boy
  - If female and >= 17 y.o.: Ms
  - If female and < 17 y.o.: Girl

### Accummulator

- Concepts: I/O read with counter
- Input:
  - `n`: the number of inputs
  - `n` `numbers`
- Output:
  - Sum of the numbers

### Odd Counter

- Concepts: simple loop, modulo, if else, counter variable
- Input:
  - `n`: the number of inputs
  - `n` `numbers`
- Output:
  - The number of odd numbers

## Challenging

### Square Root Inference

- Concepts: loop, math, updating variable
- Input: number `a`
- Output: square root of `a`
- Rule:
  - Loop n times
  - Apply this formula `x = x - ((x*x - a) / (2*x))`
  - Print the final value of `x`
