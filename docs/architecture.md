# Money Transfer System - High-Level Architecture

## 1. Overview

The Money Transfer System is a fully fledged API based system for transferring money between customer accounts.
System enables creating and modifying customers and accounts for those customers.

## 2. Main Parts

### 2.1 Authentication Service

- It ensures that only authorized users can access the system and perform actions for authenticated customer.

### 2.2 Customer Service

- Create, update, retrieve customer information.

### 2.3 Account Service

- Create, update, retrieve account information.

### 2.4 Transaction Service

- Manage money transfers between accounts.

### 2.5 Database

- Store data for customers, accounts, and transactions.

## 3. Used Technologies

- **GO, Echo Framework**
- **JWT**
- **SQLite**
