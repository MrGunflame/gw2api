# GW2API

A [Guild Wars 2 API](https://wiki.guildwars2.com/wiki/API:Main) Client in Go.

**This Module is still Work In Progress and not stable yet.**

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
    worlds, err := api.GetWorlds()
    if err != nil {
        panic(err)
    }

    fmt.Println(worlds)

    // Some endpoints will require an apikey
    // Set an apikey
    api = api.WithAcessToken("<YOUR APIKEY>")

    // Again, make a request
    account, err := gw2api.GetAccount()
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
- [x] /v2/account/home/cats
- [x] /v2/account/home/nodes
- [x] /v2/account/inventory
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
- [ ] /v2/commerce/transactions
- [ ] /v2/pvp/stats
- [ ] /v2/pvp/games
- [ ] /v2/pvp/standings
- [x] /v2/tokeninfo
- [ ] /v2/dailycrafting
- [ ] /v2/mapchests
- [ ] /v2/worldbosses
- [ ] /v2/masteries
- [ ] /v2/mounts
- [ ] /v2/mounts/skins
- [ ] /v2/mounts/types
- [ ] /v2/outfits
- [ ] /v2/pets
- [ ] /v2/professions
- [ ] /v2/races
- [ ] /v2/specializations
- [ ] /v2/skills
- [ ] /v2/traits
- [ ] /v2/legends
- [x] /v2/guild/:id
- [ ] /v2/emblem
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
- [ ] /v2/home/cats
- [ ] /v2/home/nodes
- [ ] /v2/finishers
- [ ] /v2/items
- [ ] /v2/itemstats
- [ ] /v2/materials
- [ ] /v2/pvp/amulets
- [ ] /v2/recipes
- [ ] /v2/recipes/search
- [ ] /v2/skins
- [ ] /v2/continents
- [ ] /v2/maps
- [ ] /v2/build
- [ ] /v2/colors
- [ ] /v2/currencies
- [ ] /v2/dungeons
- [ ] /v2/files
- [ ] /v2/quaggans
- [ ] /v2/minis
- [ ] /v2/novelties
- [ ] /v2/raids
- [ ] /v2/titles
- [x] /v2/worlds
- [ ] /v2/backstory/answers
- [ ] /v2/backstory/questions
- [ ] /v2/stories
- [ ] /v2/stories/seasons
- [ ] /v2/quests
- [ ] /v2/pvp
- [ ] /v2/pvp/ranks
- [ ] /v2/pvp/seasons
- [ ] /v2/pvp/seasons/:id/leaderboards
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
