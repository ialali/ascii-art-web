# Ascii-art-web

## Description

Ascii-art-web is a server-based application that brings a graphical user interface to the ASCII art creation process. Building upon the ascii-art project, this web application allows users to generate ASCII art using various banners such as shadow, standard, and thinkertoy. The server handles HTTP endpoints to render the main page and process user input for creating customized ASCII art.

## Authors

[Your Name]

## Usage: How to Run

1. Clone the repository:

   ```bash
   git clone [repository_url]
Navigate to the project directory:

bash
Copy code
cd ascii-art-web
Run the server:

bash
Copy code
go run main.go
Open your web browser and go to http://localhost:8080 to access the ASCII art web interface.

Implementation Details: Algorithm
The server, written in Go, utilizes standard Go packages to handle HTTP requests and responses. HTML templates are stored in the templates directory at the project root. The code follows good practices to ensure readability, maintainability, and adherence to Go standards.

Instructions
The main page features a text input for user input, radio buttons to select banners (shadow, standard, thinkertoy), and a button to submit a POST request to '/ascii-art'.
The result of the POST request can be displayed either in the route '/ascii-art' or appended to the home page, depending on the user's preference.
HTTP status codes are handled appropriately, responding with OK (200) for successful requests, Not Found for missing resources, Bad Request for incorrect requests, and Internal Server Error for unhandled errors.
Example Usage
Here's an example of how to use the application:

Enter text in the text input field.
Select a banner using radio buttons.
Click the submit button to generate and display the ASCII art.
