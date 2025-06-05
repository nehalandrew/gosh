# This script is used to test the gosh installation in a Docker container.

docker run -id --rm -p 2222:22 docker.io/debian bash <<EOF
#!/bin/bash
set -e

# Update and install dependencies
apt update && apt install -y wget

# Download and install gosh package
wget https://github.com/nehalandrew/gosh/releases/download/v0.0.1/gosh.deb
apt install -y ./gosh.deb

# SSH setup
mkdir -p /var/run/sshd
mkdir -p /home/gosh/.ssh
chown -R gosh:gosh /home/gosh
chown -R gosh:gosh /home/gosh/.ssh

# Set gosh user password to empty (no password)
echo "gosh:" | chpasswd -e || echo "gosh:" | chpasswd

# Modify sshd_config to allow password authentication and empty passwords
sed -i 's/#PasswordAuthentication yes/PasswordAuthentication yes/' /etc/ssh/sshd_config
sed -i 's/#PermitEmptyPasswords no/PermitEmptyPasswords yes/' /etc/ssh/sshd_config
sed -i 's/#PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config
sed -i 's/ChallengeResponseAuthentication yes/ChallengeResponseAuthentication no/' /etc/ssh/sshd_config

# Start ssh service
service ssh start

# Keep container alive so ssh server keeps running
tail -f /dev/null
EOF

# Test ssh login with no password prompt
sshpass -p "" ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null gosh@localhost -p 2222
