# Groupie Tracker

## Introduction
Groupie Tracker is a project that uses Golang as the web server. The server receives an API, which is used to create a stylish webpage to display data dynamically based on user actions. The project focuses on creating webpage events based on user interactions.

## Installation and Running the Program
1. Clone the repository:
    ```bash
    git clone https://learn.zone01kisumu.ke/git/tesiaka/groupie-tracker.git
    ```
2. Navigate into the cloned repository:
    ```bash
    cd groupie-tracker
    ```
3. Run the program:
    ```bash
    go run .
    ```
4. Open this URL in your browser:
    ```bash
    http://localhost:8089
    ```

## API Structure
The given API consists of four parts:
1. **Artists**: Contains information about bands and artists, including their name(s), image, starting year, the date of their first album, and the members.
2. **Locations**: Consists of the last and/or upcoming concert locations of the artists.
3. **Dates**: Contains details about the last and/or upcoming concert dates.
4. **Relation**: Links the other parts (artists, dates, and locations).

## Objective
The goal of the project is to create a user-friendly website where the band's information is displayed through various data visualizations, such as:
- Blocks
- Cards
- Lists
- Pages
- Graphics


## Events and Actions
This project also emphasizes creating events and actions for visualization. One of the main features involves a **client-server event**, which is an action triggered by the client (or time, etc.) that communicates with the server to receive information. This follows the request-response model:

[Request-Response](https://en.wikipedia.org/wiki/Request%E2%80%93response)

## Collaborators
- [Teddy Siaka](https://learn.zone01kisumu.ke/git/tesiaka)
- [Otieno Rogers](https://learn.zone01kisumu.ke/git/oragwelr)
