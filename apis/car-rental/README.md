# Guidelines

---

## Testing

There are bdd tests at [route](./route/) folder and databases tests at [repository](./internal/repository/) folder.  
To run all tests execute at project's root folder the command below:

```shell
go test -v ./...
```

## Implementations

- Automated tests:
  - Usage of mocks for repository tests;
  - Repository tests;
  - Endpoint tests;
    - including controller tests;
- API Rest implementation with Gin Gonic;
- Logging with zap (although not used that much in this project);
- MVC architecture;

### Possible next steps if there was enough time

> This list follows a personal priority order.

- Add Makefile shortcuts;
- Improve types and interfaces nesting;
- Dockerize database and API service;
- Implement authentication;
- Implement scheduling;
- Implement notification;
- Implement caching;
- Implement pubsub;
- Improve logs - more needed and used if previous topics get implemented;
- Add more comments to methods;

---

## Overview

- This project is a mock of a real-life problem. We intentionally made this project too much work to complete in the given time, and we do not expect you to work on this in one sitting.
- The tasks are in no particular order. Your work will be assessed not only on how much you complete but also **how** you complete the tasks you decide to build. Work will be graded on, but not limited to, code design, reusability, scalability, coding style, and writing maintainable code. Please focus more on the business logic.
- Feel free to use any standard go libraries (<https://golang.org/pkg/#stdlib>). Please email us if you need another dependency before adding a non-standard library package.
- Please email us with as many questions as you want. Questions sent during non-working hours may have a slower response time.

## Setup

  1. Install mysql locally and get it running:

  ```bash
  brew install mysql
 brew services start mysql
  ```

  2. Create local user and grant privileges:

  ```bash
  mysql -u root

  CREATE USER 'car_rental'@'localhost' IDENTIFIED WITH mysql_native_password BY 'car_rental';

  GRANT ALL PRIVILEGES ON *.* TO 'car_rental'@'localhost'; exit;
  ```

  3. Create local challenge DB:

  ```bash
  mysql -u car_rental -p
  > car_rental

  create database car-rental; exit;
  ```

  4. Fetch `gorm`, `mysql`, and `uuid` packages:

  ```bash
 go get -u github.com/gofrs/uuid
 go get -u gorm.io/driver/mysql
 go get -u gorm.io/gorm
  ```

## Run

- In the root directory, run `go run .`

## Notes

- No code is required in the `gateway` directory. That package is set up with mock functions that just print to the console. You can assume they behave as in the comments.
- To test your endpoints, call the function in `main.go`, and run `go run .`

## Submission

- Zip the entire project folder with your completed work. Run the following command in the parent directory of this project:

  ```bash
  zip -r coding-challenge-go-<your_first_name>.zip coding-challenge-go-<your_first_name>
  ```

- Email us the generated zip file.

---

## Car Rental Service

You will begin offering vehicle rentals for partners. Create a Car Rental Service for a Mobile Frontend and an Internal Manager Dashboard.

### Project Requirements

**This challenge aims to implement endpoints to our service handler that follow the specification below.**

### Booking Frontend (External Users) will

- Allow users to rent a car for X amount of hours. Users should receive a confirmation text message (use Texter Gateway inside `gateway/texter.go`).

- Users should receive a text message 24 hours before the scheduled pickup time (Use Scheduler `gateway/schedule.go`).

- Implement an endpoint that allows the user to check in and confirm Vehicle Pickup.

- Users will be charged at the time of pickup and should receive a text message stating that their vehicle has been picked up, along with the amount charged for the entire rental time.

- Users should receive a text message 1 hour before the scheduled rental return time.

- Implement an endpoint with logic to allow the user to confirm Vehicle Drop-off.

- Users should receive a text message stating that their vehicle has been dropped off, along with the final charged price (there should be an overtime fee of 1.5x rate).

- Users should receive a text message if the vehicle has yet to be returned past the scheduled due time.

- Users should be able to cancel a reservation with the following cancellation policy:
  - Canceling within 24 hours is non-refundable
  - Canceling 48 hours in advance will result in a 25% non-refundable deposit (i.e., 75% returned to the user)
  - Users should receive a text message stating that the reservation has been canceled, along with the final price they were charged

- Users should be able to update their reservation (pickup and drop-off times).

### Manager Dashboard (Internal Users)

Create endpoints that will be used to display information about our rental service to our internal Ops team. This new dashboard should be able to show:

- Real-time available vehicle inventory for a specific lot.
- Currently unavailable inventory and when each vehicle is expected to be returned.
- Overdue inventory with contact information for the renter.
- Schedule of availability for a specific vehicle.

Keep in mind that this dashboard will be used globally, so we only want to query the database once a user navigates to the dashboard.

### Bonus if you have time

- Add logging for events that should be captured.

- Add tests to your handler functions.
