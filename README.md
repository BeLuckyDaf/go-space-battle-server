# Space Battle Server (Go)

[![Build Status](https://travis-ci.org/BeLuckyDaf/go-space-battle-server.svg?branch=master)](https://travis-ci.org/BeLuckyDaf/go-space-battle-server)

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
Connect to the server with the provide username.

#### Parameters
* **Username** : _String_

#### Returns
````
{
    status: true,
    data: Player // look up in player.go
}
````
#### Example
````
/connect?username=USERNAME
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
#### Example
````
/move?username=USERNAME&token=TOKEN&target=ID
````

### World
Get the status of the world.

#### Parameters

#### Returns
NOTE: This one increases drastically with the number of points in the world. Possibly there is a better representation of the world. It was made this way with simplicity in mind, assuming that `/world` should not be called too often.

**Location types:**
* Planet = 0
* Asteroid = 1
* Station = 2

````
{
    status: true,
    data: World // look up in world.go
}
````
#### Example
````
/world
````

### Owned
Get the points IDs and the usernames of their owners,

#### Parameters

#### Returns
This returns a map, where the key is the ID (int) of the point
and the value is the username of the person, who owns it. If a point is not 
owned, it is not shown.

````
{
    status: true,
    data: map[int]string // { id: "username", ... }
}
````
#### Example
````
/owned
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
#### Example
````
/players
````

### Buy
Buy the location you are currently standing on.

#### Parameters
* **Username** : _String_
* **Token** : _String_

#### Returns
````
{
    status: true,
    data: Location // look up in location.go
}
````
#### Example
````
/buy?username=USERNAME&token=TOKEN&target=ID
````

### Destroy
Destroy the location you are currently standing on.

#### Parameters
* **Username** : _String_
* **Token** : _String_

#### Returns
````
{
    status: true,
    data: Location // look up in location.go
}
````
#### Example
````
/destroy?username=USERNAME&token=TOKEN&target=ID
````

### Attack
Attack a player in your position.

#### Parameters
* **Username** : _String_
* **Token** : _String_
* **Target** : _String_

Target player username is required.

#### Returns
````
{
    status: true,
    data: Location // look up in location.go
}
````
#### Example
````
/attack?username=MY_USERNAME&token=TOKEN&target=TARGET_USERNAME
````

### Trade
Give some of your power to another player

#### Parameters
* **Username** : _String_
* **Token** : _String_
* **Recipient** : _String_
* **Amount** : _Int_

Target player username is required.

#### Returns
````
{
    status: true,
    data: string // Non informative text message
}
````
#### Example
````
/trade?username=MY_USERNAME&token=TOKEN&recipient=TARGET_USERNAME&amount=NUMBER
````