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
