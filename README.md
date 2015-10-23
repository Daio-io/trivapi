# trivapi

[![Go ReportCard](http://goreportcard.com/badge/daio-io/trivapi)](http://goreportcard.com/report/daio-io/trivapi)

Pronounced "Triv-Appy", a trivia API. Query me for multiple choice trivia questions.

Work In Progress. Deployed to Heroku to test: (Questions results will be available soon)

## Request

    http://trivapi.herokuapp.com/randomise
    http://trivapi.herokuapp.com/category/science

## Categories

  science
  maths
  general

## Queries

Amount of questions:

    ?amount=10

Difficulty level:

    ?difficulty=2

#### Difficulty Levels

EASY = 1
MEDIUM = 2
HARD = 3
