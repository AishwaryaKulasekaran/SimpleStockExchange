# SimpleStockExchange

A simple trading app with design for buy and sell orders.

Gin is used as the http web framework

Gin handles the individual routes and the router groups.

Once the user gives the order request, InitialProcess method is called.

# InitialProcess

Once the order is obtained from the user initial check of Order queue is done and based on the queue length order is queued or sent for further processing.

# ExecuteOrder

In the next step of processing, order information is retrieved and based on the type of the order Buy/Sell is performed.

Three types checks are performed before every Buy/Sell function

  •	Price check
	
  •	Quantity check
	
  •	Type check
	
Based on the flag from each check, buying/selling/queuing the order is done.

Queuing of orders is done up to the array size. 
