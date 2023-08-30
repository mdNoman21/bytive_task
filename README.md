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
# Navigate to the project directory: 
   cd bytive-task
# Running the application 

# Option 1:Without Docker 
#  Install project Dependencies 
   go mod download
#  Build and run the application
   go run main.go
# The application will be available at http://localhost:8080



# Option 2:With Docker
#  Build the docker image
    docker build -t bytive-task .
# Run a Docker container from the image:
    docker run -p 8080:8080 bytive-task
# The application will be available at http://localhost:8080






   
