# E-commerce API

A backend API for an e-commerce platform built with Go. This project provides essential features for running an online store, including user authentication, product management, and order processing.

## ✨ Features

*   **User Authentication:** Secure user registration and login using JSON Web Tokens (JWT).
*   **Admin Seeding:** Automatic creation of an admin user on startup from environment variables.
*   **RESTful API:** A clean and organized API structure.
*   **Configuration-driven:** Easy setup using environment variables.
*   **Email Integration:** Ready for integrating with an email service for notifications.

## 🛠️ Technologies Used

*   [Go](https://golang.org/)
*   [PostgreSQL](https://www.postgresql.org/)
*   JWT for Authentication

## 🚀 Getting Started

Follow these instructions to get a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

*   [Go](https://golang.org/dl/) (version 1.18 or higher)
*   [PostgreSQL](https://www.postgresql.org/download/)
*   [Git](https://git-scm.com/)

### Installation

1.  **Clone the repository:**
    ```sh
    git clone https://github.com/Asura409/Ecommerce-api.git
    cd Ecommerce-api
    ```

2.  **Set up environment variables:**
    Create a `.env` file in the root of the project by copying the example file.
    ```sh
    cp .env.example .env
    ```
    Now, open the `.env` file and fill in your specific configuration details, such as your database credentials and a strong JWT secret.

3.  **Install dependencies:**
    This command will download the necessary Go modules.
    ```sh
    go mod tidy
    ```

4.  **Set up the database:**
    Ensure your PostgreSQL server is running. Connect to it and create the database you specified in the `DATABASE_DSN` variable in your `.env` file.

5.  **Run the application:**
    ```sh
    go run main.go
    ```
    The API server should now be running on the port specified in your `.env` file (default is `3001`).

## ⚙️ Configuration

The application is configured using environment variables stored in a `.env` file.

| Variable         | Description                                                              | Example                                                              |
| ---------------- | ------------------------------------------------------------------------ | -------------------------------------------------------------------- |
| `APP_PORT`       | The port on which the application server will listen.                    | `3001`                                                               |
| `DATABASE_DSN`   | The Data Source Name for connecting to the PostgreSQL database.          | `"host=localhost user=blog password=asura dbname=bp port=5432"`      |
| `JWT_SECRET`     | A secret key used to sign and verify JWTs.                               | `"your-super-secret-key"`                                            |
| `ADMIN_EMAIL`    | The email for the default admin user, created on startup.                | `"admin@example.com"`                                                |
| `ADMIN_PASSWORD` | The password for the default admin user.                                 | `"admin123"`                                                         |
| `EMAIL_API_KEY`  | The API key for your chosen email service provider.                      | `"your-email-api-key"`                                               |
| `EMAIL_SENDER`   | The email address from which transactional emails will be sent.          | `"noreply@yourdomain.com"`                                           |

## 🤝 Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1.  Fork the Project
2.  Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3.  Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4.  Push to the Branch (`git push origin feature/AmazingFeature`)
5.  Open a Pull Request