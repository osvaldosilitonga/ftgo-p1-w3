# Create Table

## Employee

```sql
CREATE TABLE employees (
	id INT PRIMARY KEY AUTO_INCREMENT,
	first_name VARCHAR(50) NOT NULL,
	last_name VARCHAR(50) NOT NULL,
	position VARCHAR(50)
	);
```

## Menu Items

```sql
CREATE TABLE menu_items (
	id INT PRIMARY KEY AUTO_INCREMENT,
	name VARCHAR(100) NOT NULL,
	description VARCHAR(255),
	price DOUBLE(5,2) NOT NULL,
	category VARCHAR(50)
	);
```

## Orders

```sql
CREATE TABLE orders (
	id INT PRIMARY KEY AUTO_INCREMENT,
	table_number INT,
	employee_id INT,
	order_date DATETIME,
	status VARCHAR(50),
  FOREIGN KEY (employee_id) REFERENCES employees(id)
  );
```

## Order Items

```sql
CREATE TABLE order_items (
	id INT PRIMARY KEY AUTO_INCREMENT,
	order_id INT,
	item_id INT,
	quantity INT,
	subtotal DOUBLE(5,2),
	FOREIGN KEY (order_id) REFERENCES orders(id),
	FOREIGN KEY (item_id) REFERENCES menu_items(id)
	);
```

## Payments

```sql
CREATE TABLE payments (
	id INT PRIMARY KEY AUTO_INCREMENT,
	order_id INT,
	payment_date DATETIME,
	method VARCHAR(20),
	total_amount DOUBLE(5,2),
	discount INT DEFAULT 0,
	FOREIGN KEY (order_id) REFERENCES orders(id)
	);
```

# Insert Sample Data

## Employees

```sql
INSERT INTO employees (first_name, last_name, position)
VALUES
	('John', 'Doe', 'Waiter'),
	('Jane', 'Smith', 'Cashier'),
	('Budi', 'Septiawan', 'Waiter'),
	('Umar', 'Bakri', 'Chef'),
	('Pevita', 'Pearce', 'Manager');
```

## Menu Items

```sql
INSERT INTO menu_items (name, description, price, category)
VALUES
	('Steak', 'Grilled sirloin steak', 25.99, 'Main Course'),
	('Potato Leek', 'Creamy soup with carrots, celery & leeks', 45.99, 'Main Course'),
	('Shrimp Cilanto Wrap', 'Shrimp, avocado, mixed greens, salsa, cilantro & may on a sundried tomato tortilla', 25.99, 'Main Course'),
	('Red Iceberg Salad', 'With sweet corn, blackberries, goat cheese & fresh basil', 35.99, 'Salads'),
	('Malaspina Oysters', "Fresh from Canada's Sunchine Coast. Served raw with cucumber-basil mignonette", 55.99, 'Appetizer');
```

## Orders

```sql
INSERT INTO orders (table_number, employee_id, order_date, status)
VALUES
	(2, 1, '2023-08-04', 'Pending'),
	(3, 1, '2023-07-04', 'Completed'),
	(5, 3, '2023-07-04', 'Completed'),
	(1, 1, '2023-08-04', 'Pending'),
	(4, 3, '2023-08-04', 'Cancel');
```

## Order Items

```sql
INSERT INTO order_items (order_id, item_id, quantity, subtotal)
VALUES
	(1, 1, 2, 51.98),
	(1, 2, 1, 45.99),
	(2, 3, 2, 77.97),
	(2, 1, 1, 25.99),
	(3, 4, 1, 35.99);
```

## Payments

```sql
INSERT INTO payments (order_id, payment_date, method, total_amount)
VALUES
	(1, '2023-08-04', 'Credit Card', 51.98),
	(2, '2023-07-04', 'E-Wallet', 41.4),
	(3, '2023-07-04', 'Credit Card', 51.98),
	(4, '2023-08-04', 'Cash', 32.4);
```

# Additional Instruction

## A. Retrieve all orders with their applied discounts

```sql
SELECT * FROM payments WHERE discount > 0;
```

## B. Calculate the total revenue (including discounts) for a specific day

```sql
SELECT SUM(total_amount) AS 'Total Revenue'
FROM payments
WHERE payment_date = '2023-07-04';
```
