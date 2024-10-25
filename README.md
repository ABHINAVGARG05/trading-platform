# _TRADING PLATFORM_
This is the backend of a Trading Platform in which user can buy/sell stocks and manage their portfolio along with real time stock price updation. The users also can view the details of individual stocks before making a buy. Also all transactions are stored in database.

# Project Structure
------

- `config`: Contains DB configuration.

- `controllers` This contains all the Controllers

     - `orderController.go`:This controls the buy and sell mechanism 
     - `priceController.go`: This controls the background price updation mechanism
     - `stockController.go`: This manages the mechanism for get all stocks and get stocks by their id.
    -  `userController.go`: Responsible for registering user and getting user information.

- `database`
   - `db.go`: This contains DB connection and seeding of stocks in DB.
- `models`
  - `user.go`: Defines schema for user
  - `stock.go`: Defines schema for stocks
  - `portfolio.go`: Defines schema for portfolio
  - `order.go`: Defines schema for order made

- `routes`: Contains all the routes 

## Installation and Setup
1. **Clone the Repository**
  ```bash
  git clone https://github.com/ABHINAVGARG05/trading-platform
  ```
2. **Install Dependencies**
  ```bash
  go mod tidy
  ```
3. **Install Air**
 ```bash
 go install github.com/air-verse/air@latest
 air init
 ```
4. **Run Dockerized**
  ```bash
  docker-compose up --build
  ```

## Deployment Details
### Deployment Example on a Server
1) **Install Docker and Docker Compose**:Ensure Docker and Docker Compose are installed on your server.
```bash
sudo apt update
sudo apt install docker.io docker-compose -y
```
2) **Clone Repository on Server**:
```bash
git clone https://github.com/ABHINAVGARG05/trading-platform
```
3)**Set Up Environment Variables**:
Create a ```.env``` and ```docker-compose.yml``` files with environment variables for database configuration.

4)**Build and start the application**
```bash 
docker-compose up --build
```

## SCALING
1)**Horizontal Scaling**
using ```Docker-compose.yml``` file scaling can be done by running multiple instances

2)**Database Scaling**
To support a large number of users and transactions, consider using managed database services like Amazon RDS, which provide high availability and automated backups.

3)**Cache Layer**
Use Redis or Memcached to cache frequently requested data (e.g., stock prices) and reduce database load.

## API DOCUMENTATION
### User Endpoints
### Register User
- **Endpoint**: `POST /api/v1/user/register`
- **Request Body**: 
 ```json
 {
    "ID": "userid",
    "Username": "username",
    "Balance": "Balanace"
}
```

### Get All Stocks
- **Endpoint**: `POST /api/v1/stocks`
- **ResponseBody**:
 ```json
 [
    {
        "ID": "number",
        "Symbol": "string",
        "Price": "number",
        "Volume": "number",
        "Supply": "number",
        "Demand": "number",
        "CreatedAt": "string",
        "UpdatedAt": "string"
    }
]
 ```
 ### Get Stock By Id
- **EndPoint**: `POST /api/v1/stock/id`
- **Response Body**:
 ```json
 {
    "type": "object",
    "properties": {
        "ID": {
            "type": "integer"
        },
        "Symbol": {
            "type": "string"
        },
        "Price": {
            "type": "number"
        },
        "Volume": {
            "type": "integer"
        },
        "Supply": {
            "type": "integer"
        },
        "Demand": {
            "type": "integer"
        },
        "CreatedAt": {
            "type": "string",
            "format": "date-time"
        },
        "UpdatedAt": {
            "type": "string",
            "format": "date-time"
        }
    }
}
 ```
### Buy Stock
- **Endpoint**: `api/v1/stock/buy`
- **Request Body**:
```json
{
    "user_id": user-id ,
    "stock_id": stock-id,
    "quantity":quantity
}
```
- **Response Body**:
```json
{
    "type": "object",
    "properties": {
        "ID": {
            "type": "number"
        },
        "UserID": {
            "type": "number"
        },
        "StockID": {
            "type": "number"
        },
        "Type": {
            "type": "string"
        },
        "Quantity": {
            "type": "number"
        },
        "Price": {
            "type": "number"
        },
        "CreatedAt": {
            "type": "string"
        },
        "UpdatedAt": {
            "type": "string"
        }
    }
}
```
### Sell Stock
- **EndPoint**: `api/v1/stock/sell`
- **Request Body**: 
```json
{
    "user_id": user-id ,
    "stock_id": stock-id,
    "quantity":quantity
}
```
- **Response Body**:
```json
{
    "type": "object",
    "properties": {
        "ID": {
            "type": "number"
        },
        "UserID": {
            "type": "number"
        },
        "StockID": {
            "type": "number"
        },
        "Type": {
            "type": "string"
        },
        "Quantity": {
            "type": "number"
        },
        "Price": {
            "type": "number"
        },
        "CreatedAt": {
            "type": "string"
        },
        "UpdatedAt": {
            "type": "string"
        }
    }
}
```
### Get User Information
- **Endpoint**: `api/v1/user/id`
- **Response Body**:
```json
{
    "type": "object",
    "properties": {
        "user": {
            "type": "object",
            "properties": {
                "ID": {
                    "type": "number"
                },
                "Username": {
                    "type": "string"
                },
                "Balance": {
                    "type": "number"
                },
                "CreatedAt": {
                    "type": "string"
                },
                "UpdatedAt": {
                    "type": "string"
                }
            }
        },
        "portfolio": {
            "type": "array",
            "items": {
                "type": "object",
                "properties": {
                    "ID": {
                        "type": "number"
                    },
                    "UserID": {
                        "type": "number"
                    },
                    "StockID": {
                        "type": "number"
                    },
                    "Quantity": {
                        "type": "number"
                    },
                    "AvgPrice": {
                        "type": "number"
                    },
                    "CreatedAt": {
                        "type": "string"
                    },
                    "UpdatedAt": {
                        "type": "string"
                    }
                }
            }
        }
    }
}
```
### Testing Server 

1. Use Postman
- Import [Postman Collection](https://elements.getpostman.com/redirect?entityId=36199199-9e5d57d4-cd93-426b-a673-e96e5e3741bb&entityType=collection) to test the API endpoints.
   - Set the environment variables if necessary.

