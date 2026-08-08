# Registration Algorithm
Receive request -> Validate request -> Trim email -> Convert email to lowercase -> Validate password -> Check if email exists => Hash password -> Create User model -> Save User -> Return Response

# Dependency Injection
Instead of: service := AuthService{}

inside the handler, we do:  handler := NewAuthHandler(service)
Why? Because tomorrow we can inject:
Mock service (for testing), Real service, Cached service -> without changing the handler.

# Login Flow
Receive request -> Normalize email -> Find user -> Compare password using bcrypt -> Generate JWT -> Return token

# JWT
User -> JWT Token -> Client stores token -> Future requests -> Authorization Header -> Middleware validates token

# Business rules
Name required, Email unique, Password ≥ 8 characters

# Failure Analysis
"What can go wrong?"
Email empty -> Validation error
Email already exists -> 409 Conflict
Database unavailable ->  Internal Server Error
Two users register simultaneously -> UNIQUE constraint protects us

# Registration
HTTP JSON -> RegisterRequest DTO -> AuthService -> User Model -> Repository -> Database

# Response
Database Model -> RegisterResponse DTO -> JSON

# Appointment Booking
Business: Reserve a slot. -> Failure Analysis: Double booking. -> Data Flow: AppointmentRequest DTO -> Appointment Model -> Repository -> DB -> AppointmentResponse DTO -> JSON -> DSA: Interval overlap.

# What data enters the system?
Example:{"email":"abhi@gmail.com", "password":"Secret123"}

# What should the system trust?
Answer: Almost nothing from the client. Validate everything.

# What should the database guarantee?
Examples: Unique email, Foreign keys, Not null, Transactions

# What data leaves the system?
Never return: Password, Password hash, Internal IDs that shouldn't be public, Sensitive metadata

# What happens if something fails?
Invalid JSON, Duplicate email, Slow database. Lost connection, Transaction rollback

# Register API
# Business:
| Rule              | Layer                 |
| ----------------- | --------------------- |
| Name required     | Service               |
| Email required    | Service               |
| Password required | Service               |
| Password >= 8     | Service               |
| Email normalized  | Service               |
| Email unique      | Database + Repository |
| Password hashed   | Service               |

# Goes wormg
| Problem                           | Solution                   |
| --------------------------------- | -------------------------- |
| Invalid JSON                      | Handler returns 400        |
| Empty name                        | Service validation         |
| Empty email                       | Service validation         |
| Weak password                     | Service validation         |
| Duplicate email                   | UNIQUE constraint          |
| DB down                           | Repository returns error   |
| Hashing fails                     | Service returns error      |
| Two users register simultaneously | Database UNIQUE constraint |

# Data Flow
Client sends: { "name":"Abhinandan", "email":"Abhi@gmail.com ", "password":"Secret123" }
 -> Handler -> dto.RegisterRequest

Why a DTO? -> Because this is HTTP data. Not database data.

Service receives -> dto.RegisterRequest
Service cleans it.  request.Email = strings.ToLower(strings.TrimSpace(request.Email))
Now, Abhi@gmail.com -> abhi@gmail.com

Service creates a Model.    
model.User{
    Name: request.Name,
    Email: request.Email,
    PasswordHash: hashedPassword,
}
The DTO had: Password  |   The Model has: PasswordHash
This is why DTO ≠ Model.

Repository receives *model.User
Repository only knows: "I save Users."

Database stores -> ID, Name, Email, PasswordHash, CreatedAt, Response, Database Model

Service -> DTO -> RegisterResponse -> JSON {"id":1,"name":"Abhinandan","email":"abhi@gmail.com"}

# Dependency Injection
Instead of handler := AuthHandler{}
We have   main -> database -> repository -> service -> handler -> router

In code:
repo := repository.NewGormUserRepository(db)
service := service.NewAuthService(repo)
handler := handler.NewAuthHandler(service)

Everything depends on abstractions. Nothing creates its own dependencies.

# Handler
Receive JSON -> Bind DTO -> Call Service -> Return HTTP Response

# Service
Validate Name -> Normalize Email -> Validate Password -> Find Existing User -> Hash Password -> Create User Model -> Repository.Create() -> Return Response DTO

# Login is a core business operation. So want it readable

FindByEmail(email string)
FindByPhone(phone string)
FindByUsername(username string)
FindByEmployeeID(id string)

This is actually very readable.

| Specific Methods | Generic Method              |
| ---------------- | --------------------------- |
| Easy to read     | Flexible                    |
| Type-safe        | More reusable               |
| Self-documenting | Handles many search cases   |
| More methods     | More complex implementation |

It depends on the domain.

# Admin Search API
where users can search by: email, phone, city, department, status,,technician id, role
Then we introduce something like: SearchUsers(filter UserFilter)
where: type UserFilter struct {
    Email  *string
    Phone  *string
    Status *string
    Role   *string
}
This gives flexibility where it's needed, without making the authentication code harder to understand.

HTTP Request -> RegisterRequest DTO -> Handler -> Service -> User Model -> Repository -< Database -> Repository -> Service -> RegisterResponse DTO -> Handler -> JSON

Handler speaks: dto.RegisterRequest
Repository speaks: model.User
They never talk directly.


# Why do we use c.Set("userID", userID)?
{
    "user_id": 10,
    "email": "abhi@gmail.com",
    "exp": 1785678901
}
Middleware -> c.Set("userID", 10) -> Handler -> c.Get("userID") -> Service
After the JWT is validated, we already know the authenticated user's identity. We store it in the request's gin.Context so downstream handlers and services can access it without reparsing the token or making another database query. The context is request-scoped and exists only for the lifetime of that HTTP request.

# c.Abort()?
Stop processing this request. Do not execute any remaining middleware or handlers.

r.GET("/profile", middleware.AuthMiddleware(), handler.GetProfile,)

Request -> AuthMiddleware -> Handler

# ValidateJWT
Receive Token-> Parse JWT -> Verify Signature -> Verify Signing Method (HS256) -> Verify Expiration -> Extract Claims -> Return Claims

# token.Claims.(jwt.MapClaims)
Claims is an interface. An interface can hold many concrete types.

Calims Claims - > jwt.MapClaims or MyCustomClaims or AnotherClaims

To make it know it's specifically a jwt.MapClaims.

# request model should contain customerId or customer-Name, phoneNo.

Database Normalization
Specifically: Avoid redundancy, Maintain consistency, Single source of truth


# Customer + Request
(Customer Model, Request Model, Create Customer API, Create Request API, Get Request API)

Model -> Migration -> DTO -> Repository -> Service -> Handler -> Route -> Postman Test -> DONE

# Reservation
(Appointment Window, Reserve Window, Prevent Double Booking, Transactions)

# Dispatch
(Technician, Scheduler, Assignment, Concurrency)

# Historical record → Snapshot fields

# Request should contain ReservationID
1. Reservation model can contain appointment information (Separation of Responsibility)

Request -> What customer wants.
Reservation -> What appointment was booked.
AppointmentWindow -> What slots are available.

2. It helps while checking logs (Auditing)

3. (Future Changes)

Customer reschedules three times. Now Reservation table becomes:
| ReservationID | RequestID | Window | Status    |
| ------------- | --------- | ------ | --------- |
| 1             | 100       | 9-10   | Cancelled |
| 2             | 100       | 10-11  | Cancelled |
| 3             | 100       | 2-3    | Reserved  |

Everything belongs to the same request. (one-to-many relationship)
Request remains stable. Reservations become historical records.


POST   /customers
GET    /customers
GET    /customers/:id
PUT    /customers/:id

POST   /requests
GET    /requests
GET    /requests/:id
PUT    /requests/:id

No Reservation API yet. Because a request should exist before someone books an appointment.

Customer -> Create Request -> Reserve Appointment -> Dispatch -> Complete

Customer-Request

# Repository
Customer -> Create(), GetByID(), GetAll(), Update()
Request -> Create(), GetByID(), GetAll(), Update()


# Service 
Customer
CreateCustomer(), GetCustomer(), GetCustomers(), UpdateCustomer()

Request
CreateRequest(), GetRequest(), GetRequests(), UpdateRequest()

# Handler
POST -> Bind DTO -> Call Service -> Return JSON

# gorm:"size:100"
Tells GORM how to create the database schema. 
(i) Database column size (VARCHAR(100))
(ii) Prevents storing unnecessarily large strings.

# Why []Customer instead of []*Customer?

When we query:-  var customers []model.Customer;  db.Find(&customers)
GORM fills the slice. Returning -> []model.Customer   (is simple and efficient.)

Use []*Customer only when:
Objects are very large
You want multiple references to the same object
You need to mutate the same object from multiple places

For our CRUD APIs, []Customer -> is the idiomatic Go choice.


Rule 1: Every exported function needs a constructor.
Rule 2: Repository never returns DTO
Rule 3: Handler never talks to DB
Rule 4: Service never returns model
Rule 5: Validate -> Normalize -> Repository -> Business logic -> DTO -> Return

1. What does client send? -> 2. DTO -> 3. What business rules apply? -> 4. Service -> 5. What does DB need? -> 6. Model + Repository -> 7. What should client receive? -> 8. Response DTO