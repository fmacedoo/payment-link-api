# Stripe Payment Link API

This project provides a small API for generating Stripe payment links and handling webhooks. It is built using Go and exposes the following endpoints:

- `POST /create-payment-link` - Generate a Stripe payment link.
- `POST /webhook` - Handle Stripe webhooks.

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) \(version 1.20 or newer recommended\)
- [Stripe API Key](https://stripe.com/docs/keys)

### Steps to Run

1. Clone this repository:

   ```bash
   git clone https://github.com/fmacedoo/stripe-paymentlink-api.git
   cd stripe-paymentlink-api
   ```

2. Export your Stripe API key as an environment variable:

   ```bash
   export STRIPE_API_KEY=your_stripe_api_key
   ```

3. Run the application:

   ```bash
   go run .
   ```

4. By default, the API will start at \`http://localhost:8080\`. You can use tools like [Postman](https://www.postman.com/) or \`curl\` to interact with the endpoints.

### Using Air for Live Reloading

[Air](https://github.com/cosmtrek/air) is a Go development tool that provides live reloading for your application during development.

#### Installing Air

1. Install Air using Go:

   ```bash
   go install github.com/air-verse/air@latest
   ```

2. Ensure that the Air binary is in your \`PATH\`. You can add the Go binary directory to your \`PATH\` if it\'s not already included:

   ```bash
   export PATH=\$PATH:\$(go env GOPATH)/bin
   ```

3. Verify the installation:
   ```bash
   air -v
   ```

#### Running the Application with Air

1. Create an \`air.toml\` configuration file in the project root:

   ```bash
   air init
   ```

2. Customize the configuration as needed. By default, Air watches for changes in \`.go\` files and restarts the application automatically.

3. Run the application with live reloading:
   ```bash
   air
   ```

Now, any changes you make to your Go files will trigger an automatic restart of the application.

## Endpoints

### \`/create-payment-link\`

- **Method**: \`POST\`
- **Description**: Generates a Stripe payment link.
- **Request Body**:
  ```json
  {
    "amount": 5000,
    "currency": "usd",
    "description": "Sample payment"
  }
  ```
- **Response**:
  ```json
  {
    "url": "https://payment.stripe.com/payment-link"
  }
  ```

### \`/webhook\`

- **Method**: \`POST\`
- **Description**: Handles Stripe webhooks for payment updates.
- **Request Body**: Raw Stripe webhook event.

#### Activate the Stripe webhook

- Download and install [Stripe CLI](https://docs.stripe.com/stripe-cli)

```
stripe listen --forward-to http://localhost:9800/webhook
stripe trigger payment_intent.succeeded
```

## License

This project is licensed under the [MIT License](LICENSE).