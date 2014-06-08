Config Service
---

Just a simple configuration service to help me learn golang.

Serves configuration information in the form of JSON and HTML. Both interfaces allow you to update the underlying data.

Data is stored to AWS S3 by default, but you can extend the data-store interface to provide other backings.

Use
---

```shell
go get github.com/jessejlt/configsrvs
./configsrvs -h
```

Status
---

- [x] Command line options
- [x] Environment variables for S3 credentials
- [x] Tests!
- S3
  - [x] S3 validation
  - [ ] S3 data retrieval
  - [ ] S3 data update
- JSON API
  - [x] GET
  - [ ] PUT
- Web interface
  - [x] Populate template
  - [ ] Add some style
  - [ ] Make data updatable
