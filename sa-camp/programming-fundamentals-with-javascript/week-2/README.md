# Programming Fundamentals with Javascript Mentoring Week-2

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

- Understand how to use PNPM
    - Able to init a package
    - Able to install a package
    - Able to install dev package
    - Able to instal global package
- Understand how to use Primitive Data Type
    - Able to declare a variable
    - Able to assign value to variable
    - Able to get value from variable
- Understand how to use Conditional
    - Able to create a simple conditional using if else
    - Able to created nested conditional using if else if
    - Able to use switch case and break effectively


## Pre-requisites

- Code editor (e.g. VS Code).
- Web based playground
    - Web Browser with Developer Tools.
        - Bisa dengan membuat index.html berisi DOM yang akan refresh ketika kita mengetik kode JS (perlu effort dari mentor)
        - Bisa dengan extension LiveShare atau lainnya.
        - Atau metode lainnya.
    - Online playground such as jsbin.com

## Basic

### Simple PNPM
- Concepts: 
  - PNPM setup
  - PNPM Dev and Global
- Input: 
  - None
- Output:
  - Installation of nodemon in devDependencies and ifup in global 
- Directions: 
- Create a new folder
    - Run pnpm init
    - Show package.json
  - Install nodemon using --dev flag
    - Show package.json have nodemon as devDependencies
  - Install is-up using --global flag
    - Show package.json
    - Call is-up, see its README for details.
  - Move away from the package folder
    - Run is-up again, show that it works if we use --global
    - Run nodemon, show that it failed because it's installed as part of the package
  - Using Github with NodeJS
    - Create .gitignore for node_modules folder
    - Create git repo for the package folder
    - Link to Github
    - Upload to Github
    - Ask one mentee or all to download from Github
    - Then ask them to share screen and try to run nodemon which is a devDependencies
    - If it's not working, then ask them to run pnpm install
    - Run again, it should work now
    - Explain that this is why PNPM is powerful, it's good for cooperation with other people, having the same package, especially the version
- Rule
  - N/A

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
