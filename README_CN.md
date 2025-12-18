<div align="center">
  <img src="docs/favicon.png" width="120" alt="Union API Center Logo">
  <h1>Union API Center</h1>
  <p>
    <strong>OVINC 生态系统的核心身份认证与通知服务平台</strong>
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

## 📖 简介

**Union API Center** 是 OVINC 生态系统的基础设施核心。它作为一个强大的 **OpenID Connect (OIDC) 提供商**，实现了所有 OVINC 服务之间的单点登录 (SSO)。此外，它还管理着一个集中的通知系统，处理全系统的公告和消息推送。

无论您是构建需要身份验证的微服务，还是需要实时通知的前端应用，Union API Center 都能提供必要的后端支持。

## ✨ 核心特性

*   🔐 **统一身份认证**: 符合 OIDC 标准，支持无缝 SSO 集成。
*   📢 **通知系统**: 集中管理的系统公告、警报和消息服务。
*   ☁️ **腾讯云集成**: 内置对腾讯云服务（短信、对象存储等）的支持。
*   ⚡ **异步任务处理**: 基于 **Celery** 的高性能后台任务处理。
*   🔄 **实时通信**: 由 **Django Channels** 驱动的 WebSocket 支持。
*   🌍 **国际化**: 支持多语言 (i18n)。

## 🛠 技术栈

*   **框架**: Django, Django REST Framework
*   **异步支持**: ASGI (Daphne/Uvicorn), Celery
*   **数据库**: MySQL
*   **缓存/消息队列**: Redis
*   **认证协议**: OpenID Connect (基于 `django-oidc-provider`)

## 📂 项目结构

```text
APICenter/
├── apps/            # 业务模块
├── bin/             # 二进制脚本
├── core/            # 核心配置
├── docs/            # 文档与资源
├── entry/           # 入口文件 (WSGI/ASGI)
├── locale/          # 翻译文件
├── scripts/         # 工具脚本
├── Dockerfile       # Docker 构建配置
├── Makefile         # 快捷命令
└── manage.py        # Django 管理脚本
```

## 🚀 快速开始

### 前置要求

*   Python 3.8+
*   Redis
*   MySQL

### 安装步骤

1.  **克隆仓库**

    ```bash
    git clone https://github.com/OVINC/APICenter.git
    cd APICenter
    ```

2.  **安装依赖**

    ```bash
    pip install -r requirements.txt
    ```

3.  **配置环境**

    复制示例配置文件：

    ```bash
    cp env.example .env
    ```

    > **注意**: 请在 `.env` 文件中更新您的数据库凭据、Redis 地址和其他密钥。

4.  **数据库迁移**

    ```bash
    python manage.py migrate
    ```

5.  **启动服务器**

    ```bash
    python manage.py runserver
    ```

## 🐳 Docker 支持

使用 Docker 构建并运行应用：

```bash
# 构建镜像
docker build -t union-api-center .

# 运行容器
docker run -d -p 8000:8000 --env-file .env union-api-center
```

## 💻 开发指南

我们提供了 `Makefile` 来简化常见的开发任务：

*   **代码检查**: 运行代码风格检查
    ```bash
    make lint
    ```
*   **国际化**: 更新翻译文件
    ```bash
    make messages
    ```

## 🤝 贡献

欢迎提交 Pull Request 来参与贡献！

## 📄 许可证

本项目遵循 [MIT License](LICENSE) 许可证。
