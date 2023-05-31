## How to test
    docker-compose up
    echo -e "your message" | sudo socat - UNIX:app.sock
