const output = document.getElementById("cpu-temperature");

// instance of returns a boolean
// in case the output is null
if (output instanceof HTMLElement) {
	const socket = new WebSocket("ws://localhost:8085");

	socket.onmessage = (event: MessageEvent<string>) => {
		output.textContent += `${event.data}\n`;
	};

	socket.onerror = (err: Event) => {
		console.error("WebSocket error:", err);
		output.textContent += "WebSocket error\n";
	};
} else {
	console.error("Element with ID 'cpu-temperature' not found or is not an HTMLElement");
}
