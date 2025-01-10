#!/bin/bash

case "$1" in
    "start")
        sudo docker compose up -d
        ;;
    "stop")
        sudo docker compose down
        ;;
    "restart")
        sudo docker compose restart
        ;;
    "logs")
        sudo docker compose logs -f
        ;;
    "status")
        sudo docker compose ps
        echo "---"
        echo "Memory Usage:"
        sudo docker stats --no-stream
        ;;
    "backup")
        timestamp=$(date +%Y%m%d_%H%M%S)
        backup_dir="backups/$timestamp"
        mkdir -p "$backup_dir"
        sudo docker compose exec postgis pg_dump -U gisuser gisdb > "$backup_dir/gisdb.sql"
        cp -r data/* "$backup_dir/"
        echo "Backup created in $backup_dir"
        ;;
    *)
        echo "Usage: $0 {start|stop|restart|logs|status|backup}"
        exit 1
        ;;
esac
