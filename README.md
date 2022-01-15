# LibraryApp
Web application for my golang course.
____
## Installation
1. Clone this repository
2. Download and install MySQL (You can skip this step if already installed).
3. Create database with tables from `pkg/models/sql/librarydb.sql`
4. Configure config.env with your db and mail data (like an config.env.example)
5. Run `go build ./cmd/web` to build the executable file.
6. Run/execute generated program via ide or terminal `go run ./cmd/web`
7. Open your web browser, and navigate to `127.0.0.1:8000`
