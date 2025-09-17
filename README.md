````markdown
# Lombard Checker

A simple Go-based tool to fetch wallet allocations from [Lombard Finance](https://lombard.finance) claim API.  
Supports residential proxies and parallelized requests for multiple wallets.  

---

## 🚀 Getting Started

### 1. Clone the repository
```bash
git clone https://github.com/socryptos/lombard_checker.git
````

### 2. Configure `config.go`

Navigate to `app/internal/config/config.go` and update settings:

* Enable/disable proxy via the `UseProxy` flag.
* Replace `ProxyAddress` with your **residential proxy**.

### 3. Add wallet addresses

Paste your wallet addresses (one per line) into:

```
app/data/wallets.txt
```

### 4. Run the app

Move into the `app` directory and start the program:

```bash
cd app
go run cmd/main.go
```

---

## ⚙️ Features

* Fetches claimable allocations for given wallet addresses.
* Supports proxy configuration (residential only).
* Handles multiple wallets concurrently with a worker pool.
* Provides total allocation summary.

---

## 📌 Notes

* Proxies must be provided in the format:

  ```
  http://username:password@host:port
  ```
* Recommended to use **residential proxies** only.

---
