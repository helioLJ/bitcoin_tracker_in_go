## **Project Title**: Bitcoin Price Tracker

### **Objective**
Develop a Go-based command-line application (CLI) that allows users to:
1. Fetch and display real-time Bitcoin prices in multiple currencies.
2. Periodically update prices and notify users of significant changes.
3. Log historical prices to a local file for review.

---

### **Features**
1. **Fetch Current Prices**:
   - Retrieve real-time Bitcoin prices in user-specified currencies using a public API.
   
2. **Periodic Updates**:
   - Continuously fetch and update Bitcoin prices at configurable intervals using goroutines.

3. **Price Alerts**:
   - Notify the user when Bitcoin prices cross specified thresholds.

4. **Historical Logging**:
   - Log price data to a local file with timestamps for future analysis.

---

### **Go Concepts Incorporated**
1. **Imports/Packages**:
   - Utilize standard libraries (`net/http`, `encoding/json`, `os`, `time`) and third-party packages as needed (e.g., `github.com/joho/godotenv` for managing API keys).

2. **Structs/Methods**:
   - Create structs for price data and configuration settings.
   - Define methods to fetch prices, format output, and handle updates.

3. **Error Handling**:
   - Gracefully manage HTTP request errors, JSON parsing issues, and file I/O problems.
   - Display meaningful messages to users for troubleshooting.

4. **Go Routines**:
   - Fetch Bitcoin prices concurrently without blocking user inputs.
   - Allow for background tasks like logging and threshold monitoring.

5. **Channels**:
   - Use channels for inter-goroutine communication, enabling real-time price updates and alerts.

---

### **Advanced Go Concepts (Optional)**
1. **cgo**: 
   - Add optional high-performance computation for price trend analysis.
2. **Generics**: 
   - Implement generic methods for handling API responses or data transformations.
3. **Iterators**:
   - Use iterators for managing and displaying logged price history.


---

### **Implementation Plan**

#### **1. Core Features Implementation**
1. Fetch real-time Bitcoin prices using a public API (e.g., CoinGecko or CoinMarketCap).
2. Log historical data in a human-readable format.

#### **2. Application Flow**
1. **User Input**:
   - Select currency for price tracking.
   - Set optional price thresholds for alerts.
2. **Concurrent Operations**:
   - Run real-time price updates in a background goroutine.
3. **Display Results**:
   - Continuously update the terminal with the latest price and any alerts.

#### **3. Incorporate Go Concepts**
1. **Imports/Packages**:
   - Use Go modules for managing dependencies and clean code organization.
2. **Structs/Methods**:
   - Define reusable structs for price data and configuration.
3. **Error Handling**:
   - Implement proper error handling for all operations.
4. **Go Routines**:
   - Fetch price data and monitor thresholds concurrently.
5. **Channels**:
   - Use channels to coordinate data fetching and alert notifications.

---

### **Stretch Goals**
- Add support for multiple cryptocurrencies (e.g., Ethereum, Litecoin).
- Implement a simple web server to display real-time prices using `net/http`.
- Use a visualization library to generate charts of historical price trends.

---

### **Deliverables**
1. A functional CLI application for Bitcoin price tracking.
2. Code that demonstrates mastery of Go basics:
   - Imports/Packages
   - Structs/Methods
   - Error Handling
   - Goroutines and Channels
3. Documentation for setup, usage, and an explanation of the Go concepts used.
