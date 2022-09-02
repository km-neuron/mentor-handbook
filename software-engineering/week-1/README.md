# Software Engineering Mentoring Week-1

## Basic Rules
- **This is not a demo!**
- Mentor will be the facilitator, guiding the mentee to solve the tasks.
- One person will share their screen, can be the mentor and one of the mentee.
- Mentee will take turn in giving the correct commands to be typed by person sharing his/her screen.

## Learning Objectives

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

## Pre-requisites

- Install code editor (e.g. VS Code)
- Install gitBash for Windows or git for Linux based distro (in WSL for windows)
- Create account in github

## Basic

### World Culinary Collection
- Concepts: 
  - Terminal usage
    - Create folder
    - Create file
    - Move between folders
    - Move files between folders
  - Code Editor
    - Open file
    - Edit file
    - Save file
- Input: 
  - Ideas on famous food in the world
- Output:
  - One text file and one picture file of each culinary in subfolders by country
  - One summary file in primary folder.
  - Final structure example, see folder "world culinary collection"
- Directions: 
  - Create one folder to be used as primary folder
  - Create one summary list file in primary folder
  - Create subfolder for countries, minimum 3 subfolder.
  - Decide on 3 countries that have famous culinary food.
  - Name the famous culinary food
  - Create empty file in folder country with the name of the food.
  - Google the culinary food
    - Find text to be used, can be the summary of the food or history of the food or the recipe.
    - Put the text into a text file using text editor
    - Find picture to represent the food.
      - Download the picture
      - Move the file to folder country where the text file is.
- Rule
  - Minimum of 3 countries
  - Minimum of 1 food for each country
  - One text file and one image file for each food

### My first Git
- Concepts: 
  - git config
  - git init
  - git add
  - git commit
  - git status
- Input: 
  - "Hello World"
- Output: 
  - one file containing "Hello World" inside a local git repo.
- Directions:
  - Create a folder as working directory
  - Run git config and update the data
  - Run git init inside the working directory
  - Create an empty file
  - Edit the file and write "Hello World"
  - Run git add
  - Run git commit
  - Run git status and explain
- Rule:
  - Explain git status
  - Can ask each mentee share their screen and do this step ( if time permits )

## Challenging

### My Git Culinary Collection
- Concepts: 
  - git init
  - git branch
  - Basic git commands ( add, commit, status )
  - git checkout to switch branch
  - git merge
Put World Culinary Collection into Git. Initialize a local repo, create the primary folder and summary.txt, then commit. For each country, create a new branch, then add the culinary together with the subfolder. Remember to update the summary file in primary folder. Then switch branch to master/main and do a merge.
- Input: 
  - Data from World Culinary Collection or create new ones
- Output:
  - A local git repo containing branches of country and culinary data.
- Directions:
  - Create primary folder
  - Init a local repo inside primary folder
  - Create summary file
  - git add and commit
  - Git branch for each country
  - Create folder for each country
  - Add the culinary file and image to each country
  - Update summary file in primary folder
  - git checkout to main branch
  - git merge country branch to main
- Rule:
  - Same rule as World Culinary Collection
  - Each country have it's own branch, the order for directions does not matter (create folder then branch is OK, create branch then create folder is OK).
  - If possible repeat this 5 times, each student sharing their screen and add new culinary ( googling, add, commit, merge) with other mentee supporting and mentor facilitating if mistakes are made.

