# Entectl - Ente Self-Hosted Manager

A CLI tool for managing self-hosted Ente photo storage instances using Docker Compose.

## Quick Start

### Option 1: Ansible Automation (Recommended)

```bash
# Install Ansible
pip install --user ansible

# Deploy Ente (Windows - use manual method due to locale issues)
go build -o entectl
go run main.go cluster init --config ansible/ansible-config.yaml --name ente-ansible
go run main.go cluster start --name ente-ansible

# Deploy Ente (Linux/Mac - full Ansible)
cd ansible
export PYTHONIOENCODING=utf-8
ansible-playbook -i inventory.ini playbook.yml --extra-vars "@vars.yml"

# Stop Ente
go run main.go cluster stop --name ente-ansible
# OR: ansible-playbook -i inventory.ini stop-ente.yml --extra-vars "@vars.yml"
```

### Option 2: Manual Deployment

```bash
# Build tool
go build -o entectl

# Initialize cluster
./entectl cluster init --config example-config.yaml --name my-ente

# Start services
./entectl cluster start --name my-ente

# Stop services
./entectl cluster stop --name my-ente
```

## Prerequisites

- Go 1.21+
- Docker Desktop
- Ansible (for automation)

## Configuration

Edit `ansible/vars.yml` or create your own config:

```yaml
domain: "localhost"
museum_port: 8090
db_password: "your_secure_password"
jwt_secret: "your_32_char_jwt_secret_here_123456789"
# ... more options
```

## Access Points

After deployment:
- **Photos**: http://localhost:3000
- **API**: http://localhost:8090
- **MinIO Console**: http://localhost:3201

## Services Deployed

- **Museum** - Ente API server
- **Web Apps** - Photo, accounts, albums interfaces
- **PostgreSQL** - Database
- **MinIO** - S3-compatible storage
- **Caddy** - Reverse proxy

## Management Commands

```bash
# List running services
./entectl cluster list --name my-ente

# View logs
./entectl cluster logs --name my-ente

# Remove cluster
./entectl cluster remove --name my-ente
```

## Ansible Benefits

- **One-command deployment**
- **Multi-environment support**
- **Configuration templating**
- **Idempotent operations**
- **Infrastructure as Code**

## Ansible Troubleshooting

**Windows Locale Issue:**
```bash
# If you get "Ansible requires UTF-8 encoding" error:
# Use the manual commands instead:
go run main.go cluster init --config ansible/ansible-config.yaml --name ente-ansible
go run main.go cluster start --name ente-ansible
```

**Linux/Mac:**
```bash
# Set encoding and run Ansible
export PYTHONIOENCODING=utf-8
ansible-playbook -i ansible/inventory.ini ansible/playbook.yml --extra-vars "@ansible/vars.yml"
```

**Check Status:**
```bash
# Verify deployment
go run main.go cluster list --name ente-ansible
docker compose ps
```
