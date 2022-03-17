# Go SDK

A simple Go SDK for interacting with API V1.

## Installation

You can install the module via go modules:
```bash
$ go get -u github.com/miniyarov/go-sdk
```

## Usage

### Using the Merchant API client

You can use the `MerchantAPI` module for convenient access to API methods. Some are defined in the code:

```go
package main

import (
	"fmt"

	api "github.com/miniyarov/go-sdk"
)

func main() {
	client := api.NewDefaultClient()
	merchantAPI := api.NewMerchantAPI(client, "Merchant-Key", "Secret-Key")
	resp, _ := merchantAPI.InvoiceInfo("INVOICE_HASH") // errors ignored for simplicity
	fmt.Println(resp)
}
```

## License

The MIT License (MIT). Please see [License File](LICENSE.md) for more information.
