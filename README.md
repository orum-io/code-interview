# Orum Coding Interview

Time: Please limit yourself to 2 hours
Languages: Go, SQL

## Summary

The purpose of this exercise is to design and implement the best possible solution to the exercise below in the time allotted. We're evaluating your ability to take a set of requirements and build a solution that demonstrates craftsmanship, thoughtfulness, and good architectural design.
Your solution should have some way to run locally and test the results so we can fully analyze your efforts.

## Implement an API for querying transfers data

The goal of this exercise is to design and implement a read-only API that returns one or more records from a static set of transfer data fulfilling the requirements listed below.  You are free to make assumptions as to what would be most useful for the developer integrating with your API. You should document any assumptions made and the reasoning behind them in a readme.md file.  If you think of something that would be really useful for the developer but do not have time to build the feature, document it in the roadmap.md file.

We would like 2 endpoints implemented for this feature

GET /transfers

GET /transfers/{id}

We would like the following fields returned for a transfer:
    - id
    - timestamp
    - amount
    - status
    - source person name
    - source account id
    - destination person name
    - destination account id


Be sure to include the following in your submission
A README.md with the following information
1. How to test/demo/run (if applicable)
2. Any feedback/notes (i.e., if something was hard, confusing, frustrating, etc.)
3. Anything else you'd like us to know about your submission
4. Can you think of any ways to optimize the queries?
5. (BONUS) A ROADMAP.md with what you would add/change if you had more time. Dream big.
6. (BONUS) A super-simple test suite, if applicable (even one test is a bonus)
