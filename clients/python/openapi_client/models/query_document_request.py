# coding: utf-8

"""
    Studio - AI Empower Labs

    # Studio API Documentation  ## Introduction Welcome to Studio by AI Empower Labs API documentation! We are thrilled to offer developers around the world access to our cutting-edge artificial intelligence technology and semantic search. Our API is designed to empower your applications with state-of-the-art AI capabilities, including but not limited to natural language processing, audio transcription, embedding, and predictive analytics.  Our mission is to make AI technology accessible and easy to integrate, enabling you to enhance your applications, improve user experiences, and innovate in your field. Whether you're building complex systems, developing mobile apps, or creating web services, our API provides you with the tools you need to incorporate AI functionalities seamlessly.  Support and Feedback We are committed to providing exceptional support to our developer community. If you have any questions, encounter any issues, or have feedback on how we can improve our API, please don't hesitate to contact our support team @ support@AIEmpowerLabs.com.  Terms of Use Please review our terms of use and privacy policy before integrating our API into your application. By using our API, you agree to comply with these terms.

    The version of the OpenAPI document: v1
    Contact: support@aiempowerlabs.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import annotations
import pprint
import re  # noqa: F401
import json

from pydantic import BaseModel, ConfigDict, Field, StrictFloat, StrictInt, StrictStr
from typing import Any, ClassVar, Dict, List, Optional, Union
from openapi_client.models.document_filters import DocumentFilters
from typing import Optional, Set
from typing_extensions import Self

class QueryDocumentRequest(BaseModel):
    """
    QueryDocumentRequest
    """ # noqa: E501
    query: Optional[StrictStr] = Field(default=None, description="Semantic query to find matching documents")
    index: Optional[StrictStr] = Field(default=None, description="Optional index to specify which index to search in. Defaults to 'default'")
    filter: Optional[List[DocumentFilters]] = Field(default=None, description="Optional filtering of document id(s) and/or tags")
    min_relevance: Optional[Union[StrictFloat, StrictInt]] = Field(default=None, description="Optional filter to specify minimum relevance. Typically values between 0 and 1", alias="minRelevance")
    limit: Optional[StrictInt] = Field(default=None, description="Optional filter for specifying maximum number of entries to return. Defaults to 3")
    embedding_model: Optional[StrictStr] = Field(default=None, description="Embedding model to use in query", alias="embeddingModel")
    __properties: ClassVar[List[str]] = ["query", "index", "filter", "minRelevance", "limit", "embeddingModel"]

    model_config = ConfigDict(
        populate_by_name=True,
        validate_assignment=True,
        protected_namespaces=(),
    )


    def to_str(self) -> str:
        """Returns the string representation of the model using alias"""
        return pprint.pformat(self.model_dump(by_alias=True))

    def to_json(self) -> str:
        """Returns the JSON representation of the model using alias"""
        # TODO: pydantic v2: use .model_dump_json(by_alias=True, exclude_unset=True) instead
        return json.dumps(self.to_dict())

    @classmethod
    def from_json(cls, json_str: str) -> Optional[Self]:
        """Create an instance of QueryDocumentRequest from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self) -> Dict[str, Any]:
        """Return the dictionary representation of the model using alias.

        This has the following differences from calling pydantic's
        `self.model_dump(by_alias=True)`:

        * `None` is only added to the output dict for nullable fields that
          were set at model initialization. Other fields with value `None`
          are ignored.
        """
        excluded_fields: Set[str] = set([
        ])

        _dict = self.model_dump(
            by_alias=True,
            exclude=excluded_fields,
            exclude_none=True,
        )
        # override the default output from pydantic by calling `to_dict()` of each item in filter (list)
        _items = []
        if self.filter:
            for _item in self.filter:
                if _item:
                    _items.append(_item.to_dict())
            _dict['filter'] = _items
        # set to None if query (nullable) is None
        # and model_fields_set contains the field
        if self.query is None and "query" in self.model_fields_set:
            _dict['query'] = None

        # set to None if index (nullable) is None
        # and model_fields_set contains the field
        if self.index is None and "index" in self.model_fields_set:
            _dict['index'] = None

        # set to None if filter (nullable) is None
        # and model_fields_set contains the field
        if self.filter is None and "filter" in self.model_fields_set:
            _dict['filter'] = None

        # set to None if min_relevance (nullable) is None
        # and model_fields_set contains the field
        if self.min_relevance is None and "min_relevance" in self.model_fields_set:
            _dict['minRelevance'] = None

        # set to None if limit (nullable) is None
        # and model_fields_set contains the field
        if self.limit is None and "limit" in self.model_fields_set:
            _dict['limit'] = None

        # set to None if embedding_model (nullable) is None
        # and model_fields_set contains the field
        if self.embedding_model is None and "embedding_model" in self.model_fields_set:
            _dict['embeddingModel'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of QueryDocumentRequest from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "query": obj.get("query"),
            "index": obj.get("index"),
            "filter": [DocumentFilters.from_dict(_item) for _item in obj["filter"]] if obj.get("filter") is not None else None,
            "minRelevance": obj.get("minRelevance"),
            "limit": obj.get("limit"),
            "embeddingModel": obj.get("embeddingModel")
        })
        return _obj


