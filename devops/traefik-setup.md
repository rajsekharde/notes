# Traefik setup for reverse proxy inside docker

## HTTP routing:
```bash
# Minimum config:
docker-compose.yml:
services:
    traefik:
        image: traefik:v3.6
        command:
            # Enable docker provider. Traefik reads labels from docker
            - "--providers.docker=true"
            
            # Define an HTTP entrypoint on port 80
            - "--entrypoints.web.address=:80"
            
            # Enable traefik dashboard (insecure)
            - "--api.insecure=true"
            
            # Do not expose containers unless explicitly enabled
            - "--providers.docker.exposedbydefault=false"
        ports:
            # Expose port 80 on the host
            - "80:80"
            
            # Expose dashboard on port 8080
            - "8080:8080"
        volumes:
            # Allow Traefik to read Docker metadata (labels)
            - /var/run/docker.sock:/var/run/docker.sock
        networks:
            - app-network
    backend:
        labels:
            # Explicitly enable Traefik for this container
            - "traefik.enable=true"
            
            # Router listens on the web entrypoint (port 80)
            - "traefik.http.routers.backend.entrypoints=web"
            
            # Router: match requests starting with /api
            - "traefik.http.routers.backend.rule=PathPrefix(`/api`)"
            
            # Service: forward traffic to port 8000 inside the container
            - "traefik.http.services.backend.loadbalancer.server.port=8000"
            
            # Optional but recommended: higher priority than frontend
            - "traefik.http.routers.backend.priority=10"
    frontend:
        labels:
            - "traefik.enable=true"
            
            - "traefik.http.routers.frontend.entrypoints=web"
            
            # Router: catch-all for everything else
            - "traefik.http.routers.frontend.rule=PathPrefix(`/`)"
            
            # Service: forward traffic to port 80 inside the container. Replace with frontend port
            - "traefik.http.services.frontend.loadbalancer.server.port=80"
            
            # Lower priority than backend
            - "traefik.http.routers.frontend.priority=1"
networks:
    app-network:
        driver: bridge
```
