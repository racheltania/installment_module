# Installment Module

## How to Run

### Option 1: Run Directly from Installment.exe

You can run the project directly using the executable file `installment.exe`.

### Option 2: Run from Project Source

If you want to run it directly from this project, follow these steps:

1. Download the Go modules:
    ```bash
    go mod download
    ```

2. Run the project:
    ```bash
    go run *.go
    ```

## Testing with Postman

After the project is running, you can test it using Postman.

- **Method:** POST
- **URL:** http://localhost:8080/installment
- **Body:** Raw

  ```json
  {
      "plafond": 10000000,
      "loan_duration_in_month": 12,
      "annual_interest_rate": 5,
      "start_date": "2022-08-10"
  }
