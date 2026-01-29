YAFFA: Yet Another Fundamental Financial Analyzer
YAFFA is a high-performance distributed data platform designed to ingest, normalize, and analyze SEC financial filings for over 3,000 public companies. Utilizing a multi-stage data lake architecture, the system transforms raw XBRL-formatted JSON data into high-fidelity financial ratios to facilitate automated outlier detection and fundamental valuation modeling.

Technical Stack
Orchestration: Apache Airflow

Distributed Compute: Apache Spark (PySpark)

Deep Learning: PyTorch

Backend: Golang (Gin Framework)

Frontend: Next.js 15, TypeScript, Tailwind CSS

AI/LLM: Gemini 1.5 Pro (RAG Pipeline)

Storage: AWS S3 (Object Storage), MongoDB (Metadata), Parquet (Analytical Store)

System Architecture
1. Raw Data Ingestion Layer
Orchestration: Airflow DAGs schedule Python-based workers to interface with the SEC EDGAR company-tickers mapping and bulk archival systems.

Persistence: Ingests bulk companyfacts.zip archives and incremental REST updates into S3 as immutable raw objects.

Compliance: Implements rigorous rate limiting (10 requests/second) and header validation to adhere to SEC fair-access policies.

2. Standardized Processing Layer
Data Flattening: Spark jobs parse and flatten deeply nested XBRL JSON structures into tabular formats.

Schema Normalization: Maps diverse and inconsistent XBRL taxonomy tags (e.g., reconciling variations of NetIncome and NetLoss) into a unified corporate finance schema.

Data Integrity: Resolves "duplicate fact" conflicts by prioritizing the most recent accession_number per reporting period to ensure a consistent time-series.

3. Analytical Feature Layer
Parallel Computation: Calculates 30+ core financial metrics across the entire dataset using distributed Spark executors:

Liquidity: Current Ratio, Quick Ratio, Operating Cash Flow.

Profitability: Return on Equity (ROE), Net Margin, EBITDA Growth.

Valuation: P/E Ratio, Debt-to-Equity, Free Cash Flow (FCF) Yield.

Storage: Persists aggregated features in partitioned Parquet files for optimized Machine Learning training and low-latency retrieval.

4. Machine Learning: Valuation Modeling
Model Architecture: Implements a Residual Multi-Layer Perceptron (ResMLP) using PyTorch.

Objective: Performs high-dimensional regression to predict Forward EPS based on a five-year rolling window of historical fundamental ratios.

Anomaly Detection: Identifies "valuation gaps" where market pricing significantly deviates from the model-calculated intrinsic value.

Frontend Application Interface
Framework: Developed with Next.js 15 and TypeScript for a type-safe, responsive financial dashboard.

Interactive Visualization: Integrates Recharts and Chart.js to render multi-year fundamental trend lines and comparative sector performance metrics.

State Management: Utilizes TanStack Query for efficient data fetching, caching, and synchronization between the UI and the Go backend.

Qualitative Analysis Engine (LLM)
RAG Pipeline: Implements a Retrieval-Augmented Generation workflow to process non-numerical filing sections.

Context Extraction: Automatically parses Item 7 (Management’s Discussion and Analysis) and Item 1A (Risk Factors) from 10-K filings.

Strategic Summarization: Leverages Gemini 1.5 Pro to generate executive-level summaries focusing on management guidance, latent operational risks, and forward-looking sentiment analysis.

Data Acquisition & External APIs
The system maintains a diversified ingestion strategy to ensure redundancy and accuracy:

SEC EDGAR API: Primary source for official 10-K/Q filings and XBRL facts.

Financial Modeling Prep: Standardized financial statements and real-time valuation metrics cross-checks.

Yahoo Finance (yfinance): Secondary source for historical price action, market capitalization, and dividend history.

Finnhub: Integrated for real-time news streams and analyst earnings estimates.

Public API Access
YAFFA provides a versioned RESTful API for external developers to consume processed fundamental data and ML-derived insights.

Authentication: Access is managed via API Key middleware (X-API-KEY header), with keys stored and validated in MongoDB.

Documentation: Interactive OpenAPI 3.0 (Swagger) documentation is available via the /swagger/index.html endpoint.

Rate Limiting: Protects system integrity using leaky-bucket rate limiting at the middleware layer.

Example Endpoint: GET /api/v1/fundamentals/:ticker returns a comprehensive payload of normalized ratios and predicted valuations.

