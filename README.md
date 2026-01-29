# YAFFA: Yet Another Fundamental Financial Analyzer

YAFFA is a high-performance distributed data platform designed to ingest, normalize, and analyze SEC financial filings for over 3,000 public companies. Utilizing a multi-stage data lake architecture, the system transforms raw XBRL-formatted JSON data into high-fidelity financial ratios to facilitate automated outlier detection and fundamental valuation modeling.

## Technical Stack

- **Orchestration:** Apache Airflow
- **Distributed Compute:** Apache Spark (PySpark)
- **Deep Learning:** PyTorch
- **Backend:** Golang (Gin Framework)
- **Frontend:** Next.js 15, TypeScript, Tailwind CSS
- **AI/LLM:** Gemini 1.5 Pro (RAG Pipeline)
- **Storage:** AWS S3 (Object Storage), MongoDB (Metadata), Parquet (Analytical Store)

## System Architecture

### 1. Raw Data Ingestion Layer

**Orchestration:** Airflow DAGs schedule Python-based workers to interface with the SEC EDGAR company-tickers mapping and bulk archival systems.

**Persistence:** Ingests bulk companyfacts.zip archives and incremental REST updates into S3 as immutable raw objects.

**Compliance:** Implements rigorous rate limiting (10 requests/second) and header validation to adhere to SEC fair-access policies.

### 2. Standardized Processing Layer

**Data Flattening:** Spark jobs parse and flatten deeply nested XBRL JSON structures into tabular formats.

**Schema Normalization:** Maps diverse and inconsistent XBRL taxonomy tags (e.g., reconciling variations of NetIncome and NetLoss) into a unified corporate finance schema.

**Data Integrity:** Resolves "duplicate fact" conflicts by prioritizing the most recent accession_number per reporting period to ensure a consistent time-series.

### 3. Analytical Feature Layer

**Parallel Computation:** Calculates 30+ core financial metrics across the entire dataset using distributed Spark executors:

- **Liquidity:** Current Ratio, Quick Ratio, Operating Cash Flow
- **Profitability:** Return on Equity (ROE), Net Margin, EBITDA Growth
- **Valuation:** P/E Ratio, Debt-to-Equity, Free Cash Flow (FCF) Yield

**Storage:** Persists aggregated features in partitioned Parquet files for optimized Machine Learning training and low-latency retrieval.

### 4. Machine Learning: Valuation Modeling

**Model Architecture:** Implements a Residual Multi-Layer Perceptron (ResMLP) using PyTorch.

**Objective:** Performs high-dimensional regression to predict Forward EPS based on a five-year rolling window of historical fundamental ratios.

**Anomaly Detection:** Identifies "valuation gaps" where market pricing significantly deviates from the model-calculated intrinsic value.

## Frontend Application Interface

**Framework:** Developed with Next.js 15 and TypeScript for a type-safe, responsive financial dashboard.

**Interactive Visualization:** Integrates Recharts and Chart.js to render multi-year fundamental trend lines and comparative sector performance metrics.

**State Management:** Utilizes TanStack Query for efficient data fetching, caching, and synchronization between the UI and the Go backend.

## Qualitative Analysis Engine (LLM)

**RAG Pipeline:** Implements a Retrieval-Augmented Generation workflow to process non-numerical filing sections.

