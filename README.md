# GoChatter 
**A Minimalist, High-Performance Real-Time Chat App**  
Built with Go (backend) and Vanilla JavaScript (frontend)  

## Features  
- **Real-time messaging** using raw WebSockets (`gorilla/websocket`).  
- **Zero dependencies** on frontend (no React, just DOM APIs).  
- **Binary protocol support** (faster than JSON).  
- **Concurrent-safe broadcasting** in Go (`sync.Mutex`).  
- **Single-binary deployment** (just `go build`).  

##  Tech Stack  
| Part          | Tech Used                     | Why?                                
|---------------|-------------------------------|---------------------------- 
| **Backend**   | Go (`net/http`, WebSockets)   | Raw performance, simplicity        
| **Frontend**  | Vanilla JS + WebSocket API    | No framework overhead              
| **Protocol**  | Custom binary format          | Low latency, small payloads       

## 🛠️ Installation  
### Prerequisites  
- Go 1.20+  
- Node.js (for `wscat` testing)  

### Steps  
1. Clone the repo:  
   ```bash  
   git clone https://github.com/hanapiko/GoChatter.git  
   cd GoChatter  

2. Run the Go backend:
```bash
cd backend
go run main.go
```

3 Open the frontend:
```bash
# Just open frontend/index.html in a browser!
```