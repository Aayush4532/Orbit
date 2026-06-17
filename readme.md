# Orbit

A high-concurrency event commerce platform built with Go, Redis, MongoDB, and Gin.

Orbit is designed to handle scenarios where a large number of users compete for limited inventory in real time. The system focuses on fairness, atomic allocation, low-latency booking, and scalable event management.

---

## Overview

Orbit allows sellers to create events and publish limited inventory products. Once an event goes live, inventory is loaded into Redis, enabling fast booking operations without placing load on MongoDB.

The booking engine uses Redis Lua Scripts to guarantee atomic inventory allocation under heavy concurrency.

Examples:

* Event ticket sales
* Limited seat course registrations
* Flash sales
* Limited inventory launches
* First-Come-First-Serve booking systems

---

## Architecture

MongoDB acts as the primary persistence layer.

Redis acts as the real-time booking engine.

```text
Seller
  ↓
MongoDB
  ↓
Event Live
  ↓
Redis Warmup

Buyer
  ↓
Gin API
  ↓
Redis Lua Script
  ↓
Atomic Inventory Allocation
  ↓
Success / Sold Out
```

---

## Core Features

### Event Management

* Create Events
* Update Events
* Publish Events
* Event Lifecycle Management

### Product Management

* Add Products
* Update Products
* Inventory Management
* Event Specific Products

### Booking Engine

* Redis Based Inventory
* Atomic Booking using Lua Scripts
* First-Come-First-Serve Allocation
* Idempotent Booking
* Concurrent Request Handling
* Sold Out Protection

### User System

* Registration
* Authentication
* JWT Authorization
* Role Based Access Control

---

## Tech Stack

### Backend

* Golang
* Gin Framework

### Database

* MongoDB

### Cache & Booking Engine

* Redis
* Redis Lua Scripts

### Authentication

* JWT
* HTTP Cookies

---

## Booking Flow

When an event becomes live:

1. Products are loaded from MongoDB.
2. Worker Pool pushes inventory into Redis.
3. Buyers interact only with Redis.
4. Atomic Lua Scripts allocate inventory.
5. Successful bookings are recorded.
6. Sold-out inventory is rejected instantly.

---

## Inventory Model

Inventory Stock:

```text
product:{productId}:{eventId}
```

Example:

```text
product:abc123:event789 = 50
```

Product Metadata:

```text
productmeta:{productId}:{eventId}
```

Stored Fields:

* Price
* Seller ID
* Product Title
* Currency

---

## Concurrency Strategy

Orbit is designed to survive high traffic scenarios.

Features:

* Redis Atomic Operations
* Lua Scripts
* Idempotency Keys
* Worker Pools
* Pipeline Writes
* Fast Failure for Sold Out Inventory

Goal:

```text
Inventory = 100

10000 Concurrent Requests

Exactly 100 Winners
Exactly 9900 Failures
Inventory Final = 0
```

---

## Future Improvements

* Payment Gateway Integration
* Reservation Window
* Redis Streams
* Booking Persistence Workers
* WebSocket Inventory Updates
* Distributed Rate Limiting
* Admission Control Layer
* Reverse Proxy Layer
* Multi Region Deployment

---

## Project Structure

```text
api/
cmd/
configs/

internal/
├── buyer/
├── seller/
├── repositories/
├── middleware/
├── models/
├── db/
├── worker/
├── utils/

pkg/
```

---

## Why This Project?

Most booking systems fail because inventory allocation is treated as a simple CRUD problem.

Orbit focuses on the difficult part:

* Fair allocation
* Concurrent bookings
* Low latency inventory access
* High traffic handling

The project is built to explore production-grade backend engineering concepts such as caching, distributed systems, atomic operations, concurrency control, and scalable booking architectures.
