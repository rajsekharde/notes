## Cloudflare proxy turned ON

Request path:
Client - Cloudflare (HTTPS)
Cloudflare - Traefik(Server) (HTTP)

Cloudflare responsibilities:
- TLS termination
- DDOS protection
- WAF (firewall)
- Bot protection
- Rate limiting
- Hide server IP

Cloudflare forwards traffic to server in HTTP by default
Cloudflare - server connection can be made HTTPS using Cloudflare Origin Certificates

## Cloudflare proxy turned OFF

Request path:
Client - Traefik(Server) (HTTPS)

Server:
- TLS termination
- LetsEncrypt certs (or your own certs)
- Full end-to-end encryption
- No Cloudflare protection

## Basic deployment setup for projects:
Add HTTP routing with Traefik with hostname included in container labels
Enable Cloudflare proxy in the DNS record
Set SSL/TLS as Flexible on Cloudflare
Request path:
Client - Cloudflare (HTTPS)
Cloudflare - Server (HTTP)
