# Orum Coding Interview

Time: Please limit yourself to 2 hours

## Languages

You can choose between Go or TypeScript (Node.js) for your implementation. Either implementation will require knowledge of SQL as well.

Scaffolding for Go is located in the `go/` directory.

Scaffolding for TypeScript is located in the `typescript/` directory.

For whichever language you choose, you can use the directory for that language as the root of your project.

## Summary

We're evaluating your ability to take a set of requirements and build a solution that demonstrates craftsmanship, thoughtfulness, and good architectural design within the allotted time.

## Implement an API for querying transfers data

The goal of this exercise is to design and implement an API that makes use of a provided sqlite database containing transfer, person, and account data to fulfill the requirements listed below.  You are free to make assumptions as to what would be most useful for the developer integrating with your API. You should document any assumptions made and the reasoning behind them in a readme.md file.  If you think of something that would be really useful for the developer but do not have time to build the feature, document it in the roadmap.md file.
Your solution should have some way to run locally and test the results so we can fully analyze your efforts. We ❤️ tests at Orum so please include what you see fit. 

### Database

We have provided a sqlite database containing the data you will need to query: `transfers.db`.

The database has the following schema:

```sql
CREATE TABLE customers (
    id TEXT NOT NULL PRIMARY KEY,
    first_name TEXT,
    last_name TEXT,
    date_of_birth DATE,
    email TEXT
);

CREATE TABLE accounts (
    id TEXT NOT NULL PRIMARY KEY,
    customer_id TEXT NOT NULL,
    routing_number TEXT,
    account_number TEXT,

    FOREIGN KEY (customer_id) REFERENCES customers(id)
);

CREATE TABLE transfers (
    id TEXT NOT NULL PRIMARY KEY,
    timestamp DATETIME NOT NULL,
    amount INTEGER NOT NULL,
    status TEXT NOT NULL,
    source_account_id TEXT,
    dest_account_id TEXT,

    FOREIGN KEY (source_account_id) REFERENCES accounts(id),
    FOREIGN KEY (dest_account_id) REFERENCES accounts(id)
);
```

### Endpoints

#### Transfers

We would like 2 endpoints implemented for this feature:

GET /transfers (with pagination)

GET /transfers/{id}

The response should contain the following attributes:
```text
- id
- timestamp
- amount
- status
- source customer name
- source account id
- destination customer name
- destination account id
```

Bonus: Feel free to add anything to the API that would make it easier for a developer to consume.

#### Accounts

We would like 1 endpoint implement for this feature:

POST /accounts

The request should contain the following attributes:
```text
- customer id
- account number
- routing number
```

The request should only save to the database if both the routing number is valid and the customer already exists in the database, otherwise it should return an error to the user.

To determine whether a routing number is valid, it must conform to the following formula, where d<sub>n</sub> refers to the value of the digit in the n<sup>th</sup> place of the routing number.

(3(d<sub>1</sub> + d<sub>4</sub> + d<sub>7</sub>) + 7(d<sub>2</sub> + d<sub>5</sub> + d<sub>8</sub>) + (d<sub>3</sub> + d<sub>6</sub> + d<sub>9</sub>)) mod 10 = 0

For more information, please refer to the following link: [ABA Routing Transit Number](https://en.wikipedia.org/wiki/ABA_routing_transit_number#Check_digit)

The response should contain the following attributes:
```text
- id
- customer id
- account number
- routing number
```

## Results

Be sure to include the following in your submission
A README.md with the following information

1. How to test/demo/run (if applicable)
2. Any feedback/notes (i.e., if something was hard, confusing, frustrating, etc.)
3. Anything else you'd like us to know about your submission
4. Can you think of any ways to optimize the queries?
5. (BONUS) A ROADMAP.md with what you would add/change if you had more time. Dream big.

## FAQ

### Can I use external libraries?

Yes, please do. You can use any dependencies you like. Our only ask is that you explain your choices or are ready to do so in a followup interview.

### Can I use an ORM?

Sure can. We'll make sure to evaluate your SQL skills in a followup interview.

### Should I write tests?

Yes. We take pride in our tests at Orum so please include them.

## Submitting your code

When you're ready to submit your code, we ask that you create a **private** github repo and add the following users as collaborators.

[@jimstrate](https://github.com/jimstrate)
[@TheManSpeaker] (https://github.com/TheManSpeaker)


## Questions

Please email code-interview@orum.io if you have any questions regarding the assignment.
