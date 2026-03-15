from typing import Generic, TypeVar, List

from pydantic import BaseModel, ConfigDict
from pydantic.alias_generators import to_camel


T = TypeVar("T")


class PaginatedDataResponse(BaseModel, Generic[T]):
    model_config = ConfigDict(alias_generator=to_camel, populate_by_name=True)
    data: List[T]
    total: int
