<p align="center">
  <img src="docs/favicon.png" width="120px" alt="Union API Center Logo">
</p>

# Union API Center

Union API Center is a centralized service platform providing unified authentication and notification services for the OVINC ecosystem. It serves as an OpenID Connect (OIDC) provider and manages system-wide notifications.

## Features

- **Unified Authentication**: Acts as an OIDC Provider to handle single sign-on (SSO) for all OVINC services.
- **Notification System**: Centralized management for system notices and alerts.
- **Tencent Cloud Integration**: Built-in support for Tencent Cloud services.
- **Asynchronous Tasks**: Powered by Celery for background processing.
- **Real-time Communication**: Utilizes Django Channels for WebSocket support.

## Tech Stack

- **Framework**: Django, Django REST Framework
- **Asynchronous**: ASGI (Daphne/Uvicorn), Celery
- **Database**: MySQL
- **Cache/Message Broker**: Redis
- **Authentication**: django-oidc-provider

## Getting Started

### Prerequisites

- Python 3.8+
- Redis
- MySQL

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository_url>
   cd APICenter
   ```

2. **Install dependencies**
   ```bash
   pip install -r requirements.txt
   ```

3. **Configuration**
   Copy the example environment file and configure it:
   ```bash
   cp env.example .env
   ```
   Update `.env` with your database and Redis credentials.

4. **Database Migrations**
   ```bash
   python manage.py migrate
   ```

5. **Run the Server**
   ```bash
   python manage.py runserver
   ```

### Docker Support

The project includes a `Dockerfile` for containerized deployment.

```bash
docker build -t union-api-center .
```

## Development

- **Linting**: Run `make lint` to check code style.
- **Messages**: Run `make messages` to update translation files.

## License

This project is licensed under the terms specified in the [LICENSE](LICENSE) file.
