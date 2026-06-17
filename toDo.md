POST /api/seller/events/:eventId/pause/:id # to pause the particular event.
POST /api/seller/events/:eventId/stop/:id  # to completely end  the particular event.

# these are after event routes.
GET  /api/seller/events/booking/fetch/:eventId  # to get all the booking of details of the event.
GET  /api/seller/events/analytics/:id # analystics of the event


//----------------------- Abstracted ToDo-------------------------//
1. Main Booking Engine. 
2. go worker to push data from redis to the database. 
4. Admin Routes. 
5. Client Side interaction Website for this. 
6. Reverse Proxy Server. 

---------------//----------------------------MVP END--------------------------------//----------