const socket = new WebSocket("ws://localhost:8080/ws");
const chatDiv = document.getElementById("chat");
const msgInput = document.getElementById("msg");

// Receive messages
socket.onmessage = (event) => {
    const msg = document.createElement("div");
    msg.textContent = event.data;
    chatDiv.appendChild(msg);
    chatDiv.scrollTop = chatDiv.scrollHeight; // Auto-scroll
};

// Send messages
document.getElementById("send").addEventListener("click", () => {
    if (msgInput.value.trim()) {
        socket.send(msgInput.value);
        msgInput.value = "";
    }
});

// Send on Enter key
msgInput.addEventListener("keypress", (e) => {
    if (e.key === "Enter") {
        socket.send(msgInput.value);
        msgInput.value = "";
    }
});