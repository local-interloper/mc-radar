from pydantic import BaseModel, ConfigDict
from pydantic.alias_generators import to_camel
from pypika.queries import QueryBuilder


class Pagination(BaseModel):
    model_config = ConfigDict(alias_generator=to_camel, populate_by_name=True)

    first: int
    last: int

    def apply(self, query_builder: QueryBuilder) -> QueryBuilder:
        return query_builder.offset(self.first).limit(self.last - self.first)
