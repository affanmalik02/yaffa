# YAFFA: Yet Another Fundamental Financial Analyzer

**YAFFA** is a distributed data platform designed to ingest, normalize, and analyze SEC financial filings for 3,000+ public companies. By leveraging the Medallion Architecture, it transforms raw XBRL JSON data into high-fidelity financial ratios used for outlier detection and valuation modeling.

## 🛠 Tech Stack
- **Orchestration:** Apache Airflow
- **Distributed Compute:** Apache Spark (PySpark)
- **Deep Learning:** PyTorch
- **Backend:** Golang (Gin Framework)
- **Storage:** AWS S3 (Data Lake) & MongoDB (Metadata)

## 🏗 System Architecture

### 1. Ingestion Layer (Bronze)
- **Airflow DAGs** trigger Python workers to crawl the SEC EDGAR `company-tickers` mapping.
- Downloads bulk `companyfacts.zip` or hits REST endpoints to store raw JSON filings into **S3 (Bronze Layer)**.
- Implements **Rate Limiting** (10 requests/sec) to comply with SEC fair-access policies.

### 2. Processing Layer (Silver)
- **Spark Jobs** flatten the deeply nested JSON facts.
- **Normalization:** Maps varied XBRL tags (e.g., `NetIncome` vs `NetLoss`) to a unified schema.
- **Deduplication:** Handles the "Duplicate Fact" problem in SEC data by filtering for the latest `accession_number` per reporting period.

### 3. Analytics Layer (Gold)
- Calculates **30+ Financial Ratios** in parallel:
  - Liquidity (Current Ratio, Quick Ratio)
  - Profitability (ROE, Net Margin, EBITDA Growth)
  - Valuation (P/E, Debt-to-Equity, FCF Yield)
- Stores aggregated "Gold" records in Parquet for high-performance ML training.

### 4. Machine Learning (Valuation Model)
- **Architecture:** A Residual Multi-Layer Perceptron (ResMLP) in **PyTorch**.
- **Task:** Regression to predict "Forward EPS" based on 5 years of historical fundamental ratios.
- **Inference:** Identifies companies where the current market price deviates significantly from the fundamental "Fair Value" prediction.

## 🚀 Key Achievements
- Processed **10+ years of historical data** for the entire S&P 500 in under 15 minutes using a Spark cluster.
- Built a **Go-based REST API** to serve "Top Undervalued" queries with sub-50ms latency using indexed lookups.