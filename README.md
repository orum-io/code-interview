# Orum Coding Interview

Time: Please limit yourself to 2 hours

## Languages

You can choose between Go or TypeScript (Node.js) for your implementation. Either implementation will require knowledge of SQL as well.

Scaffolding for Go is located in [this branch](https://github.com/orum-io/code-interview/tree/go).

Scaffolding for TypeScript is located in [this branch](https://github.com/orum-io/code-interview/tree/typescript).

## Summary

We're evaluating your ability to take a set of requirements and build a solution that demonstrates craftsmanship, thoughtfulness, and good architectural design within the allotted time.

## Implement an API for querying transfers data

The goal of this exercise is to design and implement a read-only API that returns one or more records from a static set of transfer data fulfilling the requirements listed below.  You are free to make assumptions as to what would be most useful for the developer integrating with your API. You should document any assumptions made and the reasoning behind them in a readme.md file.  If you think of something that would be really useful for the developer but do not have time to build the feature, document it in the roadmap.md file.
Your solution should have some way to run locally and test the results so we can fully analyze your efforts.

We have provided a sqlite database containing the data you will need to query: `transfers.db`.

### Endpoints

We would like 2 endpoints implemented for this feature:

GET /transfers (with pagination)

GET /transfers/{id}

### Fields

We would like the following fields returned for a transfer:

    - id
    - timestamp
    - amount
    - status
    - source person name
    - source account id
    - destination person name
    - destination account id

Bonus: Feel free to add anything to the API that would make it easier for a developer to consume.

## Results

Be sure to include the following in your submission
A README.md with the following information

1. How to test/demo/run (if applicable)
2. Any feedback/notes (i.e., if something was hard, confusing, frustrating, etc.)
3. Anything else you'd like us to know about your submission
4. Can you think of any ways to optimize the queries?
5. (BONUS) A ROADMAP.md with what you would add/change if you had more time. Dream big.

## Submitting your code

When you're ready to submit your code, we ask that you create a **private** github repo and add the following users as collaborators.

[@jimstrate](https://github.com/jimstrate)

[@TheManSpeaker](https://github.com/TheManSpeaker)

[@bengentree](https://github.com/bengentree)
