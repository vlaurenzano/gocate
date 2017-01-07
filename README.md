[![Build Status](https://travis-ci.org/vlaurenzano/gocate.svg?branch=master)](https://travis-ci.org/vlaurenzano/gocate)
# gocate
gocate is an implentation of the locate command written in golang.

# Usage
Gocate searches for all occurences of a string in it's database. First use `gocate -u` to populate the database, then  use:
`gocate string` to find all occurences. Gocate will present you all matches as well as your top 5 matches (scored based on string distance to the end of path).

#Configuration
Gocate uses environment variables for configuration and falls back to default values if not provided.

- `GOCATE_DB_LOCATION` Where to store the gocate database. By default it is stored in /tmp
- `GOCATE_BUILD_INDEX_STRATEGY` Concurrent | Iterative. The default is concurrent which uses n goroutines to walk the file system. Iterative uses golang's filepath.Walk.
- `GOCATE_N_BUILD_JOBS` How many concurrent jobs to use when building the database. Defaults to 100. Can be tuned for performance. If set too high, will hit the os' open file limit.
