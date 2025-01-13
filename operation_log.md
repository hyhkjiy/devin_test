# Operation Log for Web Application Development

## Setup Phase

1. Created project structure with backend and frontend directories
2. Initialized Go module with Iris framework and SQLite support
3. Created frontend React application with TypeScript
4. Installed and configured Tailwind CSS
5. Implemented JWT authentication for user login
6. Enhanced user response format with proper error handling
7. Completed user management features (registration and login)

## API Endpoints Planning

### User Management
1. POST /api/users/register
   - Purpose: Register new user
   - Request: { username: string, password: string }
   - Response: { id: int, username: string }

2. POST /api/users/login
   - Purpose: User authentication
   - Request: { username: string, password: string }
   - Response: { token: string }

### Article Management
1. POST /api/articles
   - Purpose: Create new article
   - Auth: Required
   - Request: { title: string, content: string }
   - Response: { id: int, title: string, content: string, publishTime: datetime }

2. GET /api/articles
   - Purpose: List all articles
   - Auth: Not required
   - Response: [{ id: int, title: string, publishTime: datetime }]

3. GET /api/articles/:id
   - Purpose: Get article details
   - Auth: Not required
   - Response: { id: int, title: string, content: string, publishTime: datetime, author: string }

## Database Schema

### Users Table
```sql
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
```

### Articles Table
```sql
CREATE TABLE articles (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    user_id INTEGER NOT NULL,
    publish_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
```
## Implementation Progress
1. Implemented user registration endpoint with SQLite storage
2. Added password hashing using bcrypt
3. Configured CORS middleware for frontend integration
4. Tested user registration functionality
5. Using in-memory SQLite database for development
