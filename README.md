# Transaction Service

## Overview
This is a simple transaction service with three endpoints to create accounts, retrieve account information, and create transactions.

## Endpoints
### Create an Account
**POST /accounts**

**Request Body:**
```json
{
  "document_number": "12345678900"
}
{
  "status": "success",
  "message": "Account created successfully",
  "data": {
    "account_id": 1,
    "document_number": "12345678900"
  },
  "code": 0
}

Retrieve Account Information
GET /accounts/:accountId  Response: