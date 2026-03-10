
```markdown
# gRPC API Contracts & Versioning (Go)

This project demonstrates **API Contracts and Versioning using gRPC and Golang**.

It shows how to maintain **backward compatibility** while evolving APIs by introducing **versioned protobuf contracts (v1, v2)**.

The project contains two API versions:

- **v1** → Basic user information
- **v2** → Extended user information (adds email field)

This allows **old clients to continue working while new clients use enhanced APIs**.

---

# Architecture


Client (v1 or v2)
│
▼
gRPC Server
│
▼
Versioned API Contracts
(Proto v1 / Proto v2)



---

# Project Structure



grpc-api-versioning
│
├── go.mod
│
├── proto
│   ├── v1
│   │   └── user.proto
│   │
│   └── v2
│       └── user.proto
│
├── pb
│   ├── v1
│   └── v2
│
├── server
│   └── main.go
│
└── client
├── client_v1.go
└── client_v2.go


---

# API Contract

In gRPC, the **API contract is defined using Protocol Buffers (.proto)**.

Example:


service UserService {
rpc GetUser (UserRequest) returns (UserResponse);
}



This defines the **request and response structure for communication**.

---

# Version 1 API

`proto/v1/user.proto`

Features:

- Returns basic user details
- Fields:
  - id
  - name

Example Response



{
id: 1
name: "Chirag"
}




# Version 2 API

`proto/v2/user.proto`

New field added:

- email

Example Response


{
id: 1
name: "Chirag"
email: "[chirag@email.com](mailto:chirag@email.com)"
}



Older clients using **v1 continue to work without changes**.

---

# Generate gRPC Code

Run the following command to generate Go code from proto files.



protoc --go_out=. --go-grpc_out=. proto/v1/user.proto
protoc --go_out=. --go-grpc_out=. proto/v2/user.proto


# Run the Server



go run ./server



Server will start on:



localhost:50051

---

# Run Client (Version 1)



go run ./client/client_v1.go



Output


User: id:1 name:"Chirag"



# Run Client (Version 2)



go run ./client/client_v2.go

Output



User: id:1 name:"Chirag" email:"[chirag@email.com](mailto:chirag@email.com)"



# Key Concepts Demonstrated

- API Contracts using Protocol Buffers
- API Versioning Strategy
- Backward Compatibility
- gRPC Service Implementation
- Versioned Client Consumption

---

# Best Practices for gRPC Versioning

1. Never remove existing fields
2. Only add new fields
3. Never reuse field numbers
4. Use versioned packages (`v1`, `v2`)
5. Maintain backward compatibility

---

# Technologies Used

- Golang
- gRPC
- Protocol Buffers

---

# Use Cases

This pattern is commonly used in:

- Microservices architectures
- Fintech APIs
- Large-scale distributed systems
- Public API lifecycle management

---

# Author

Chirag Raul

GitHub: https://github.com/Chirag711
```
