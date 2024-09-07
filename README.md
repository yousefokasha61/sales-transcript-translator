# Sales Transcript Translator

## Overview

This project is a Go-based application that serves as a Sales Transcript Translator. It is containerized using Docker
and managed with Docker Compose. The application is designed to be easy to set up and run in a local development
environment.

## Prerequisites

Before you can run this project, make sure you have the following installed on your machine:

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/sales-transcript-translator.git
cd sales-transcript-translator
```

### 2. Set the `OPENAI_API_KEY` Environment Variable in `.env` file

```bash
# .env
OPENAI_API_KEY=your_openai_api_key_here
```

### 3. Build the Docker Image

```bash
docker-compose build
```

### 3. Start the Application

```bash
docker-compose up
```

### 4. Access the Application

```bash
http://localhost:8080
```
