
# Groupie-Tracker-Search-bar

Groupie-Tracker is a web application that provides information about musical artists, including their locations, concert dates, and band members. This application fetches data from an external API and presents it in a user-friendly format, allowing users to explore various artists and their concert details.

## Features

- **Homepage**: Explore and view various artists' information.
- **Artist Search**: Search for artists by name using the search bar with autocomplete suggestions or manual search.
- **Artist Details**: View detailed information about artists including their band members, first album, concert dates, and locations.
- **Error Handling**: Custom error pages for various scenarios such as 404 errors, server errors, and incorrect HTTP methods.
- **Responsive Design**: Styled with custom CSS for better user experience on various devices.

## API Integration

Groupie-Tracker retrieves its data from the [Groupie Trackers API](https://groupietrackers.herokuapp.com/api). It fetches details such as:
- Artists
- Locations
- Concert dates
- Relations (mapping artists to their concert dates and locations)

## Getting Started

### Prerequisites

- Go 1.16 or higher
- Internet connection (to fetch API data)

### Installation

1. Clone the repository:

   ```bash
   git clone https://learn.zone01kisumu.ke/git/tesiaka/groupie-tracker.git
   ```

2. Navigate to the project directory:

   ```bash
   cd groupie-tracker
   ```

3. Install dependencies:

   ```bash
   go mod tidy
   ```

4. Run the application:

   ```bash
   go run main.go
   ```

5. Open your browser and navigate to:

   ```
   http://localhost:8089/
   ```

### Project Structure

- `main.go`: Entry point of the application, handling routing and server setup.
- `handlers/`: Contains various handler functions for serving web pages and handling API data.
  - `homepage.go`: Serves the homepage.
  - `artistinfo.go`: Fetches and displays artist information.
  - `aboutus.go`: Displays the "About Us" page.
  - `errorhandlers.go`: Handles custom error pages (500, 404, incorrect method, etc.).
  - `searchartist.go`: Manages the manual search for artists by name, directing users to their detailed information upon finding a match.
  - `staticserver.go`: Serves static files like CSS, images, and JavaScript.
- `api/`: Handles communication with the external API.
  - `api.go`: Contains functions to fetch data from the API (artists, locations, dates, relations).
- `models/`: Contains data structures for handling and parsing API responses.
  - `models.go`: Defines structures like `Artist`, `Location`, `Date`, and `Relation`.
- `autocomplete/`: Contains logic for handling search autocompletion.
- `templates/`: HTML templates for rendering pages.
- `static/`: CSS, images, and JavaScript files.

### Search Functionality

- **Manual Artist Search**: Users can search for artists using the search bar. The `searchartist.go` file handles the manual search functionality, where the input is compared with the list of artists retrieved from the API. If a match is found, the user is redirected to the detailed artist information page.

### Error Handling

The application includes custom error handling for various scenarios:

- **404 Page**: Displayed when the requested page is not found.
- **500 Page**: Shown when there is a server error, with a template defined in `Errortemplate/error500.html`.
- **Wrong Method Page**: Displays when a user tries to access a route with an invalid HTTP method, using the template in `Errortemplate/wrongmethodused.html`.
- **No Internet Connection**: Handled with a custom template, ensuring the user is aware of connectivity issues.
- **Artist Not Found**: If a user tries to search for an artist that does not exist, they are redirected to a bad request error page.

### API Data Structure

The app fetches the following from the API:
- **Artists**: General artist information (ID, name, members, creation date, first album, image).
- **Locations**: Locations where concerts will be held.
- **Dates**: Concert dates for each artist.
- **Relations**: Links concert dates to specific locations for each artist.

### License

This project is licensed under the MIT License.

---

### Future Improvements

- Add pagination to artist listings.
- Improve UI/UX with better mobile responsiveness.
- Implement caching to reduce API requests and improve performance.
- Add more error handling and logging for debugging purposes.

## Collaborators

[Teddy siaka](https://learn.zone01kisumu.ke/git/tesiaka)

[OTIENO ROGERS](https://learn.zone01kisumu.ke/git/oragwelr)