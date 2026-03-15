from fastapi import FastAPI, APIRouter
from routers import servers_router
import middleware


app = FastAPI(docs_url=None, redoc_url=None, openapi_url=None)

app.add_middleware(middleware.AuthMiddleware)

core_router = APIRouter(prefix="/api")

core_router.include_router(servers_router)

app.include_router(core_router)

