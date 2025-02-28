# (CNF-Q) CNF Queue Service

## Description

Queue Service is a fast and lightweight queuing system written in Go. It allows adding and retrieving data from queues via a REST API. It supports token-based authentication and handles multiple queues simultaneously.

## 🚀 Features

- Create and manage queues
- Push and pop messages (FIFO)
- Peek at the last element in the queue
- List available queues
- Token authentication

## 📌 Installation

### Option 1: Build your binary file
Please clone repository.
```
git clone https://github.com/cnfsoftwarw/cnf-q.git
```
and compile the file.
```
make build
```
The compiled version of file should be available at `./bin/cnf-q-service`

### Option 2: Run docker image
To run docker image please run the following command.
```
docker run -p 8080:8080 cnfsoftware/cnf-q-server
```

## 📡 API Endpoints

---

### 🔹 Push an item to a queue
```
POST /queue/:name/push
```
Body: Any binary data.

Example cURL:
```
curl -X POST http://localhost:8080/queue/myqueue/push \
-H "X-Auth-Token: valid-token" \
--data "Hello, Queue!"
```

### 🔹 Pop an item from a queue
```
POST /queue/:name/pop
```

Example cURL:
```
curl -X POST http://localhost:8080/queue/myqueue/pop -H "X-Auth-Token: Bearer valid-token"
```

### 🔹 Peek at the last element in a queue (without removing it)
```
GET /queue/:name/peek
```

Example cURL:
```
curl -X POST http://localhost:8080/queue/myqueue/pop -H "X-Auth-Token: valid-token"
```

### 🔹 Retrieve a list of queues
```
GET /queues
```

Example cURL:
```
curl -X POST http://localhost:8080/queue/myqueue/pop -H "X-Auth-Token: valid-token"
```

Example response:

```
{
    "queues": [
        "queue-1",
        "queue-3",
        "queue-2"
    ]
}

```

## 📜 License

This project is released under the MIT license.

📧 For questions or issues, feel free to open an Issue on GitHub! 🚀

## 📜 Maintenance

The maintenance of the project belongs to [CNF Software](https://cnf.software).
