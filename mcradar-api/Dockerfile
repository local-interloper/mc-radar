FROM ghcr.io/astral-sh/uv:alpine3.23

WORKDIR /app

COPY . .

RUN uv sync

CMD ["uv", "run", "fastapi", "run", "app/main.py"]