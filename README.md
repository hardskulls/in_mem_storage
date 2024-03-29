# Project
This is an example of an in-memory database.

## Design
The project is designed following the principles of `clean architecture` and `DDD`.

## Link: [deployed instance](https://in-mem-db-go.onrender.com)

## Usage
### General
Send your command to one of the available `API` entries using `&`-separated `key=value` 
pairs encoded in the body of request.

#### Rate Limiting
The rate limit for an `IP` address can be set by sending `for` and `limit` keys, 
where `for` is a valid `IP` address, and `limit` is a limit until the next 
request in milliseconds.
- Path `api/v1/rate_limiter`

#### Time to live
Can be set using the `Set` command.

#### Commands
There are 4 commands available in total:
- Set (requires `key`, `value` and `expires_after`)
- Get (requires `key`)
- Delete (requires `key`)
- Update (requires `key`, `value`)

`key` and `value` can be anything, `expires_after` 
must be the record's `time to live` in seconds.
- Path `api/v1/db_cmd`
