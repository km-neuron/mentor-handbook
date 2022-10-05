# Programming with Javascript Mentoring Week-1

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

- Understand how to use Primitive Data Type
    - Able to declare a variable
    - Able to assign value to variable
    - Able to get value from variable
- Understand how to use Conditional
    - Able to create a simple conditional using if else
    - Able to created nested conditional using if else if
    - Able to use switch case and break effectively
- Understand how to use Looping
    - Able to create a simple loop using for
    - Able to create a simple loop using while
    - Able to create nested loop using for
- Understand how to use Function
    - Able to create a function
    - Able to send parameter to a function and use the parameter inside the function
    - Able to return a value out from a function and use the returned value outside the function.
    - Able to create arrow function
    - Able to use modular function by using function inside another function

## Pre-requisites

- Code editor (e.g. VS Code).
- Web based playground
    - Web Browser with Developer Tools.
        - Bisa dengan membuat index.html berisi DOM yang akan refresh ketika kita mengetik kode JS (perlu effort dari mentor)
        - Bisa dengan extension LiveShare atau lainnya.
        - Atau metode lainnya.
    - Online playground such as jsbin.com

## Basic

### Simple Calculation
- Concepts: 
  - Variable declaration, assignment and retrieval
  - Simple conditional using variables
- Input: 
  - Two numbers, one arithmetic operator ( +, -, *, /)
- Output:
  - Result of a simple arithmetic calculation
- Directions: 
  - Create 3 variables
    - First variable is a number.
    - Second variable is a number.
    - Third variable is an arithmetic operator.
    - Use console.log to show the variables.
  - Create conditional
    - Still using only 3 variables, I know it's not optimal but we are teaching not coding.
    - First conditional if operator is +, will add the two numbers and output the result.
    - Second conditional is -, will subtract and output the result.
    - Third conditional is *, will multiply and output the result.
    - Fourth conditional is /, will divide and output the result.
    - Fifth conditional is %, will get remainder of a division and output the result.
    - It's preferred to store the result into a variable before using console.log (optional)
    - Use console.log to show the result.
  - Use switch case
    - Using the conditional if...else above, convert to switch case
    - Add the operator "add" and show that we can omit break
- Rule
  - N/A

### Simple Calculation with Function
- Concepts: 
  - Create and use functions
  - Use conditionals with functions
- Input: 
  - Same three variables as Simple Calculation
- Output: 
  - Result of simple arithmetic calculation using function
- Directions:
  - For each operation ( add, subtract, multiply, divide, modulus ), create a function.
  - Show what happens if we don't add the parameter.
  - Show what happens if we didn't add return inside the function.
- Rule:
  - N/A

### Multiplication Table
- Concepts: 
  - Create and use function showing result of multiplication
  - Use looping with function to show the result of multiplication 1 to 10 by a number
- Input: 
  - A single variable
- Output: 
  - Result in a series of console log showing the result of the multiplication
- Directions:
  - Create a variable with a number.
  - Create a function that receive one parameter.
  - Inside the function create a loop from 1 to 10
  - Use the parameter to multiply inside the loop
  - Create a variable inside the function to receive the result.
  - With each iteration, add string in the following format '1 x 1 = 1\n2 x 2 = 4\n...' all the way up to 10, so the result with be a string of multiplication table.
  - Return the result
  - Use console log to call the function
- Rule:
  - N/A

## Challenging

### Simple Calculator
- Concepts: 
  - Variable
  - Conditional
  - Looping
  - Function
- Input: 
  - A string of calculations, e.g. case1:"1 add 3; 1 add 4;" or case2:"1 add 2 minus 3"
- Output:
  - Calculation result, e.g. case1: "4; 5" or case2: "0"
- Directions:
  - Create a variable containing the string for calculations
  - Use str.split() to split the string
  - Create functions for each arithmetic operation
  - Use conditional to do the calculation
  - Arrange the resulting string for output to console.log
- Rule:
  - Try not to use modular function, save that for next week.
