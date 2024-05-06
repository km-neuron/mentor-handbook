# Fly Basic Echo Server

This folder shows examples of a basic echo server for deployment to Fly.

While the primary material is deployment to Fly, the basic echo server is good to be built together with students, so they understand basic client-server communication.

This server starts on port 8080 and listens for requests at the /echo endpoint. It expects a query parameter named message and echoes it back in the response. You can test it by navigating to http://localhost:8080/echo?message=Hello in your web browser after starting the server. Naturally when hosted on Fly, the localhost:8080 need to be changed.

Feel free to add more functionalities to the basic echo server. And maybe updating the server to be something beyond basic echo server.
