# Programming with Javascript Mentoring Week-2

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

- Modular Function
    - Able to use function inside another function
    - Able to pass parameter between function
- Array
    - Able to create an array
    - Able to access value inside an array
    - Able to change the value inside an array
    - Able to create two dimensional array
    - Able to access two dimensional array to get value or to change value
- Object Literal
    - Able to create an object
    - Able to access value inside an object
    - Able to change the value inside an object
    - Able to create nested object
    - Able to access nested object to get value or go change value
- Array of Object
    - Able to create an array of object
    - Able to access value inside the array of object to get the value
    - Able to add new object into the array of object
    - Able to remove object from array of object

## Pre-requisites

- Code editor (e.g. VS Code).
- Web based playground
    - Web Browser with Developer Tools.
        - Bisa dengan membuat index.html berisi DOM yang akan refresh ketika kita mengetik kode JS (perlu effort dari mentor)
        - Bisa dengan extension LiveShare atau lainnya.
        - Atau metode lainnya.
    - Online playground such as jsbin.com

## Basic

### Simple Attendance Data with Array
- Concepts: 
  - Create and make changes to an array
- Input: 
  - One string separated by comma containing 'Present/Absent/Sick'
- Output:
  - An array of the data from the string
- Directions: 
  - Create a string containing value of "Present, Absent, Sick" multiple times, e.g. "Present, Absent, Sick, Present, Present, Present"
  - Create an empty array
  - Use looping to store the data into the array
  - show the result using console.log
  - show a faster method using str.split()
  - show the result using console.log
- Rule
  - N/A

### Simple Attendance Data with Object
- Concepts: 
  - Create and make changes to an object
- Input: 
  - One string containing names, e.g. ('Budi, Ani, Sigit')
  - array from previous challenge
- Output: 
  - Three object from the data inside the string
- Directions:
  - Create a string containing names, 3 should be enough
  - Create an empty object with key 'name' and 'data'
  - Create 3 objects, one for each name.
- Rule:
  - N/A

### Simple Attendance data with Array of Object
- Concepts: 
  - Create an array of object from existing array and object
- Input: 
  - Three objects from the previous challenge
  - Three Array / string from the previous challenge
- Output: 
  - Array of object containing the data
- Directions:
  - Create an empty array
  - Push each object into the array
  - Update the objects key "data" with data from an array
  - Console.log the array of object
- Rule:
  - N/A

## Challenging

### Advanced Attendance Data
- Concepts: 
  - Able to access array of object
  - Able to use modular function concept
- Input: ( We kind of reset the data here )
  - Three string containing 'Present, Absent, Sick' in random, make it 5-10 data each.
  - Three string containing names
- Output:
  - An array of object containing key of name, rawData, Present, Absent, Sick
  - name and rawData should be clear
  - Present, Absent, Sick will have the value of the number of times from the rawData they shows up.
- Directions:
  - Create the arrays, using function
  - Create the objects using function
  - Create the array of objects using function
  - Do calculation using function
  - Create a main function containing all the functions
  - Show result using console.log
- Rule:
  - N/A
