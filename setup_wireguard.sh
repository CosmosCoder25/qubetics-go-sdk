#!/bin/bash
set -e  # Exit on error

# Configuration
WG_DIR="/etc/wireguard"
WG_CONF="$WG_DIR/wg0.conf"
WG_INTERFACE="wg0"
WG_NETWORK="10.8.0.1/24"
WG_PORT="51820"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to check if running as root
check_root() {
    if [ "$(id -u)" -ne 0 ]; then
        echo -e "${RED}Error: This script must be run as root${NC}" >&2
        exit 1
    fi
}

# Function to generate keys
generate_keys() {
    echo -e "${YELLOW}Generating WireGuard keys...${NC}"
    
    # Create directory if it doesn't exist
    mkdir -p "$WG_DIR"
    chmod 700 "$WG_DIR"
    
    # Generate private key
    if [ ! -f "$WG_DIR/private.key" ]; then
        wg genkey | tee "$WG_DIR/private.key" | wg pubkey > "$WG_DIR/public.key"
        chmod 600 "$WG_DIR/private.key"
        chmod 644 "$WG_DIR/public.key"
    else
        echo -e "${YELLOW}Private key already exists, skipping generation${NC}"
    fi
    
    # Generate pre-shared key if it doesn't exist
    if [ ! -f "$WG_DIR/preshared.key" ]; then
        wg genpsk > "$WG_DIR/preshared.key"
        chmod 600 "$WG_DIR/preshared.key"
    fi
}

# Function to create WireGuard config
create_wg_config() {
    echo -e "${YELLOW}Creating WireGuard configuration...${NC}"
    
    # Check if config already exists
    if [ -f "$WG_CONF" ]; then
        echo -e "${YELLOW}WireGuard configuration already exists, backing up to $WG_CONF.bak${NC}"
        cp "$WG_CONF" "${WG_CONF}.bak"
    fi
    
    # Get the primary network interface
    PRIMARY_IFACE=$(ip route | grep '^default' | awk '{print $5}' | head -n 1)
    
    # Create new config
    cat > "$WG_CONF" <<EOL
[Interface]
PrivateKey = $(cat "$WG_DIR/private.key")
Address = $WG_NETWORK
ListenPort = $WG_PORT
SaveConfig = true
PostUp = iptables -A FORWARD -i $WG_INTERFACE -j ACCEPT; iptables -t nat -A POSTROUTING -o $PRIMARY_IFACE -j MASQUERADE
PostDown = iptables -D FORWARD -i $WG_INTERFACE -j ACCEPT; iptables -t nat -D POSTROUTING -o $PRIMARY_IFACE -j MASQUERADE
EOL

    chmod 600 "$WG_CONF"
}

# Function to display setup information
show_setup_info() {
    echo -e "\n${GREEN}WireGuard setup complete!${NC}"
    echo -e "Public key: ${YELLOW}$(cat "$WG_DIR/public.key")${NC}"
    echo -e "Interface: ${WG_INTERFACE}"
    echo -e "Network: ${WG_NETWORK}"
    echo -e "Port: ${WG_PORT}"
    echo -e "\nTo start WireGuard: ${YELLOW}systemctl start wg-quick@$WG_INTERFACE${NC}"
    echo -e "To enable on boot: ${YELLOW}systemctl enable wg-quick@$WG_INTERFACE${NC}"
}

# Main execution
main() {
    check_root
    generate_keys
    create_wg_config
    show_setup_info
}

main "$@"