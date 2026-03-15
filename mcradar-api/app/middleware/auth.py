from http.client import UNAUTHORIZED
import logging
import os

from fastapi import Request, Response
from starlette.middleware.base import BaseHTTPMiddleware

logger = logging.getLogger(__name__)

API_KEY = os.environ.get("API_KEY")


class AuthMiddleware(BaseHTTPMiddleware):
    async def dispatch(self, request: Request, call_next):
        token = request.headers.get("Authorization")

        if token is None:
            return Response(status_code=UNAUTHORIZED)

        token = token.removeprefix("Bearer ")

        if token != API_KEY:
            return Response(status_code=UNAUTHORIZED)

        return await call_next(request)
