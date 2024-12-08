
# Auth REST API Service

This project is an implementation of an authentication and authorization REST API service using **Golang** and the **Gin framework**. The service supports:

1. **User Sign Up**: Register a new user with email and password.
2. **User Sign In**: Authenticate user credentials and issue a JWT token.
3. **Authorization**: Verify token validity for secured routes.
4. **Token Revocation**: Revoke tokens to log out users.
5. **Token Refresh**: Refresh expired tokens for continued access.

---

## **Features**
- Secure authentication using JWT.
- In-memory user and token storage for simplicity (Can use a database, but compatibility may vary across environments).
- Error handling with proper HTTP response codes.
- Middleware for token validation.

---

## **Setup Instructions**

### **Prerequisites**
- **Go 1.20+**
- **Git**

### **Steps to Run**

1. Clone the repository:

   ```bash
   git clone https://github.com/Jatin642-rg/Auth-login-logout.git
   cd Auth-login-logout
   ```

2. Create a `.env` file in the project root:

   ```plaintext
   JWT_SECRET=your_secret_key
   ```

3. Run the application:

   ```bash
   go run main.go
   ```

4. The API service will start on `http://localhost:8080`.

---

## **API Endpoints**

### **1. Sign Up**
- **URL:** `/signup`
- **Method:** `POST`
- **Description:** Registers a new user.
- **Request Body:**
  ```json
  {
    "email": "test@example.com",
    "password": "password123"
  }
  ```
- **Responses:**
  - **Success:** `201 Created`
    ```json
    {
      "message": "User created successfully"
    }
    ```
  - **Failure:** `400 Bad Request` or `409 Conflict`

---

### **2. Sign In**
- **URL:** `/signin`
- **Method:** `POST`
- **Description:** Authenticates a user and returns a JWT token.
- **Request Body:**
  ```json
  {
    "email": "test@example.com",
    "password": "password123"
  }
  ```
- **Responses:**
  - **Success:** `200 OK`
    ```json
    {
      "token": "<JWT_TOKEN>"
    }
    ```
  - **Failure:** `401 Unauthorized`

---

### **3. Authorization**
- **URL:** `/authorize`
- **Method:** `GET`
- **Description:** Validates a provided token.
- **Headers:**
  ```plaintext
  Authorization: Bearer <JWT_TOKEN>
  ```
- **Responses:**
  - **Success:** `200 OK`
    ```json
    {
      "message": "Access granted",
      "email": "test@example.com"
    }
    ```
  - **Failure:** `401 Unauthorized`

---

### **4. Revoke Token**
- **URL:** `/revoke`
- **Method:** `POST`
- **Description:** Revokes a user's token.
- **Headers:**
  ```plaintext
  Authorization: Bearer <JWT_TOKEN>
  ```
- **Responses:**
  - **Success:** `200 OK`
    ```json
    {
      "message": "Token revoked successfully"
    }
    ```
  - **Failure:** `401 Unauthorized`

---

### **5. Refresh Token**
- **URL:** `/refresh`
- **Method:** `POST`
- **Description:** Refreshes an expired token.
- **Headers:**
  ```plaintext
  Authorization: Bearer <JWT_TOKEN>
  ```
- **Responses:**
  - **Success:** `200 OK`
    ```json
    {
      "token": "<NEW_JWT_TOKEN>"
    }
    ```
  - **Failure:** `401 Unauthorized`

---

## **Project Structure**

```
Auth-login-logout/
├── main.go               # Entry point for the application
├── handlers/             # API route handlers
│   ├── auth.go           # Authentication and authorization handlers
│   ├── token.go          # Token management handlers
├── middleware/           # Middleware for token validation
│   ├── auth_middleware.go
├── models/               # Data models
│   ├── user.go           # User data structure
├── storage/              # In-memory storage
│   ├── memory_store.go
├── utils/                # Utility functions
│   ├── jwt.go            # Functions for JWT token creation and validation
├── .env                  # Environment variables
├── README.md             # Documentation
```

---

## **Testing the API**

### **Using cURL**

1. **Sign Up**
   ```bash
   curl -X POST -H "Content-Type: application/json"    -d '{"email":"test@example.com","password":"password123"}'    http://localhost:8080/signup
   ```

2. **Sign In**
   ```bash
   curl -X POST -H "Content-Type: application/json"    -d '{"email":"test@example.com","password":"password123"}'    http://localhost:8080/signin
   ```

3. **Authorize**
   ```bash
   curl -X GET -H "Authorization: Bearer <JWT_TOKEN>"    http://localhost:8080/authorize
   ```

4. **Revoke Token**
   ```bash
   curl -X POST -H "Authorization: Bearer <JWT_TOKEN>"    http://localhost:8080/revoke
   ```

5. **Refresh Token**
   ```bash
   curl -X POST -H "Authorization: Bearer <JWT_TOKEN>"    http://localhost:8080/refresh
   ```

---

## **License**
This project is licensed under the MIT License.
