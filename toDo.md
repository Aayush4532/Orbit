POST /api/seller/events/:eventId/pause/:id # to pause the particular event.
POST /api/seller/events/:eventId/stop/:id  # to completely end  the particular event.

# these are after event routes.
GET  /api/seller/events/booking/fetch/:eventId  # to get all the booking of details of the event.
GET  /api/seller/events/analytics/:id # analystics of the event


//----------------------- Abstracted ToDo-------------------------//
3. Main Booking Engine. 
4. Reverse Proxy Server.
5. go worker to push data from redis to the database.
6. buyer routes.
7. Admin Routes.
8. Client Side interaction Website for this. 

---------------//----------------------------MVP END--------------------------------//----------