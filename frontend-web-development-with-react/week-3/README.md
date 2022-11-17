# Frontend Web Development with React Mentoring Week-3

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

- Understand how routing works.
- Able to implement React Router DOM to create SPA with multiple page.
  - Able to use URL Params to pass parameter from URL.
- Able to implement ChakraUI as UI element.
  - Able to use Common Elements from ChakraUI.

## Pre-requisites

- Code editor (e.g. VS Code).
- Web Browser
- Nodejs and PNPM installed in local computer

## Basic

### Simple Book Collection

- Concepts:
  - React Router DOM.
    - Able to create multipage website.
- Input:
  - A list of books with ISBN in JSON format: booksData = [{"isbn":"0747558191", "title":"Harry Potter and the Philoshoper's Stone","description":"Harry Potter has never even heard of Hogwarts when the letters start dropping on the doormat at number four, Privet Drive. Addressed in green ink on yellowish parchment with a purple seal, they are swiftly confiscated by his grisly aunt and uncle. Then, on Harry's eleventh birthday, a great beetle-eyed giant of a man called Rubeus Hagrid bursts in with some astonishing news: Harry Potter is a wizard, and he has a place at Hogwarts School of Witchcraft and Wizardry. An incredible adventure is about to begin!"}, {"isbn":"0439708184", "title":"Harry Potter and the Sorcerer's Stone","description":"Harry Potter has no idea how famous he is. That's because he's being raised by his miserable aunt and uncle who are terrified Harry will learn that he's really a wizard, just as his parents were. But everything changes when Harry is summoned to attend an infamous school for wizards, and he begins to discover some clues about his illustrious birthright. From the surprising way he is greeted by a lovable giant, to the unique curriculum and colorful faculty at his unusual school, Harry finds himself drawn deep inside a mystical world he never knew existed and closer to his own noble destiny."}, {"isbn":"0439064872","title":"Harry Potter and the Chamber of Secret","description":"Ever since Harry Potter had come home for the summer, the Dursleys had been so mean and hideous that all Harry wanted was to get back to the Hogwarts School for Witchcraft and Wizardry. But just as he's packing his bags, Harry receives a warning from a strange, impish creature who says that if Harry returns to Hogwarts, disaster will strike. And strike it does. For in Harry's second year at Hogwarts, fresh torments and horrors arise, including an outrageously stuck-up new professor and a spirit who haunts the girls' bathroom. But then the real trouble begins -- someone is turning Hogwarts students to stone. Could it be Draco Malfoy, a more poisonous rival than ever? Could it possibly be Hagrid, whose mysterious past is finally told? Or could it be the one everyone at Hogwarts most suspects...Harry Potter himself!"}]
- Output:
  - A website showing the books in a list and link for book details that will show in another page. And a search page to filter books by name.
- Directions:
  - Example is available at folder "/books-collection"
  - Use CRA to create React app.
  - Install using pnpm react-router-dom
  - Aktifkan React Router DOM mengikuti halaman berikut: https://neurons.ruangguru.com/course/frontend-development/frontend-web-development-with-react/react-router-dom/introduction-to-react-router-dom
  - Import Routes and Route ke App.js
  - Create three routes:
    - endpoint "/" as "Home"
    - endpoint "book" as "Book List"
    - endpoint "book/:isbn" as "Book Detail" as child of "book"
  - Create folder "/routes" in "/src"
  - Create Components for each route
    - "Home" contains "Hello, welcome to my Book Collection"
      - Add Link that will redirect to endpoint "book"
    - "Book List" contains List of Books
      - Setup the input booksData here
      - Create Component to render the data in folder "/cards"
        - Book name
        - Link to book detail using ISBN as param.
    - "Book Detail" contain Information about the book. 
      - Grab the isbn from useParams(), pass it to BookDetail Component that we'll create
      - Create Component to render the data in folder "/cards"
        - This is not ideal, but we also setup the input booksData here
        - We'll search by isbn the actual book.
        - Then we'll show
          - Book name
          - Book description
  - Play around some more creating or modifying new route
  - If you want, you can store the data as JSON and load the JSON file, or create a wrapper to simulate API call.
  - You can also use fetch() and actually use 3rd party API.
  - There is no styling, feel free to style it before hand.
  
### Simple Book Collection with Styling

- Concepts:
  - Same as Simple Book Collection
  - ChakraUI as styling component
- Input:
  - Simple Book Collection
- Output:
  - Simple Book Collection with ChakraUI styling
- Directions:
  - Re-use the Simple Book Collection
  - Style using ChakraUI, remember to install ChakraUI first using pnpm

## Challenging

### Use 3rd party API and showcase the data

- Concepts:
  - Free
- Input:
  - Free
- Output:
  - Free
- Directions:
  - Free
