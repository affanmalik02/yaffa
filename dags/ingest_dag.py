from datetime import datetime, timedelta

from airflow import DAG
from airflow.operators.python import PythonOperator

import logging

logger = logging.getLogger(__name__)

default_args = {
    'owner': 'yaffa',
    'retries': 2,
    'retry_delay': timedelta(minutes=5),
}

"""
Airflow DAG skeleton for YAFFA ingestion pipeline.

TODO:
- Implement the individual task functions.
- Wire up the DAG with the task dependencies.
"""

def fetch_ticker_map():
	"""Fetch mapping of tickers from EDGAR or local cache."""
	# TODO: implement
	raise NotImplementedError

def download_bulk_archives():
	"""Download companyfacts bulk archives."""
	# TODO: implement
	raise NotImplementedError

def schedule_incremental_updates():
	"""Schedule incremental SEC REST updates."""
	# TODO: implement
	raise NotImplementedError

def validate_request_headers(headers=None):
	"""Validate headers before calling SEC endpoints."""
	# TODO: implement
	raise NotImplementedError

def rate_limit_request():
	"""Rate-limit wrapper (10 req/s default)."""
	# TODO: implement
	raise NotImplementedError

def orchestrate_ingest_pipeline():
	"""Top-level orchestration entrypoint for DAG tasks."""
	# TODO: implement
	raise NotImplementedError

with DAG(
	dag_id='yaffa_ingest_pipeline',
	default_args=default_args,
	description='YAFFA SEC EDGAR ingestion pipeline',
	schedule_interval='@daily',
	start_date=datetime(2024, 1, 1),
	catchup=False,
) as dag:
	ingest_task = PythonOperator(
		task_id='orchestrate_ingest',
		python_callable=orchestrate_ingest_pipeline,
	)

	ingest_task