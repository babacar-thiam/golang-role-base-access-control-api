## **API Documentation**

### **Overview**
This API implements role-based access control (RBAC) for managing roles and users. Roles include `ADMIN`, `CLIENT`, and `PROVIDER`.

- **ADMIN**: Full access to manage roles and users.
- **CLIENT**: Limited access, mostly to their own data.
- **PROVIDER**: Similar to CLIENT but tailored for providers.

---

### **Authentication**

Use JWT for securing the endpoints. Include the `Authorization` header in requests:

```
Authorization: Bearer <token>
```

---

### **Endpoints**

#### **1. Seed Roles**

Automatically seeds the roles `ADMIN`, `CLIENT`, and `PROVIDER` when the application initializes.

---

#### **2. Get All Roles**

**Access**: `ADMIN` only

| Method | Endpoint | Description        |
|--------|----------|--------------------|
| GET    | `/roles` | Retrieve all roles |

**Example Request:**

```
GET /roles
Authorization: Bearer <admin-token>
```

**Example Response:**

```
HTTP 200 OK
[
  {
    "id": 1,
    "role_name": "ADMIN"
  },
  {
    "id": 2,
    "role_name": "CLIENT"
  },
  {
    "id": 3,
    "role_name": "PROVIDER"
  }
]
```

**Error Response:**

```
HTTP 403 Forbidden
{
  "error": "Access denied"
}
```

---

#### **3. Get All Users**

**Access**: `ADMIN` only

| Method | Endpoint | Description        |
|--------|----------|--------------------|
| GET    | `/users` | Retrieve all users |

**Example Request:**

```
GET /users
Authorization: Bearer <admin-token>
```

**Example Response:**

```
HTTP 200 OK
[
  {
    "id": 1,
    "username": "admin_user",
    "role": "ADMIN"
  },
  {
    "id": 2,
    "username": "client_user",
    "role": "CLIENT"
  }
]
```

**Error Response:**

```
HTTP 403 Forbidden
{
  "error": "Access denied"
}
```

---

#### **4. Get User Information by ID**

**Access**: Only for the authenticated user whose ID matches the request.

| Method | Endpoint      | Description                     |
|--------|---------------|---------------------------------|
| GET    | `/users/{id}` | Retrieve information for a user |

**Example Request:**

```
GET /users/2
Authorization: Bearer <user-token>
```

**Example Response:**

```
HTTP 200 OK
{
  "id": 2,
  "username": "client_user",
  "role": "CLIENT",
  "email": "client@example.com"
}
```

**Error Responses:**

- **Unauthorized Access (Invalid Token)**:
  ```
  HTTP 401 Unauthorized
  {
    "error": "Invalid or missing token"
  }
  ```

- **Access Denied (Trying to access another user’s data)**:
  ```
  HTTP 403 Forbidden
  {
    "error": "Access denied"
  }
  ```

---

### **Middleware and Authorization**

#### **Role-Based Middleware**

- **ADMIN Check**: Used for `/roles` and `/users` endpoints to ensure only admins can access.
- **Authenticated User Check**: Ensures the user can access only their own data using `id`.

---

### **Error Handling**

| HTTP Status Code | Description                        |
|------------------|------------------------------------|
| 401              | Unauthorized (Invalid or no token) |
| 403              | Forbidden (Access denied)          |
| 404              | Not found (Invalid resource ID)    |
| 500              | Internal server error              |

---