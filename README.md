# [API] Named
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/diegovictor/named-api?style=flat-square)
[![MIT License](https://img.shields.io/badge/license-MIT-green?style=flat-square)](https://raw.githubusercontent.com/DiegoVictor/named-api/main/LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)<br>
[![Run in Insomnia}](https://insomnia.rest/images/run.svg)](https://insomnia.rest/run/?label=Named&uri=https%3A%2F%2Fraw.githubusercontent.com%2FDiegoVictor%2Fnamed-api%2Fmain%2FInsomnia_2022-01-30.json)

Responsible for provide data to the [web](https://github.com/DiegoVictor/named-web) front-end. Generates names based on preseted datasets (list of example names), but you can upload a custom dataset (it must have at least 23 names). Was utilized the Markov Chain algorithm to generate the names, a custom script was created based on this article: [Generating Startup names with Markov Chains](https://towardsdatascience.com/generating-startup-names-with-markov-chains-2a33030a4ac0).

Live Demo: https://name-generator-api.herokuapp.com

## Table of Contents
* [Requirements](#requirements)
* [Usage](#usage)
  * [Routes](#routes)
    * [Requests](#requests)

# Requirements
* [Go](https://go.dev/)

# Usage
To start up the app run:
```
$ go run .
```
Or:
```
go run main.go
```

## Routes
|route|HTTP Method|params|description
|:---|:---:|:---:|:---:
|`/datasets`|GET| - |Retrieve available datasets.
|`/names`|GET|`dataset` query parameter.|Generate names for the dataset provided.
|`/feedbacks`|POST|Body with an array of feedbacks with `name` and `value`.|Save good feedbacks
|`/upload`|POST|Body with dataset [form data](https://developer.mozilla.org/docs/Web/API/FormData) (See insomnia file for good example).|Create a custom dataset.


### Requests
* `POST /feedbacks`

Request body:
```
[
  {
    "name": "Jordan",
    "value": 1
  }
]
```

* `POST /upload`

Request body:
```
"file"=<file>
```
