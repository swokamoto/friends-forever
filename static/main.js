document.addEventListener("DOMContentLoaded", function() {
    const responseDiv = document.getElementById("response");

    const eventSource = new EventSource("/events");
    eventSource.onmessage = function(event) {
        const newMessage = document.createElement("div");
        newMessage.textContent = event.data;
        responseDiv.appendChild(newMessage);
    };

    document.getElementById("helloButton").addEventListener("click", function() {
        fetch("/hello")
            .then(response => response.text())
            .then(data => {
                const helloMessage = document.createElement("div");
                helloMessage.innerHTML = data;
                responseDiv.appendChild(helloMessage);
            });
    });
});