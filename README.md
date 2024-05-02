# Story Publishing Platform (not done yet)
This is a web platform designed for publishing and reading stories. 
Users can publish their own stories, read stories from other users, like and comment on them. 
The platform also includes an admin page for managing users, stories, and comments.

## Technologies Used
- **Golang**: Programming language used for backend development.
- **Gin Framework**: Web framework used for building the API endpoints.
- **Redis**: In-memory data structure store used for caching.
- **PostgreSQL**: Relational database used for storing user data, stories, and comments.
- **JWT (JSON Web Tokens)**: Used for authentication and authorization.
- **Swagger**: API documentation tool for documenting the API endpoints.
- **Unit Testing**: Test cases written to ensure the correctness of the functionality.
- **Pagination**: Implemented cursor-based and offset pagination for handling large datasets.
- **Clean Architecture**: Followed clean architecture principles to keep the codebase organized and maintainable.
- **Logging**: Logging implemented to track application activities and errors.

## Features
- **Story Platform**: Displays a list of all stories with filtering options on the homepage.
- **Authentication & Authorization**: Users can register with an email and password, and then login to access their personal account.
- **User Account**: Logged-in users can write and publish stories, save stories as drafts, and access their own profile.
- **Story Interaction**: Users can read public stories, like them, leave comments, and reply to comments.
- **Admin Panel**: Admins have the ability to delete users, stories, comments, etc.

## Installation
1. Clone the repository:

```bash
git clone https://github.com/isido5ik/StoryPublishingPlatform.git
```
2. Install dependencies:
```bash
go mod tidy
```
3. Set up PostgreSQL and Redis databases and configure the connection details in the application.

4. Run the application:
```bash
go run main.go
```

## Usage
1. Register a new user account or login with existing credentials.
2. Explore stories on the homepage and interact with them by liking, commenting, etc.
3. Write your own stories and publish them on the platform.
4. Admins can manage users, stories, and comments through the admin panel.

## Contributing
Contributions are welcome! If you have any ideas, suggestions, or improvements, feel free to open an issue or create a pull request.

## Contact
For any inquiries or support, please contact do5ymzh4n@gmail.com.
