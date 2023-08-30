# Bytive Task Management System

The Bytive Task Management System is a simple web application built in Go that allows users to manage tasks and projects. This application provides APIs to create, retrieve, update, and delete tasks and projects.

## How to Use

### Prerequisites

Before you begin, make sure you have the following installed:

- Go (1.17 or later)
- Docker (for containerization, optional)

### Installation

1. Clone the repository to your local machine:

   ```bash
   git clone https://github.com/mdNoman21/bytive-task.git
   ```

2. Navigate to the project directory: 
   ```bash
   cd bytive-task
   ```

   
### Running the application 

Option 1:Without Docker 

1. Install project Dependencies
   ```bash 
   go mod download
   ```

2. Build and run the application
   ```bash
   go run main.go
    ```



Option 2:With Docker
1.  Build the docker image and run a Docker container from the image
   ```bash 
    docker build -t bytive-task 
    docker run -p 8080:8080 bytive-task
      ```



### JSON Request Format
- When sending JSON requests to create or update project information, ensure that you follow the specified format for the date, start_time, and end_time fields. These fields use the ISO 8601 format, which includes both date and time information.

1. Date Format
The date field should be formatted using the ISO 8601 date format with time zone information:
      ```bash
      "date": "2023-08-30T00:00:00Z"
      ```

2. Start Time Format
The start_time field should also be formatted in ISO 8601 format, including the time zone:
      ```bash
      "start_time": "2023-08-30T16:00:00Z"
      ```

3. End Time Format
Similarly, the end_time field should follow the ISO 8601 format:

      ```bash
      "end_time": "2023-08-30T17:30:00Z" 
      ```
       

##  Usage
Once the application is up and running, you can interact with it using APIs. Here are some of the available endpoints:

Create a Project: POST /createProject
Get All Projects: GET /getProjects
Get a Project by ID: GET /getProject/:id
Update All Project End Times: PUT /updateEndTimeAll?timeToAdd={timeToAdd}
Delete a Project: DELETE /deleteProject/:id
The application will be available at http://localhost:8080
