const fs = require("fs");
const WebSocket = require("ws");

// Path to your downloaded WAV file
const wavFilePath = "C:\\Users\\HP\\Downloads\\file_example_WAV_1MG.wav"; // Update this with the exact file name
// Replace 'your_file_name.wav' with the actual name of your WAV file

// Read the WAV file as binary data
const wavData = fs.readFileSync(wavFilePath);

// Connect to the WebSocket server
const ws = new WebSocket("ws://localhost:8080/convert");

ws.on("open", () => {
    console.log("Connected to WebSocket server");

    // Send the WAV data to the server
    ws.send(wavData);
});

ws.on("message", (data) => {
    console.log("Received FLAC data from server");

    // Save the FLAC data to a file
    fs.writeFileSync("C:\\Users\\HP\\Downloads\\output.flac", data);
    console.log("FLAC data saved to output.flac");

    // Close the connection
    ws.close();
});

ws.on("close", () => {
    console.log("Disconnected from WebSocket server");
});

ws.on("error", (error) => {
    console.error("WebSocket error:", error);
});
