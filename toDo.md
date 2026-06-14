POST /api/seller/events/create # to create a single event.
GET  /api/seller/events/get     # to get all events of that particular seller.
GET  /api/seller/events/get/:id 
PUT  /api/seller/events/update/:id  
DELETE /api/seller/events/delete/:id 

POST /api/seller/events/:eventId/registerProducts # to register products in the event.
GET  /api/seller/events/:eventId/getProducts           <----# to get all products of the event.  for all buyers as well as sellers.--->
GET  /api/seller/events/:eventId/getProduct/:id      <---------# to get particular product. for all buyers as well as sellers. ----->
PUT  /api/seller/events/:eventId/updateProduct/:id # to update particualr product
DELETE /api/seller/events/:eventId/deleteProduct/:id # to delete particular product

POST /api/seller/events/:eventId/live/:id # to start live the particular event 
POST /api/seller/events/:eventId/pause/:id # to pause the particular event.
POST /api/seller/events/:eventId/stop/:id  # to completely end  the particular event.

# these are after event routes.
GET  /api/seller/events/booking/fetch/:eventId  # to get all the booking of details of the event.
GET  /api/seller/events/analytics/:id # analystics of the event