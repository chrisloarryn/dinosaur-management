<h1>Dinosaur Management API</h1>
<p>This repository contains an API for managing dinosaurs using the hexagonal architecture.</p>

<h2>Project Structure</h2>
    <p>The project is organized into different directories:</p>
    <ul>
        <li><code>cmd</code>: Main entry point of the application.</li>
        <li><code>internal</code>: Contains core application logic.</li>
        <li><code>domain</code>: Defines the core domain model.</li>
        <li><code>usecase</code>: Implements business use cases.</li>
        <li><code>repository</code>: Defines repository interfaces.</li>
        <li><code>infrastructure</code>: Provides data storage implementations.</li>
        <li><code>adapter</code>: Adapts use cases to the API layer.</li>
        <li><code>api</code>: Implements the API server using the Echo framework.</li>
    </ul>

<h2>Getting Started</h2>
    <p>To run the API server:</p>
    <pre><code>go run cmd/main.go</code></pre>

<h2>API Endpoints</h2>
    <ul>
        <li><strong>GET /dinosaurs</strong>: List all dinosaurs.</li>
        <li><strong>POST /dinosaurs</strong>: Create a new dinosaur.</li>
    </ul>

<h2>Dependencies</h2>
    <p>This project uses the following libraries:</p>
    <ul>
        <li><a href="https://echo.labstack.com/">Echo Framework</a> for building the API server.</li>
    </ul>

<h2>License</h2>
    <p>This project is licensed under the MIT License.</p>