# Ente Ansible Automation

Automated deployment of Ente self-hosted photo storage using Ansible.

## Prerequisites

- Ansible installed
- Docker Desktop running
- Go installed

## Quick Start

```bash
# Deploy Ente (Ansible)
ansible-playbook -i inventory.ini playbook.yml --extra-vars "@vars.yml"

# Deploy Ente (Manual)
go build -o entectl
./entectl cluster init --config ansible-config.yaml --name ente-ansible
./entectl cluster start --name ente-ansible

# Stop Ente
ansible-playbook -i inventory.ini stop-ente.yml --extra-vars "@vars.yml"
# OR: ./entectl cluster stop --name ente-ansible
```

## Custom Configuration

Edit `vars.yml` to customize:
- Cluster name
- Domain
- Ports
- Security keys
- Database password

## Access Points

After deployment:
- Photos: http://localhost:3000
- API: http://localhost:8090
- MinIO: http://localhost:3201

## Files

- `playbook.yml` - Main deployment playbook
- `stop-ente.yml` - Stop services playbook
- `vars.yml` - Configuration variables
- `templates/ente-config.yaml.j2` - Config template
- `inventory.ini` - Ansible inventory