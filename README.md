# Cron Parser

This is a simple cron parser that allows you to interpret cron expressions.

## Getting Started

Follow the steps below to set up the cron parser:

### Step 0: Code Walkthrough

The codebase is rather small, with:

- [constants](constants/constants.go) : File to store magic numbers,
  etc.
- [parser](parser/parser.go) : function to encapsulate a
  CronExpression.
- [validate](parser/validate.go) : Util class 
  which is responsible
  for taking a raw expression and converting it into `CronExpression`.
- [Main](main.go) : Entry point.

### Step 1: Initialization

Clone the repository and unzip in your local machine
* Install golang: https://golang.org/doc/install
* Download and Install all the dependent packages
```bash
 go mod tidy
```
* Build the project
```bash
 go build
 ```

### Step 2: Run

Run the `go run main.go` script to along with the expression to see the results.

```bash
go run main.go "*/15 0 1,15 * 1-5 /usr/bin/find"
```

### Test

You can run the tests using

```bash
go test cron-parser/tests
```

or you can check out the test files @[Test Files](tests/validate_test.go)
