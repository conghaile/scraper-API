# Documentation

This API was created to allow HTML scrapers to write data to a Postgres database remotely. In the future, the repository will include more endpoints for more data sources, and a Docker image for deployment on cloud services such as AWS.

Currently, the API has a single endpoint for posts scraped from [the Warosu /biz/ archive](https://warosu.org/biz). It supports POST and GET requests; the latter requires the request to be fomatted as such:

```
{"number": <int>,
"text": [
    <string>,
    <string>,
    <string>
    ],
"time": <int>}
```
The "text" array should have one element per sentence in a scraped post, and the array may be arbitrarily long.
