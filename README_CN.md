<p align="center">
  <img src="docs/favicon.png" width="120px" alt="Union API Center Logo">
</p>

# Union API Center

Union API Center 是 OVINC 生态系统的核心服务平台，提供统一的身份认证和通知服务。它作为 OpenID Connect (OIDC) 提供商，管理整个系统的单点登录（SSO）及消息通知。

## 功能特性

- **统一身份认证**: 作为 OIDC 提供商，为所有 OVINC 服务提供单点登录支持。
- **通知系统**: 集中管理的系统通知与公告服务。
- **腾讯云集成**: 内置腾讯云服务支持。
- **异步任务**: 基于 Celery 的后台任务处理。
- **实时通信**: 使用 Django Channels 提供 WebSocket 支持。

## 技术栈

- **框架**: Django, Django REST Framework
- **异步支持**: ASGI (Daphne/Uvicorn), Celery
- **数据库**: MySQL
- **缓存/消息队列**: Redis
- **认证**: django-oidc-provider

## 快速开始

### 前置要求

- Python 3.8+
- Redis
- MySQL

### 安装步骤

1. **克隆仓库**
   ```bash
   git clone <repository_url>
   cd APICenter
   ```

2. **安装依赖**
   ```bash
   pip install -r requirements.txt
   ```

3. **配置环境**
   复制示例配置文件并进行修改：
   ```bash
   cp env.example .env
   ```
   请在 `.env` 文件中更新您的数据库和 Redis 配置。

4. **数据库迁移**
   ```bash
   python manage.py migrate
   ```

5. **启动服务器**
   ```bash
   python manage.py runserver
   ```

### Docker 支持

本项目包含 `Dockerfile`，支持容器化部署。

```bash
docker build -t union-api-center .
```

## 开发指南

- **代码检查**: 运行 `make lint` 进行代码风格检查。
- **国际化**: 运行 `make messages` 更新翻译文件。

## 许可证

本项目遵循 [LICENSE](LICENSE) 文件中声明的许可条款。