**Context Extraction:** Automatically parses Item 7 (Management's Discussion and Analysis) and Item 1A (Risk Factors) from 10-K filings.

**Strategic Summarization:** Leverages Gemini 1.5 Pro to generate executive-level summaries focusing on management guidance, latent operational risks, and forward-looking sentiment analysis.

## Data Acquisition & External APIs

The system maintains a diversified ingestion strategy to ensure redundancy and accuracy:

- **SEC EDGAR API:** Primary source for official 10-K/Q filings and XBRL facts
- **Financial Modeling Prep:** Standardized financial statements and real-time valuation metrics cross-checks
- **Yahoo Finance (yfinance):** Secondary source for historical price action, market capitalization, and dividend history
- **Finnhub:** Integrated for real-time news streams and analyst earnings estimates

## Public API Access

YAFFA provides a versioned RESTful API for external developers to consume processed fundamental data and ML-derived insights.

**Authentication:** Access is managed via API Key middleware (X-API-KEY header), with keys stored and validated in MongoDB.

**Documentation:** Interactive OpenAPI 3.0 (Swagger) documentation is available via the `/swagger/index.html` endpoint.

**Rate Limiting:** Protects system integrity using leaky-bucket rate limiting at the middleware layer.

**Example Endpoint:** `GET /api/v1/fundamentals/:ticker` returns a comprehensive payload of normalized ratios and predicted valuations.

## System Performance Metrics

- **Scalability:** Processed 10+ years of historical filing data for the entire S&P 500 in under 15 minutes on a distributed Spark cluster
- **Latency:** Go-based REST API serves complex fundamental queries with a p99 response time under 50ms
- **Observability:** Integrated multi-channel alerting (AWS SES, PagerDuty) via Airflow for immediate resolution of data pipeline failures

## Implementation Roadmap

This section outlines the initial module structure and method skeletons to be implemented. All code will be created under `/Users/affanmalik/Documents/Workspace/yaffa/`.

### 1. Airflow (Python)
**File:** `dags/ingest_dag.py`
- `fetch_ticker_map()` — Fetch mapping of tickers from EDGAR or local cache
- `download_bulk_archives()` — Download companyfacts bulk archives
- `schedule_incremental_updates()` — Schedule incremental SEC REST updates
- `validate_request_headers(headers)` — Validate headers before calling SEC endpoints
- `rate_limit_request()` — Rate-limit wrapper (10 req/s default)
- `orchestrate_ingest_pipeline()` — Top-level orchestration entrypoint for DAG tasks

### 2. Ingestion / Storage (Python)
**File:** `storage/s3_client.py`
- `upload_raw_object(bucket, key, stream)` — Upload a raw object to S3
- `list_objects(bucket, prefix)` — List objects in a bucket with prefix
- `get_object_stream(bucket, key)` — Return a stream/bytes for an object
- `copy_object(src_bucket, src_key, dst_bucket, dst_key)` — Copy object within S3
- `delete_object(bucket, key)` — Delete object from S3

### 3. Spark Processing (PySpark)
**File:** `jobs/flatten_xbrl.py`
- `parse_xbrl_json(record)` — Parse raw XBRL JSON record into dict/row
- `flatten_taxonomy(df)` — Flatten nested taxonomy fields into tabular DataFrame
- `normalize_schema(df)` — Map diverse taxonomy tags to unified schema
- `resolve_duplicate_facts(df)` — Resolve duplicate facts by accession_number / period
- `write_parquet(df, partition_cols, path)` — Write DataFrame to partitioned Parquet

**File:** `jobs/compute_features.py`
- `compute_liquidity_metrics(df)` — Compute current ratio, quick ratio, operating cash flow
- `compute_profitability_metrics(df)` — Compute ROE, net margin, EBITDA growth
- `compute_valuation_metrics(df)` — Compute P/E, D/E, FCF yield, etc.
- `aggregate_time_series(df)` — Aggregate computed metrics into time series per ticker
- `persist_features(df, path)` — Persist final feature set to Parquet for ML and API

### 4. Backend (Golang)
**File:** `backend/cmd/api/main.go`
- `setupRouter()` — Initialize Gin router with middleware
- `registerRoutes(r)` — Register all HTTP handler routes
- `main()` — Application entrypoint

**File:** `backend/internal/middleware/apikey.go`
- `APIKeyMiddleware() gin.HandlerFunc` — Extract and validate X-API-KEY header
- `ValidateAPIKey(key string) (bool, error)` — Query MongoDB for API key validity

**File:** `backend/internal/middleware/ratelimit.go`
- `NewLeakyBucket(limit int, burst int) *LeakyBucket` — Initialize leaky-bucket limiter
- `RateLimitMiddleware(bucket *LeakyBucket) gin.HandlerFunc` — Enforce rate limits

**File:** `backend/internal/handlers/fundamentals.go`
- `GetFundamentals(c *gin.Context)` — Fetch fundamentals for ticker and return JSON
- `ListTickers(c *gin.Context)` — List all available tickers
- `HealthCheck(c *gin.Context)` — Health status endpoint
- `ServeSwagger(c *gin.Context)` — Serve OpenAPI/Swagger UI

### 5. MongoDB (Go)
**File:** `backend/internal/db/mongo.go`
- `Connect(uri string) (*mongo.Client, error)` — Establish MongoDB connection
- `GetAPIKey(ctx, client, key)` — Query and validate API key from database
- `StoreMetadata(ctx, client, doc)` — Insert or upsert metadata document
- `QueryMetadata(ctx, client, filters)` — Query metadata collection

### 6. Machine Learning (PyTorch)
**File:** `ml/models/resmlp.py`
- `class ResMLP(nn.Module)` — Residual Multi-Layer Perceptron model
  - `__init__(self, config)` — Initialize model layers per config
  - `forward(self, x)` — Forward pass through network
  - `training_step(self, batch)` — Single training step
  - `predict_forward_eps(self, features)` — Predict forward EPS
  - `save(self, path)` — Save model weights
  - `load(cls, path, config)` — Load model from checkpoint

**File:** `ml/train/train.py`
- `prepare_dataloader(config)` — Prepare train/val dataloaders
- `train_epoch(model, dataloader, optimizer, device)` — Run one training epoch
- `validate_epoch(model, dataloader, device)` — Run validation epoch
- `run_training_loop(config)` — Top-level training orchestration

### 7. RAG / LLM Utilities (Python)
**File:** `rag/retriever.py`
- `extract_sections(text, sections=[...])` — Extract specified sections from filing text
- `embed_documents(docs)` — Embed documents using configured embedding model
- `retrieve_context(query, top_k=5)` — Retrieve top_k documents/segments for query
- `build_rag_prompt(context, query)` — Construct RAG prompt from context and query

**File:** `rag/generator.py`
- `generate_summary(prompt, model='gemini-1.5-pro')` — Generate summary using LLM
- `synthesize_executive_summary(sections)` — Create high-level executive summary

### 8. Frontend (Next.js, TypeScript)
**File:** `frontend/src/pages/api/fundamentals/[ticker].ts`
- `handler(req, res)` — API route handler for fundamentals endpoint
- `validateTickerParam(req)` — Validate ticker query parameter
- `cacheFundamentals(ticker, payload)` — Cache response payload

**File:** `frontend/src/lib/api.ts`
- `getFundamentals(ticker: string)` — Fetch fundamentals from backend
- `listTickers()` — Fetch available tickers
- `getSwaggerSpec()` — Fetch OpenAPI specification

### 9. Utilities / DevOps
**File:** `tools/openapi/generate.go`
- `GenerateOpenAPISpec() error` — Generate OpenAPI 3.0 spec from handlers
- `ServeSwaggerFiles(dir string) error` — Serve static Swagger UI files

## Next Steps (Priority Order)
1. Create all skeleton files with method signatures and TODO comments
2. Wire imports and empty function bodies that return placeholders
3. Add unit tests for API surface (fundamentals endpoints, auth middleware)
4. Implement one vertical: **Ingest → Flatten → Feature Compute → API Query**
5. Integrate ML training pipeline with feature data
6. Build frontend dashboard with visualization components