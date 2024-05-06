# Fly - Weather App

This folder shows examples of a weather app for deployment to Fly.

While the primary material is deployment to Fly, the weather app is good to be built together with students, so they understand how to fetch data from a public API and displays it.

This weather app connects to Open Weather Map (https://openweathermap.org/api). You need to sign up to get the API key, so make sure you do this before hand, it does takes some time before you get the API key.

The city is static, in this case London, you can update the code so it can receive from HTML form, a city which you can pass to the API.

The route handler is `/weather`, so from the local development, you can access it via: http://localhost:8080/weather and after deployment to Fly, remember to change localhost:8080 to the domain name received from Fly.
