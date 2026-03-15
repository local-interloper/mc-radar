from typing import Annotated

from fastapi import Depends
from psycopg import Connection
import psycopg_pool
from os import environ

password = environ.get("POSTGRES_PASSWORD")
host = environ.get("POSTGRES_HOST", "127.0.0.1")
database = environ.get("POSTGRES_DB", "postgres")

pool = psycopg_pool.ConnectionPool(
    f"host={host} user=postgres password={password} sslmode=disable",
    min_size=5,
    max_size=50,
)


def get_db():
    with pool.connection() as connection:
        yield connection


DatabaseDep = Annotated[Connection, Depends(get_db)]
