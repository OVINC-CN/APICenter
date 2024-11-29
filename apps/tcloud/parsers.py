class TCICallbackParser:
    """
    callback parser
    """

    def __init__(self, data: dict):
        self.data = data

    @property
    def is_sensitive(self) -> bool:
        return self.data["JobsDetail"]["Result"] != 0

    @property
    def audit_id(self) -> str:
        return self.data["JobsDetail"]["JobId"]

    @property
    def detail(self) -> dict:
        return self.data

    @property
    def bucket_id(self) -> str:
        return self.data["JobsDetail"]["BucketId"]

    @property
    def bucket_region(self) -> str:
        return self.data["JobsDetail"]["Region"]

    @property
    def creation_time(self) -> str:
        return self.data["JobsDetail"]["CreationTime"]


class TCIDocumentCallbackParser(TCICallbackParser):
    """
    document callback parser
    """

    @property
    def is_sensitive(self) -> bool:
        return self.data["JobsDetail"]["Suggestion"] != 0
