# Railway - Basic API server

This folder shows examples of a basic API server for deployment to Railway.

While the primary material is deployment to Railway, the basic API server is good to be built together with students, so they understand how to create APIs in Go and how to deploy it on Railway.

This server starts on port 8080 and listens for requests at the / endpoint. You can test it by navigating to http://localhost:8080/ in your web browser after starting the server. Naturally when hosted on Railway, the localhost:8080 need to be changed.

It connects to https://zenquotes.io/ and output a single quote onto a HTML page on the / endpoint. Or you can use the /random endpoint that outputs JSON.

Feel free to add more functionalities to the basic API server.
