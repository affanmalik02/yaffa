# yaffa: Yet Another Fundamental Financial Analyzer

YAFFA is an end-to-end financial intelligence platform that merges professional-grade fundamental analysis with personal portfolio tracking. By combining distributed processing of SEC filings, real-time market data via yfinance, and secure bank linking via Plaid, YAFFA provides users with a "ground-truth" view of their investments versus intrinsic business value.

## Technical Stack

- **Orchestration:** Apache Airflow
- **Distributed Compute:** Apache Spark (PySpark)
- **Deep Learning:** PyTorch (Valuation Modeling)
- **Backend:** Golang (Gin Framework)
- **Account Integration:** Plaid API (Holdings & Investments)
- **Frontend:** Next.js 15, TypeScript, Tailwind CSS
- **AI/LLM:** Qwen2 7B / Baichuan2 7B (Open-source, locally-hosted via Ollama or vLLM)
- **Storage:** AWS S3 (Data Lake), MongoDB (User Metadata & Tokens), Parquet (Analytical Store)

## System Architecture

### 1. Multi-Source Ingestion Layer

The system synchronizes data from three distinct pipelines to provide a holistic market view:

**SEC EDGAR:** Automated extraction of historical 10-K/Q filings for "Ground Truth" fundamental data.

**yfinance API:** Integration for real-time market pricing, historical volatility, adjusted closes, and company metadata.

**Plaid API:** Secure OAuth-based linking of user bank and brokerage accounts to ingest real-time portfolio holdings and cost-basis data.

### 2. Standardized Processing Layer (Spark)

**Normalization:** Spark jobs reconcile inconsistent XBRL tags from SEC filings into a unified schema.

**Cross-Source Mapping:** A mapping service links SEC CIK identifiers with yfinance tickers and Plaid security IDs to ensure data consistency across the platform.

**Deduplication:** Implements logic to resolve conflicting data points between official filings and third-party market aggregators.

### 3. Portfolio & Fundamental Analytics

**Feature Engineering:** Distributed computation of 30+ financial ratios (ROE, P/E, FCF Yield) alongside portfolio-specific metrics like Diversification Score and Sector Exposure.

**Valuation Engine:** A PyTorch-based ResMLP model trains on historical fundamentals to predict a stock's "Fair Value," which is then compared against the user's actual brokerage cost-basis.

### 4. Analytical Feature Layer

**Parallel Computation:** Calculates 30+ core financial metrics across the entire dataset using distributed Spark executors:

- **Liquidity:** Current Ratio, Quick Ratio, Operating Cash Flow
- **Profitability:** Return on Equity (ROE), Net Margin, EBITDA Growth
- **Valuation:** P/E Ratio, Debt-to-Equity, Free Cash Flow (FCF) Yield
- **Portfolio:** Diversification Score, Sector Exposure, Cost-Basis Analysis

**Storage:** Persists aggregated features in partitioned Parquet files for optimized Machine Learning training and low-latency retrieval.

### 5. Machine Learning: Valuation Modeling

**Model Architecture:** Implements a Residual Multi-Layer Perceptron (ResMLP) using PyTorch.

**Objective:** Performs high-dimensional regression to predict Forward EPS and Fair Value based on a five-year rolling window of historical fundamental ratios.

**Anomaly Detection:** Identifies "valuation gaps" where market pricing significantly deviates from the model-calculated intrinsic value.

## User Interface & Features

### Analysis Dashboard

A high-density dashboard built with Next.js and Recharts that allows users to:

- **Search & Analyze:** Deep-dive into any ticker to view fundamental health scores and AI-generated summaries of SEC risk factors
- **Portfolio Oversight:** View linked brokerage holdings synced via Plaid, overlaying YAFFA's "Fair Value" predictions on top of current market prices
- **Decision Support:** Visualize the delta between a company's fundamental performance and its current market sentiment

### Qualitative AI Insights

**RAG Pipeline:** Utilizes open-source Qwen2 7B or Baichuan2 7B models, deployed locally via Ollama or vLLM, to parse the "Management Discussion & Analysis" (MD&A) sections of filings. This approach eliminates API costs and latency while maintaining strong multilingual financial reasoning.

**Synthesized Summaries:** Generates executive summaries that highlight management guidance and operational risks, providing context that raw numbers often miss.

**Cost Efficiency:** By hosting the LLM locally on modest hardware (8GB+ GPU or CPU inference), YAFFA avoids per-token charges from commercial APIs, making the system economically viable for production-scale filing analysis.

## Backend Infrastructure (Go)

**Unified API:** A versioned REST API built in Gin that aggregates data from S3 (fundamentals), MongoDB (user portfolio), and yfinance (price).

**Plaid Integration:** Manages the exchange of public tokens for access tokens, stored securely with encryption at rest.

**LLM Service Proxy:** Go service that routes filing text to a local Ollama/vLLM endpoint for inference and caches results in MongoDB.

**Security & Auth:** Implements API Key validation and rate limiting to protect internal data services.

**Performance:** Utilizes Go's concurrency model (Goroutines) to fetch market data and user holdings in parallel, ensuring sub-100ms dashboard refreshes.

## Implementation Roadmap

### File Structure

The project is organized as follows:

```
yaffa/
├── dags/                          # Airflow orchestration
│   └── ingest_dag.py
├── storage/                       # S3 and data storage
│   └── s3_client.py
├── jobs/                          # Spark processing jobs
│   ├── flatten_xbrl.py
│   └── compute_features.py
├── backend/                       # Golang API server
│   ├── cmd/api/
│   │   └── main.go
│   └── internal/
│       ├── middleware/
│       │   ├── apikey.go
│       │   └── ratelimit.go
│       ├── handlers/
│       │   └── fundamentals.go
│       ├── models/
│       │   └── types.go
│       └── db/
│           └── mongo.go
├── ml/                            # Machine Learning (PyTorch)
│   ├── models/
│   │   └── resmlp.py
│   └── train/
│       └── train.py
├── rag/                           # RAG & LLM utilities
│   ├── retriever.py
│   └── generator.py
├── frontend/                      # Next.js frontend
│   └── src/
│       ├── pages/api/
│       │   └── fundamentals/
│       │       └── [ticker].ts
│       └── lib/
│           └── api.ts
├── tools/                         # Utilities & DevOps
│   └── openapi/
│       └── generate.go
└── README.md
```

### Implementation Modules

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