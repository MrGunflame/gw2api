# GW2API

[![Go Reference](https://pkg.go.dev/badge/gitlab.com/MrGunflame/gw2api.svg)](https://pkg.go.dev/gitlab.com/MrGunflame/gw2api)
[![Go Report Card](https://goreportcard.com/badge/gitlab.com/MrGunflame/gw2api)](https://goreportcard.com/report/gitlab.com/MrGunflame/gw2api)

A [Guild Wars 2 API](https://wiki.guildwars2.com/wiki/API:Main) Wrapper Client in Go.

## Installation

`go get -u gitlab.com/MrGunflame/gw2api`

## Examples

Create a new api session and use it to fetch endpoints.
```
package main

import (
    "fmt"

    "gitlab.com/MrGunflame/gw2api"
)

func main() {
    // Create a new session
    api := gw2api.New()

    // Make a request
    worlds, err := api.Worlds()
    if err != nil {
        panic(err)
    }

    fmt.Println(worlds)

    // Some endpoints will require an apikey
    // Set an api token
    api = api.WithAcessToken("<YOUR APIKEY>")

    // Make a request
    account, err := gw2api.Account()
    if err != nil {
        panic(err)
    }

    fmt.Println(account)
}
```

## Supported Endpoints

- [x] /v2/achievements
- [x] /v2/achievements/daily
- [x] /v2/achievements/daily/tomorrow
- [x] /v2/achievements/groups
- [x] /v2/achievements/categories

- [x] /v2/account
- [x] /v2/account/achievements
- [x] /v2/account/bank
- [x] /v2/account/dailycrafting
- [x] /v2/account/dungeons
- [x] /v2/account/dyes
- [x] /v2/account/finishers
- [x] /v2/account/gliders
- [x] /v2/account/home
- [x] /v2/account/home/cats
- [x] /v2/account/home/nodes
- [x] /v2/account/inventory
- [x] /v2/account/legendaryarmory
- [x] /v2/account/luck
- [x] /v2/account/mailcarriers
- [x] /v2/account/mapchests
- [x] /v2/account/masteries
- [x] /v2/account/mastery/points
- [x] /v2/account/materials
- [x] /v2/account/minis
- [x] /v2/account/mounts/skins
- [x] /v2/account/mounts/types
- [x] /v2/account/novelties
- [x] /v2/account/outfits
- [x] /v2/account/pvp/heroes
- [x] /v2/account/raids
- [x] /v2/account/recipes
- [x] /v2/account/skins
- [x] /v2/account/titles
- [x] /v2/account/wallet
- [x] /v2/account/worldbosses
- [x] /v2/characters
- [x] /v2/pvp/stats
- [x] /v2/pvp/games
- [x] /v2/pvp/standings
- [x] /v2/tokeninfo
- [x] /v2/dailycrafting
- [x] /v2/mapchests
- [x] /v2/worldbosses
- [x] /v2/masteries
- [x] /v2/mounts
- [x] /v2/mounts/skins
- [x] /v2/mounts/types
- [x] /v2/outfits
- [x] /v2/pets
- [x] /v2/professions
- [x] /v2/races
- [x] /v2/specializations
- [x] /v2/skills
- [x] /v2/traits
- [x] /v2/legendaryarmory
- [x] /v2/legends
- [x] /v2/guild/:id
- [x] /v2/emblem
- [x] /v2/guild/permissions
- [x] /v2/guild/search
- [x] /v2/guild/upgrades
- [x] /v2/guild/:id/log
- [x] /v2/guild/:id/members
- [x] /v2/guild/:id/ranks
- [x] /v2/guild/:id/stash
- [x] /v2/guild/:id/treasury
- [x] /v2/guild/:id/teams
- [x] /v2/guild/:id/upgrades
- [x] /v2/home/cats
- [x] /v2/home/nodes
- [x] /v2/finishers
- [x] /v2/items
- [x] /v2/itemstats
- [x] /v2/materials
- [x] /v2/pvp/amulets
- [x] /v2/recipes
- [x] /v2/recipes/search
- [x] /v2/skins
- [x] /v2/continents
- [x] /v2/maps
- [x] /v2/build
- [x] /v2/colors
- [x] /v2/currencies
- [x] /v2/dungeons
- [x] /v2/files
- [x] /v2/quaggans
- [x] /v2/minis
- [x] /v2/novelties
- [x] /v2/raids
- [x] /v2/titles
- [x] /v2/worlds
- [x] /v2/backstory/answers
- [x] /v2/backstory/questions
- [x] /v2/stories
- [x] /v2/stories/seasons
- [x] /v2/quests
- [x] /v2/pvp
- [x] /v2/pvp/ranks
- [x] /v2/pvp/seasons
- [x] /v2/pvp/seasons/:id/leaderboards
- [x] /v2/commerce/delivery
- [x] /v2/commerce/exchange
- [x] /v2/commerce/exchange/coins
- [x] /v2/commerce/exchange/gems
- [x] /v2/commerce/listings
- [x] /v2/commerce/prices
- [x] /v2/commerce/transactions
- [x] /v2/wvw/abilities
- [x] /v2/wvw/matches
- [x] /v2/wvw/objectives
- [x] /v2/wvw/ranks
- [x] /v2/wvw/upgrades