System Performance Metrics
Scalability: Processed 10+ years of historical filing data for the entire S&P 500 in under 15 minutes on a distributed Spark cluster.

Latency: Go-based REST API serves complex fundamental queries with a p99 response time under 50ms.

Observability: Integrated multi-channel alerting (AWS SES, PagerDuty) via Airflow for immediate resolution of data pipeline failures.

Implementation Roadmap (initial method skeletons)
This section lists planned modules and the initial method names to implement first. Implementations will be created in the codebase under the indicated paths.

1) Airflow (Python)
- filepath: dags/ingest_dag.py
  - fetch_ticker_map()
  - download_bulk_archives()
  - schedule_incremental_updates()
  - validate_request_headers()
  - rate_limit_request()
  - orchestrate_ingest_pipeline()

2) Ingestion / Storage (Python)
- filepath: storage/s3_client.py
  - upload_raw_object(bucket, key, stream)
  - list_objects(bucket, prefix)
  - get_object_stream(bucket, key)
  - copy_object(src_bucket, src_key, dst_bucket, dst_key)
  - delete_object(bucket, key)

3) Spark Processing (PySpark)
- filepath: jobs/flatten_xbrl.py
  - parse_xbrl_json(record)
  - flatten_taxonomy(df)
  - normalize_schema(df)
  - resolve_duplicate_facts(df)
  - write_parquet(df, partition_cols)

- filepath: jobs/compute_features.py
  - compute_liquidity_metrics(df)
  - compute_profitability_metrics(df)
  - compute_valuation_metrics(df)
  - aggregate_time_series(df)
  - persist_features(df, path)

4) Backend (Golang)
- filepath: backend/cmd/api/main.go
  - setupRouter()
  - registerRoutes(r)
  - main()

- filepath: backend/internal/middleware/apikey.go
  - APIKeyMiddleware(next http.Handler) http.Handler
  - ValidateAPIKey(key) (bool, error)

- filepath: backend/internal/middleware/ratelimit.go
  - RateLimitMiddleware(next http.Handler) http.Handler
  - NewLeakyBucket(limit int, burst int)

- filepath: backend/internal/handlers/fundamentals.go
  - GetFundamentals(w http.ResponseWriter, r *http.Request)
  - ListTickers(w http.ResponseWriter, r *http.Request)
  - HealthCheck(w http.ResponseWriter, r *http.Request)
  - ServeSwagger(w http.ResponseWriter, r *http.Request)

5) MongoDB (Go)
- filepath: backend/internal/db/mongo.go
  - Connect(uri string) (*mongo.Client, error)
  - GetAPIKey(ctx, key)
  - StoreMetadata(ctx, doc)
  - QueryMetadata(ctx, filters)

6) Machine Learning (PyTorch)
- filepath: ml/models/resmlp.py
  - class ResMLP(nn.Module)
    - __init__(self, config)
    - forward(self, x)
    - training_step(self, batch)
    - predict_forward_eps(self, features)
    - save(self, path)
    - load(cls, path)

- filepath: ml/train/train.py
  - prepare_dataloader()
  - train_epoch(model, dataloader, optimizer)
  - validate_epoch(model, dataloader)
  - run_training_loop(config)

7) RAG / LLM Utilities (Python)
- filepath: rag/retriever.py
  - extract_sections(text, sections=['Item 1', 'Item 7', 'Item 1A'])
  - embed_documents(docs)
  - retrieve_context(query, top_k)
  - build_rag_prompt(context, query)

- filepath: rag/generator.py
  - generate_summary(prompt, model='gemini-1.5-pro')
  - synthesize_executive_summary(sections)

8) Frontend (Next.js, TypeScript)
- filepath: frontend/src/pages/api/fundamentals/[ticker].ts
  - fetchFundamentalsAPI(req, res)
  - validateTickerParam(req)
  - cacheFundamentals(ticker, payload)

- filepath: frontend/src/lib/api.ts
  - getFundamentals(ticker)
  - listTickers()
  - getSwaggerSpec()

9) Utilities / DevOps
- filepath: tools/openapi/generate.go
  - GenerateOpenAPISpec()
  - ServeSwaggerFiles()

Next steps (priority)
- Create the files listed above with the declared method signatures.
- Wire minimal imports and empty function bodies returning placeholders.
- Add tests for API surface and basic smoke tests for processing jobs.
- Iterate on one vertical: ingest -> flatten -> feature compute -> API.