import boto3
from typing import Optional, List, BinaryIO
import logging

logger = logging.getLogger(__name__)

class S3Client:
	def __init__(self, bucket_name: str, region: str = 'us-east-1'):
		self.s3 = boto3.client('s3', region_name=region)
		self.bucket_name = bucket_name

	def upload_raw_object(self, key: str, stream: BinaryIO) -> bool:
		"""Upload a raw object to S3."""
		# TODO: implement
		raise NotImplementedError

	def list_objects(self, prefix: str) -> List[str]:
		"""List objects in a bucket with prefix."""
		# TODO: implement
		raise NotImplementedError

	def get_object_stream(self, key: str) -> Optional[BinaryIO]:
		"""Return a stream/bytes for an object."""
		# TODO: implement
		raise NotImplementedError

	def copy_object(self, src_key: str, dst_key: str, dst_bucket: Optional[str] = None) -> bool:
		"""Copy object within S3 or between buckets."""
		# TODO: implement
		raise NotImplementedError

	def delete_object(self, key: str) -> bool:
		"""Delete object from S3."""
		# TODO: implement
		raise NotImplementedError