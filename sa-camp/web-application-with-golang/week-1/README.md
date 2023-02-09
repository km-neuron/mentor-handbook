# Web Application Mentoring Week-1

## Basic Rules

- **This is not a demo!**
- The challenges are for reference, can be used as is, customized or completely use your own challenges. Keep in mind the Learning Objectives!
- Mentor will be the facilitator, guiding the mentee to solve the tasks.
  - Share screen method
    - One person will share their screen, can be the mentor and one of the mentee.
    - Mentee will take turn in giving the correct syntax to be typed by person sharing his/her screen.
  - Using tools such as replit <https://replit.com/>
    - Everyone can login using their own account.
    - Mentor can have a tab for each mentee to monitor their work.
    - Details over at Discord: <https://discord.com/channels/996138478817513632/996141512704401579/1019923694254043136>
  
## Learning Objectives

- Understand how to interact with files (I/O)
- Understand how to create a HTTP server
- Understand how to use HTTP client
- Understand REST API.
- Able to retrieve data using REST API.
- Able to serve data using REST API.

## Pre-requisites

- Code editor (e.g. VS Code).
- Golang installed in local computer
- Postman installed in local computer
- Postgres installed in local computer

## Basic

### Read and write simple JSON file

- Concepts:
  - IO File, IO JSON
- Input:
  - A simple JSON structure inside a file, can be as simple as first_name, last_name, birthdate
- Output:
  - Will read the JSON file and output to terminal, then update the JSON file with a new data.
- Direction:
  - Create a simple JSON file containing one data.
  - Create a struct to hold the JSON data.
  - Read the file using io
  - Read and convert the JSON content to struct.
  - Output the JSON struct to terminal.
  - Append a new JSON data to struct
  - Write the new JSON to file, can use append or overwrite the file.

### Connect to 3rd party API

- Concepts:
  - Connect to a free 3rd party API
- Input:
  - Free 3rd party API data
- Output:
  - The data retrieved shown to terminal
- Directions:
  - Create the client
  - Connect to 3rd party API data
  - If processing is needed then process it.
  - Output the data to terminal.

### Simple web server

- Concepts:
  - Serve a simple calculator web server
- Input:
  - None
- Output:
  - A HTML file that will have two input boxes and 4 buttons (+, -, *, /)
- Directions:
  - Setup ListenAndServe()
  - Can use template
  - Show a HTML file with form containing 2 input boxes and 4 buttons (+, -, *, /)
  - On click of the button will perform the calculation and output to terminal or redirect to new page with the calculation and the result, or on the same page, simply create a container.

### Quote App

- Concepts:
  - REST API
  - HTTP Client
  - Handling JSON format
  - HTTP Server
- Input:
  - Find any free quote API or choose from these two:
    - <https://animechan.vercel.app/>
    - <https://zenquotes.io/>
- Output:
  - Single qoute output to HTML page.
- Directions:
  - Read the documentation for the API.
  - Pick the random endpoint or endpoint that will output a single quote, e.g.
    - <https://animechan.vercel.app/api/random>
    - <https://zenquotes.io/api/random>
  - Use Postman to send the API and see what the output format is.
  - Create HTTP Client that will send the request.
  - Create HTTP Server that will output the single quote to the browser.
  - Style the quote as desired.
  - See example at "/quoteApp/" folder.

## Challenging

### Full application for anime or pokemon or other API

- Concepts:
  - Create a web server that will retrieve data from 3rd party API and then display the title only with a link to a detail page. If clicked, then the detail page will have more data.
- Input:
  - 3rd party API data
- Output:
  - A web application that can display data and have a favorite button implemented.
- Directions:
  - Create a client first to get the data from 3rd party API
  - Create a server to serve the data
  - First show the list with link to a detail page using route
  - When the link is clicked then the detail page will show detail data based on the title.
  - Have a link to get back to the list.

### Quote App with Database

- Concepts:
  - Same as Quote App
  - Database CRUD
    - Create database
    - Create Table
    - Insert Data into Table
    - Select Data from Table
- Input:
  - Same as Quote App
- Output:
  - Quote inserted to database
- Directions:
  - Best to continue using the Quote App.
  - Create database first to store the quotes.
  - Create table to store the quotes.
  - Add insert into database for the quotes, if there is endpoint that can get .more than 1 data, then use that endpoint.
  - Use select to return the count of quotes in the table.
  - And use select to return one quote of the day.
