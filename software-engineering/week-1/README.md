# Software Engineering Mentoring Week-1

## Learning Objectives

- Understand daily activities as a software engineer.
  - Able to formulate algorithm to solve a problem.
  - Able to create pseudocode for simple application.
  - Understand how networking works in the context of web application
    - Able to identify parts of a URI
    - Able to access website by port

- Able to use tools commonly used by software engineer.
  - Able to use Unix based terminal
    - Know the directory he/she is working on
    - Can list files and foler inside a directory
    - Can move between directories
    - Can create a new file
    - Can delete file
    - Can open file with code editor
  - Able to use code editor
    - Can open file with code editor
    - Can work with code editor, add/edit/remove text from a file opened with code editor
    - Can save file with code editor
  - Able to use basic git commands
      - git config
    - git init
    - git status
    - git add
    - git commit
    - git branch
  - Able to connect with Github and work with Github repo
    - git clone
    - git checkout
    - git push

## Pre-requisities

- Install code editor (e.g. VS Code)
- Install gitBash for Windows or git for Linux based distro (in WSL for windows)
- Create account in github

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
