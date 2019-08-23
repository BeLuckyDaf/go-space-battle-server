# Space Battle Server (Go)

This article describes the API for this awesome little server.

## Response
All responses follow this style: 
````
{
    status: true // or false and then error message in data
    data: { ... }
}
````

## API

### Connect
Connect to the server with the provided username.

#### Parameters
* **Username** : _String_

#### Returns
````
{
    status: true,
    data: Player // look up in player.go
}
````

### Move
Move the player to the specified location. The target point must be adjacent
to the current position of the player.

#### Parameters
* **Username** : _String_
* **Token** : _String_
* **Location** : _Integer_ (new location)

#### Returns
````
{
    status: true,
    data: Player // look up in player.go, the player will have new location
}
````

### Status
Get the status of the world.

#### Parameters

#### Returns
````
{
    status: true,
    data: Server // look up in server.go
}
````

### Players
Get the list of all players in the world.

#### Parameters

#### Returns
````
{
    status: true,
    data: [Player, ...] // look up in player.go
}
````