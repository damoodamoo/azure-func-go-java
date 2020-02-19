# Azure Functions with Go
This repo contains some sample code experimenting with the new http worker coming in Azure Functions. More details and samples for this feature can be found at [Pragna Gopa's](https://github.com/pragnagopa/) repo [here](https://github.com/pragnagopa/functions-http-worker/).

## Structure
Essentially, this repo contains a simple, dirty, API written in Go. With the new http worker feature, we can spin up our own process along with the functions runtime, and functions will simply marshall requests from triggers + bindings to our own http endpoints. 

There is no 'function code' - just json triggers and bindings, which point at our own HTTP endpoints which are not aware of functions. 

# To Run
- Make sure you've got functions core tools installed and up to date -> [docs](https://docs.microsoft.com/en-us/azure/azure-functions/functions-run-local).
- Clone this repo
- Update the values in `local.settings.json` to point at your own storage accounts / cosmos etc as needed
- Build the Go API:
```
go build ./go/go-http-server
```
- Run the functions host:
```
func start
```
- Hit the endpoints in Postman / your api testing tool. For the `add` endpoint, use the following json schema:
```json
{
	"id": 1,
	"name": "bananas"
}
```

### Operations:
- `/api/add`: `POST` the above schema
- `/api/get?id=1`: `GET` an item
- `/api/list`: `GET` all items
- `/api/send-items`: `GET`. Send all items to a storage queue
- `process-items` (non-http): Trigger on queue and post to cosmos and secondary queue


## Disclaimer
All code is sample, ugly, and likely to break :)


