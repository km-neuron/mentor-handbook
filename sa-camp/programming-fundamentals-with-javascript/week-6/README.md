# Programming Fundamentals with Javascript Mentoring Week-6

## Basic Rules
- **This is not a demo!**
- ### **The challenges are for reference, can be used as is, customized or completely use your own challenges. Keep in mind the Learning Objectives!**
- Mentor will be the facilitator, guiding the mentee to solve the tasks.
  - Share screen method
    - One person will share their screen, can be the mentor and one of the mentee.
    - Mentee will take turn in giving the correct syntax to be typed by person sharing his/her screen.
  - Using tools such as replit <https://replit.com/>
    - Everyone can login using their own account.
    - Mentor can have a tab for each mentee to monitor their work.
    - Details over at Discord: <https://discord.com/channels/996138478817513632/996141512704401579/1019923694254043136>

## Learning Objectives

- Understand how to use CommonJS Modules in NodeJS
    - Able to use CommonJS export
    - Able to use CommonJS require
    - Able to use pass variables between two files using CommonJS module
- Understand how to use ES Modules in NodeJS
    - Able to use ESModules export default
    - Able to use ESModules import
    - Able to use ESModules named export
    - Able to use pass variables between two files using ESModules
- OOP (Theoritical)
    - Explain basic OOP theory
        - class as a blueprint
        - property as variable in OOP
        - method as function in OOP
        - instantiate using new
        - this as reference to self in OOP
    - Explain method chaining theory
- Callback
    - Able to use create callback function
- Promise
    - Able to create a promise
    - Able to use a promise
- Async Await
    - Able to create an async function
    - Able to use await inside the function
- Able to fetch data from server using REST API.
- Able to create a dummy JSON server.
- Able to present the data from the server.
- Understand how to use environment variables
    - Able to create an environment variable in his/her own computer
    - Able to use the environment variable using NodeJS process
    - Able to user environment variables using dotenv package

## Pre-requisites

- Code editor (e.g. VS Code).
- Web based playground
    - Web Browser with Developer Tools.
        - Bisa dengan membuat index.html berisi DOM yang akan refresh ketika kita mengetik kode JS (perlu effort dari mentor)
        - Bisa dengan extension LiveShare atau lainnya.
        - Atau metode lainnya.
    - Online playground such as jsbin.com

## Basic

### Simple Calculator with ESModules
- Concepts: 
  - Function and simple math calculation
  - PNPM setup
  - ES Modules
- Input: 
  - None
- Output:
  - Two files 
- Directions: 
- Create a new folder for the calculator files
    - Run pnpm init
    - Show package.json
  - Use the file from Simple Calculator
    - Export the functions using ESModules, using named export in simple calculation file
    - Use ESModules import to get functions from the simple calculation file in process file
    - Run the process file and show that we're using the process file to run the calculator
- Rule
  - N/A

### Simple Shopping using Promise
- Concepts: 
  - Create and use Promise
- Input: 
  - Same as above
- Output: 
  - Same as above
- Directions:
  - Same as above, but using Promise
- Rule:
  - N/A

### Quote App

- Concepts:
  - REST API
  - Data Fetch
- Input:
  - Find any free quote API or choose from these two:
    - https://animechan.vercel.app/
    - https://zenquotes.io/
- Output:
  - Single qoute output in Postman.
- Directions:
  - Pick the random endpoint or endpoint that will output a single quote, e.g.
    - https://animechan.vercel.app/api/random
    - https://zenquotes.io/api/random
  - Use Postman to send the API and see what the output format is.

### Simple Environment Variables
- Concepts: 
  - PNPM
  - Environment variables
- Input: 
  - Environment variables
- Output: 
  - Can display the environment variables
- Directions:
  - Use the terminal/command prompt to EXPORT a variable, for example name with mentee name in his/her own computer, i.e. EXPORT nama=mentee_name
  - Create a folder and setup using PNPM init
  - Create a file and inside use process.env to grab the variable and output using console.log
  - It should work, it not check if perhaps the variable was not exported properly in his/her own computer.
  - Don't take too long, next
  - Install dotenv
  - Create a file callend .env
  - Fill the file with some variables, for example nama and favorite
  - Create a new file and use dotenv to console log the variables inside .env file
  - Experiment more
- Rule:
  - N/A

## Challenging

### Advanced Shopping
- Concepts: 
  - Able to use either Callback, Promise or Async Await
- Input:
  - Same as above
- Output:
  - An array of string containing the repeated string 'I bought "product_name" for "product_price"' in the middle.
  - The beginning of the array will have the string 'I have "start_money_in_wallet"'
  - The end of the array will have the string 'I still have "end_money_in_wallet left"'
- Directions:
  - Create callback/promise/async await
  - Output should be stored in a string
  - Create global array to store the string
  - Create the start string
  - Create the end string
  - Combine the strings into an array
- Rule:
  - N/A
