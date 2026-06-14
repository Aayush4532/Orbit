POST /api/seller/events/:eventId/live/:id # to start live the particular event 
POST /api/seller/events/:eventId/pause/:id # to pause the particular event.
POST /api/seller/events/:eventId/stop/:id  # to completely end  the particular event.

# these are after event routes.
GET  /api/seller/events/booking/fetch/:eventId  # to get all the booking of details of the event.
GET  /api/seller/events/analytics/:id # analystics of the event