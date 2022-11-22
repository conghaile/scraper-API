# Documentation

This API was created to allow HTML scrapers to write data to a Postgres database remotely. In the future, the repository will include a Docker image for deployment on cloud services such as AWS.

Currently, the API has a single endpoint for posts scraped from [the Warosu /biz/ archive](https://warosu.org/biz). It supports POST requests and requires the body to be fomatted as such:

```
{"number": <int>,
"text": [
    <string>,
    <string>,
    <string>
    ],
"time": <int>}
```
The array in the "text" field may be arbitrarily long. If you intend on doing NLP on post text in the future, it's best that each element in the "text" array contains a single sentence from the post for ease of parsing later.
