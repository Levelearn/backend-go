# Levelearn Backend

This repository contains the backend code for a gamified learning application, migrated from Express.js to **Go (Golang)**. It provides a high-performance RESTful API for managing users, courses, progress, rewards, and other game-related functionalities.

## Table of Contents
* [Introduction](#introduction)
* [Technologies Used](#technologies-used)
* [Features](#features)
* [Project Structure](#project-structure)
* [Installation](#installation)
* [Running the Application](#running-the-application)
* [API Documentation](#api-documentation)

---

## Introduction
This backend application serves as the core engine for a gamified learning platform. It handles data persistence, user authentication (JWT), authorization, and the business logic for awarding points, badges, and rewards based on user progress. The migration to Go was performed to leverage better concurrency, type safety, and memory efficiency.

## Technologies Used
* **Go (Golang)**: Core programming language.
* **Gin Gonic**: High-performance HTTP web framework.
* **GORM**: The developer-friendly ORM for Go (Replacement for Prisma).
* **MySQL**: Relational database management system.
* **Godotenv**: For managing environment variables.
* **Golang-migrate**: (Optional) For database migrations.

## Features
* **User Management**: Secure authentication and profile handling.
* **Course Management**: Full CRUD for learning materials.
* **Chapter & Material**: Hierarchical content structure.
* **Assessment & Assignment**: Evaluation systems and submission tracking.
* **Gamification Logic**: Real-time point calculation, leveling, and rewards.

## Project Structure
This project follows a clean architecture pattern tailored for Go:
```text
.
├── cmd/            # Application entry point (main.go)
├── config/         # Database and environment configurations
├── controllers/    # Request handlers and logic
├── models/         # Database schemas and GORM structs
├── repositories/   # Data access layer
├── routes/         # API endpoint definitions
├── utils/          # Helpers (JWT, Password Hashing, etc.)
├── .env            # Environment variables
└── go.mod          # Go module dependencies
