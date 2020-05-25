<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="">
    <img src="https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcRj2GH5OYJicARKz2cV2epalmOOjp93dztz0bGaIaHaWg0szgAb&usqp=CAU" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center"># upper-tweet</h3>

  <p align="center">
    An awesome README for upper-tweet
    <br />
    <strong>Explore the docs »</strong></a>
    <br />
    <a>View Demo</a>
    ·
    <a>Report Bug</a>
    ·
    <a>Request Feature</a>
  </p>
</p>

⚠️🚧⚠️

<!-- TABLE OF CONTENTS -->
## Table of Contents

* [About the Project](#about-the-project)
  * [Built With](#built-with)
* [Getting Started](#getting-started)
  * [Prerequisites](#prerequisites)
  * [Installation](#installation)
* [Usage](#usage)


<!-- ABOUT THE PROJECT -->
## About The Project

### Built With
* [Gin](https://github.com/gin-gonic/gin)
* [upper](https://github.com/upper/db)


<!-- GETTING STARTED -->
## Getting Started

### Prerequisites

This is an example of how to list things you need to use the software and how to install them.
* docker
[cmmon.. ;d](https://docs.docker.com/get-docker/)

<!-- USAGE EXAMPLES -->
## Usage

Build the app: 
```shell script
make build
```

Run migrations: 
```shell script
make migrate-db-up
```

Drop migrations (1): 
```shell script
make migrate-db-down
```

Show logs: 
```shell script
make logs
```

Jump into psql: 
```shell script
make psql
```