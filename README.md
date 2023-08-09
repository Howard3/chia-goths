# ChiA GoTHS - Chi Alpine.js, Go, Tailwind CSS, HTMX, and SQLite Template Project

![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)

## Introduction

ChiA GoTHS is a template project that combines various technologies to help you rapidly build modern web applications. It includes Go for server-side programming, HTMX and Alpine.js for seamless client-server communication and interactivity, Tailwind CSS for fast and responsive UI design, and SQLite for a lightweight and embeddable database solution.

This project serves as a foundation for your web development endeavors, allowing you to quickly get started and leverage the power of these technologies together.

## Features

- Go backend with Go-Chi for routing.
- HTMX and Alpine.js for smooth and efficient client-server interactions.
- Tailwind CSS for a utility-first, responsive UI design.
- SQLite for an easy-to-use and embedded database.
- Automatic CSRF protection using Gorilla CSRF
- Zero Allocation Logger with Zerolog

## Prerequisites

To run this project, you'll need the following installed on your system:

- Go (https://golang.org/)
- Node.js (https://nodejs.org/) for npm (Node Package Manager)
- Tailwind CSS (https://tailwindcss.com/docs/installation)
- SQLite (https://www.sqlite.org/download.html)
- Taskfile (https://taskfile.dev/installation/) optional but recommended.
- Flyctl (https://fly.io/docs/hands-on/install-flyctl/) if you want to deploy to [fly.io](https://fly.io/)

## Getting Started

1. Clone the repository:

```bash
git clone https://github.com/Howard3/chia-goths.git
```
2. Install frontend dependencies:
```bash
npm install
```
3. Start the dev server
```bash
task dev
```

## Configuration
Static Assets: Place your static files (images, CSS, JavaScript, etc.) in the assets directory.
Templates: This project uses gohtml templates, and they are stored in the templates directory.

## Contributing
Contributions are welcome! If you find any issues or have suggestions for improvement, please open an issue or submit a pull request.

## License
This project is licensed under the MIT License.
