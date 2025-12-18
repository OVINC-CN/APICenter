<div align="center">
  <img src="docs/favicon.png" width="120" alt="Union API Center Logo">
  <h1>Union API Center</h1>
  <p>
    <strong>Centralized Authentication & Notification Service for OVINC Ecosystem</strong>
  </p>

  <p>
    <a href="LICENSE">
      <img src="https://img.shields.io/badge/License-MIT-green.svg" alt="License">
    </a>
    <img src="https://img.shields.io/badge/Python-3.12+-blue.svg" alt="Python">
    <img src="https://img.shields.io/badge/Django-5.x-green.svg" alt="Django">
  </p>

  <p>
    <a href="README.md">English</a> | <a href="README_CN.md">简体中文</a>
  </p>
</div>

---

## 📖 Introduction

**Union API Center** is the core infrastructure of the OVINC ecosystem. It serves as a robust **OpenID Connect (OIDC) Provider**, enabling Single Sign-On (SSO) across all OVINC services. Additionally, it manages a centralized notification system, handling system-wide alerts and messaging.

Whether you are building a microservice that needs authentication or a frontend app requiring real-time notifications, Union API Center provides the necessary backend support.

## ✨ Key Features

*   🔐 **Unified Authentication**: OIDC-compliant provider for seamless SSO integration.
*   📢 **Notification System**: Centralized management for system notices, alerts, and announcements.
*   ☁️ **Tencent Cloud Integration**: Native support for Tencent Cloud services (SMS, OSS, etc.).
*   ⚡ **Asynchronous Processing**: High-performance background task handling using **Celery**.
*   🔄 **Real-time Communication**: WebSocket support powered by **Django Channels**.
*   🌍 **Internationalization**: Multi-language support (i18n) ready.

## 🛠 Tech Stack

*   **Framework**: Django, Django REST Framework
*   **Asynchronous**: ASGI (Daphne/Uvicorn), Celery
*   **Database**: MySQL
*   **Cache & Broker**: Redis
*   **Auth Protocol**: OpenID Connect (via `django-oidc-provider`)

## 📂 Project Structure

```text
APICenter/
├── apps/            # Application modules
├── bin/             # Binary scripts
├── core/            # Core configuration and settings
├── docs/            # Documentation and assets
├── entry/           # Entry points (WSGI/ASGI)
├── locale/          # Translation files
├── scripts/         # Utility scripts
├── Dockerfile       # Docker build configuration
├── Makefile         # Command shortcuts
└── manage.py        # Django management script
```

## 🚀 Getting Started

### Prerequisites

*   Python 3.8+
*   Redis
*   MySQL

### Installation

1.  **Clone the repository**

    ```bash
    git clone https://github.com/OVINC/APICenter.git
    cd APICenter
    ```

2.  **Install dependencies**

    ```bash
    pip install -r requirements.txt
    ```

3.  **Configuration**

    Copy the example environment file:

    ```bash
    cp env.example .env
    ```

    > **Note**: Update `.env` with your database credentials, Redis URL, and other secrets.

4.  **Database Migrations**

    ```bash
    python manage.py migrate
    ```

5.  **Run the Server**

    ```bash
    python manage.py runserver
    ```

## 🐳 Docker Support

Build and run the application using Docker:

```bash
# Build the image
docker build -t union-api-center .

# Run the container
docker run -d -p 8000:8000 --env-file .env union-api-center
```

## 💻 Development

We provide a `Makefile` to simplify common development tasks:

*   **Linting**: Check code style
    ```bash
    make lint
    ```
*   **Translation**: Update translation files
    ```bash
    make messages
    ```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

This project is licensed under the [MIT License](LICENSE).
