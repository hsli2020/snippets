package main

// https://tutorialedge.net/golang/consuming-restful-api-with-go/

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
)

// A Response struct to map the Entire Response
type Response struct {
    Name    string    `json:"name"`
    Pokemon []Pokemon `json:"pokemon_entries"`
}

// A Pokemon Struct to map every pokemon to.
type Pokemon struct {
    EntryNo int            `json:"entry_number"`
    Species PokemonSpecies `json:"pokemon_species"`
}

// A struct to map our Pokemon's Species which includes it's name
type PokemonSpecies struct {
    Name string `json:"name"`
}

func main() {
    response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    var responseObject Response
    json.Unmarshal(responseData, &responseObject)

    fmt.Println(responseObject.Name)
    fmt.Println(len(responseObject.Pokemon))

    for i := 0; i < len(responseObject.Pokemon); i++ {
        fmt.Println(responseObject.Pokemon[i].Species.Name)
    }
}
/*
{
  "descriptions": [
    {
      "description": "Red/Blue/Yellow Kanto dex",
      "language": {
        "name": "en",
        "url": "https://pokeapi.co/api/v2/language/9/"
      }
    },
    {
      "description": "Rot/Blau/Gelb Kanto Dex",
      "language": {
        "name": "de",
        "url": "https://pokeapi.co/api/v2/language/6/"
      }
    },
    {
      "description": "Pokédex régional de Kanto dans Rouge/Bleu/Jaune",
      "language": {
        "name": "fr",
        "url": "https://pokeapi.co/api/v2/language/5/"
      }
    }
  ],
  "id": 2,
  "is_main_series": true,
  "name": "kanto",
  "names": [
    {
      "language": {
        "name": "fr",
        "url": "https://pokeapi.co/api/v2/language/5/"
      },
      "name": "Kanto"
    },
    {
      "language": {
        "name": "de",
        "url": "https://pokeapi.co/api/v2/language/6/"
      },
      "name": "Kanto"
    },
    {
      "language": {
        "name": "en",
        "url": "https://pokeapi.co/api/v2/language/9/"
      },
      "name": "Kanto"
    }
  ],
  "pokemon_entries": [
    {
      "entry_number": 1,
      "pokemon_species": {
        "name": "bulbasaur",
        "url": "https://pokeapi.co/api/v2/pokemon-species/1/"
      }
    },
    {
      "entry_number": 2,
      "pokemon_species": {
        "name": "ivysaur",
        "url": "https://pokeapi.co/api/v2/pokemon-species/2/"
      }
    },
    {
      "entry_number": 3,
      "pokemon_species": {
        "name": "venusaur",
        "url": "https://pokeapi.co/api/v2/pokemon-species/3/"
      }
    },
    {
      "entry_number": 4,
      "pokemon_species": {
        "name": "charmander",
        "url": "https://pokeapi.co/api/v2/pokemon-species/4/"
      }
    },
    {
      "entry_number": 5,
      "pokemon_species": {
        "name": "charmeleon",
        "url": "https://pokeapi.co/api/v2/pokemon-species/5/"
      }
    },
    {
      "entry_number": 6,
      "pokemon_species": {
        "name": "charizard",
        "url": "https://pokeapi.co/api/v2/pokemon-species/6/"
      }
    },
    {
      "entry_number": 7,
      "pokemon_species": {
        "name": "squirtle",
        "url": "https://pokeapi.co/api/v2/pokemon-species/7/"
      }
    },
    {
      "entry_number": 8,
      "pokemon_species": {
        "name": "wartortle",
        "url": "https://pokeapi.co/api/v2/pokemon-species/8/"
      }
    },
    {
      "entry_number": 9,
      "pokemon_species": {
        "name": "blastoise",
        "url": "https://pokeapi.co/api/v2/pokemon-species/9/"
      }
    },
    {
      "entry_number": 10,
      "pokemon_species": {
        "name": "caterpie",
        "url": "https://pokeapi.co/api/v2/pokemon-species/10/"
      }
    },
    {
      "entry_number": 11,
      "pokemon_species": {
        "name": "metapod",
        "url": "https://pokeapi.co/api/v2/pokemon-species/11/"
      }
    },
    {
      "entry_number": 12,
      "pokemon_species": {
        "name": "butterfree",
        "url": "https://pokeapi.co/api/v2/pokemon-species/12/"
      }
    },
    {
      "entry_number": 13,
      "pokemon_species": {
        "name": "weedle",
        "url": "https://pokeapi.co/api/v2/pokemon-species/13/"
      }
    },
    {
      "entry_number": 14,
      "pokemon_species": {
        "name": "kakuna",
        "url": "https://pokeapi.co/api/v2/pokemon-species/14/"
      }
    },
    {
      "entry_number": 15,
      "pokemon_species": {
        "name": "beedrill",
        "url": "https://pokeapi.co/api/v2/pokemon-species/15/"
      }
    },
    {
      "entry_number": 16,
      "pokemon_species": {
        "name": "pidgey",
        "url": "https://pokeapi.co/api/v2/pokemon-species/16/"
      }
    },
    {
      "entry_number": 17,
      "pokemon_species": {
        "name": "pidgeotto",
        "url": "https://pokeapi.co/api/v2/pokemon-species/17/"
      }
    },
    {
      "entry_number": 18,
      "pokemon_species": {
        "name": "pidgeot",
        "url": "https://pokeapi.co/api/v2/pokemon-species/18/"
      }
    },
    {
      "entry_number": 19,
      "pokemon_species": {
        "name": "rattata",
        "url": "https://pokeapi.co/api/v2/pokemon-species/19/"
      }
    },
    {
      "entry_number": 20,
      "pokemon_species": {
        "name": "raticate",
        "url": "https://pokeapi.co/api/v2/pokemon-species/20/"
      }
    },
    {
      "entry_number": 21,
      "pokemon_species": {
        "name": "spearow",
        "url": "https://pokeapi.co/api/v2/pokemon-species/21/"
      }
    },
    {
      "entry_number": 22,
      "pokemon_species": {
        "name": "fearow",
        "url": "https://pokeapi.co/api/v2/pokemon-species/22/"
      }
    },
    {
      "entry_number": 23,
      "pokemon_species": {
        "name": "ekans",
        "url": "https://pokeapi.co/api/v2/pokemon-species/23/"
      }
    },
    {
      "entry_number": 24,
      "pokemon_species": {
        "name": "arbok",
        "url": "https://pokeapi.co/api/v2/pokemon-species/24/"
      }
    },
    {
      "entry_number": 25,
      "pokemon_species": {
        "name": "pikachu",
        "url": "https://pokeapi.co/api/v2/pokemon-species/25/"
      }
    },
    {
      "entry_number": 26,
      "pokemon_species": {
        "name": "raichu",
        "url": "https://pokeapi.co/api/v2/pokemon-species/26/"
      }
    },
    {
      "entry_number": 27,
      "pokemon_species": {
        "name": "sandshrew",
        "url": "https://pokeapi.co/api/v2/pokemon-species/27/"
      }
    },
    {
      "entry_number": 28,
      "pokemon_species": {
        "name": "sandslash",
        "url": "https://pokeapi.co/api/v2/pokemon-species/28/"
      }
    },
    {
      "entry_number": 29,
      "pokemon_species": {
        "name": "nidoran-f",
        "url": "https://pokeapi.co/api/v2/pokemon-species/29/"
      }
    },
    {
      "entry_number": 30,
      "pokemon_species": {
        "name": "nidorina",
        "url": "https://pokeapi.co/api/v2/pokemon-species/30/"
      }
    },
    {
      "entry_number": 31,
      "pokemon_species": {
        "name": "nidoqueen",
        "url": "https://pokeapi.co/api/v2/pokemon-species/31/"
      }
    },
    {
      "entry_number": 32,
      "pokemon_species": {
        "name": "nidoran-m",
        "url": "https://pokeapi.co/api/v2/pokemon-species/32/"
      }
    },
    {
      "entry_number": 33,
      "pokemon_species": {
        "name": "nidorino",
        "url": "https://pokeapi.co/api/v2/pokemon-species/33/"
      }
    },
    {
      "entry_number": 34,
      "pokemon_species": {
        "name": "nidoking",
        "url": "https://pokeapi.co/api/v2/pokemon-species/34/"
      }
    },
    {
      "entry_number": 35,
      "pokemon_species": {
        "name": "clefairy",
        "url": "https://pokeapi.co/api/v2/pokemon-species/35/"
      }
    },
    {
      "entry_number": 36,
      "pokemon_species": {
        "name": "clefable",
        "url": "https://pokeapi.co/api/v2/pokemon-species/36/"
      }
    },
    {
      "entry_number": 37,
      "pokemon_species": {
        "name": "vulpix",
        "url": "https://pokeapi.co/api/v2/pokemon-species/37/"
      }
    },
    {
      "entry_number": 38,
      "pokemon_species": {
        "name": "ninetales",
        "url": "https://pokeapi.co/api/v2/pokemon-species/38/"
      }
    },
    {
      "entry_number": 39,
      "pokemon_species": {
        "name": "jigglypuff",
        "url": "https://pokeapi.co/api/v2/pokemon-species/39/"
      }
    },
    {
      "entry_number": 40,
      "pokemon_species": {
        "name": "wigglytuff",
        "url": "https://pokeapi.co/api/v2/pokemon-species/40/"
      }
    },
    {
      "entry_number": 41,
      "pokemon_species": {
        "name": "zubat",
        "url": "https://pokeapi.co/api/v2/pokemon-species/41/"
      }
    },
    {
      "entry_number": 42,
      "pokemon_species": {
        "name": "golbat",
        "url": "https://pokeapi.co/api/v2/pokemon-species/42/"
      }
    },
    {
      "entry_number": 43,
      "pokemon_species": {
        "name": "oddish",
        "url": "https://pokeapi.co/api/v2/pokemon-species/43/"
      }
    },
    {
      "entry_number": 44,
      "pokemon_species": {
        "name": "gloom",
        "url": "https://pokeapi.co/api/v2/pokemon-species/44/"
      }
    },
    {
      "entry_number": 45,
      "pokemon_species": {
        "name": "vileplume",
        "url": "https://pokeapi.co/api/v2/pokemon-species/45/"
      }
    },
    {
      "entry_number": 46,
      "pokemon_species": {
        "name": "paras",
        "url": "https://pokeapi.co/api/v2/pokemon-species/46/"
      }
    },
    {
      "entry_number": 47,
      "pokemon_species": {
        "name": "parasect",
        "url": "https://pokeapi.co/api/v2/pokemon-species/47/"
      }
    },
    {
      "entry_number": 48,
      "pokemon_species": {
        "name": "venonat",
        "url": "https://pokeapi.co/api/v2/pokemon-species/48/"
      }
    },
    {
      "entry_number": 49,
      "pokemon_species": {
        "name": "venomoth",
        "url": "https://pokeapi.co/api/v2/pokemon-species/49/"
      }
    },
    {
      "entry_number": 50,
      "pokemon_species": {
        "name": "diglett",
        "url": "https://pokeapi.co/api/v2/pokemon-species/50/"
      }
    },
    {
      "entry_number": 51,
      "pokemon_species": {
        "name": "dugtrio",
        "url": "https://pokeapi.co/api/v2/pokemon-species/51/"
      }
    },
    {
      "entry_number": 52,
      "pokemon_species": {
        "name": "meowth",
        "url": "https://pokeapi.co/api/v2/pokemon-species/52/"
      }
    },
    {
      "entry_number": 53,
      "pokemon_species": {
        "name": "persian",
        "url": "https://pokeapi.co/api/v2/pokemon-species/53/"
      }
    },
    {
      "entry_number": 54,
      "pokemon_species": {
        "name": "psyduck",
        "url": "https://pokeapi.co/api/v2/pokemon-species/54/"
      }
    },
    {
      "entry_number": 55,
      "pokemon_species": {
        "name": "golduck",
        "url": "https://pokeapi.co/api/v2/pokemon-species/55/"
      }
    },
    {
      "entry_number": 56,
      "pokemon_species": {
        "name": "mankey",
        "url": "https://pokeapi.co/api/v2/pokemon-species/56/"
      }
    },
    {
      "entry_number": 57,
      "pokemon_species": {
        "name": "primeape",
        "url": "https://pokeapi.co/api/v2/pokemon-species/57/"
      }
    },
    {
      "entry_number": 58,
      "pokemon_species": {
        "name": "growlithe",
        "url": "https://pokeapi.co/api/v2/pokemon-species/58/"
      }
    },
    {
      "entry_number": 59,
      "pokemon_species": {
        "name": "arcanine",
        "url": "https://pokeapi.co/api/v2/pokemon-species/59/"
      }
    },
    {
      "entry_number": 60,
      "pokemon_species": {
        "name": "poliwag",
        "url": "https://pokeapi.co/api/v2/pokemon-species/60/"
      }
    },
    {
      "entry_number": 61,
      "pokemon_species": {
        "name": "poliwhirl",
        "url": "https://pokeapi.co/api/v2/pokemon-species/61/"
      }
    },
    {
      "entry_number": 62,
      "pokemon_species": {
        "name": "poliwrath",
        "url": "https://pokeapi.co/api/v2/pokemon-species/62/"
      }
    },
    {
      "entry_number": 63,
      "pokemon_species": {
        "name": "abra",
        "url": "https://pokeapi.co/api/v2/pokemon-species/63/"
      }
    },
    {
      "entry_number": 64,
      "pokemon_species": {
        "name": "kadabra",
        "url": "https://pokeapi.co/api/v2/pokemon-species/64/"
      }
    },
    {
      "entry_number": 65,
      "pokemon_species": {
        "name": "alakazam",
        "url": "https://pokeapi.co/api/v2/pokemon-species/65/"
      }
    },
    {
      "entry_number": 66,
      "pokemon_species": {
        "name": "machop",
        "url": "https://pokeapi.co/api/v2/pokemon-species/66/"
      }
    },
    {
      "entry_number": 67,
      "pokemon_species": {
        "name": "machoke",
        "url": "https://pokeapi.co/api/v2/pokemon-species/67/"
      }
    },
    {
      "entry_number": 68,
      "pokemon_species": {
        "name": "machamp",
        "url": "https://pokeapi.co/api/v2/pokemon-species/68/"
      }
    },
    {
      "entry_number": 69,
      "pokemon_species": {
        "name": "bellsprout",
        "url": "https://pokeapi.co/api/v2/pokemon-species/69/"
      }
    },
    {
      "entry_number": 70,
      "pokemon_species": {
        "name": "weepinbell",
        "url": "https://pokeapi.co/api/v2/pokemon-species/70/"
      }
    },
    {
      "entry_number": 71,
      "pokemon_species": {
        "name": "victreebel",
        "url": "https://pokeapi.co/api/v2/pokemon-species/71/"
      }
    },
    {
      "entry_number": 72,
      "pokemon_species": {
        "name": "tentacool",
        "url": "https://pokeapi.co/api/v2/pokemon-species/72/"
      }
    },
    {
      "entry_number": 73,
      "pokemon_species": {
        "name": "tentacruel",
        "url": "https://pokeapi.co/api/v2/pokemon-species/73/"
      }
    },
    {
      "entry_number": 74,
      "pokemon_species": {
        "name": "geodude",
        "url": "https://pokeapi.co/api/v2/pokemon-species/74/"
      }
    },
    {
      "entry_number": 75,
      "pokemon_species": {
        "name": "graveler",
        "url": "https://pokeapi.co/api/v2/pokemon-species/75/"
      }
    },
    {
      "entry_number": 76,
      "pokemon_species": {
        "name": "golem",
        "url": "https://pokeapi.co/api/v2/pokemon-species/76/"
      }
    },
    {
      "entry_number": 77,
      "pokemon_species": {
        "name": "ponyta",
        "url": "https://pokeapi.co/api/v2/pokemon-species/77/"
      }
    },
    {
      "entry_number": 78,
      "pokemon_species": {
        "name": "rapidash",
        "url": "https://pokeapi.co/api/v2/pokemon-species/78/"
      }
    },
    {
      "entry_number": 79,
      "pokemon_species": {
        "name": "slowpoke",
        "url": "https://pokeapi.co/api/v2/pokemon-species/79/"
      }
    },
    {
      "entry_number": 80,
      "pokemon_species": {
        "name": "slowbro",
        "url": "https://pokeapi.co/api/v2/pokemon-species/80/"
      }
    },
    {
      "entry_number": 81,
      "pokemon_species": {
        "name": "magnemite",
        "url": "https://pokeapi.co/api/v2/pokemon-species/81/"
      }
    },
    {
      "entry_number": 82,
      "pokemon_species": {
        "name": "magneton",
        "url": "https://pokeapi.co/api/v2/pokemon-species/82/"
      }
    },
    {
      "entry_number": 83,
      "pokemon_species": {
        "name": "farfetchd",
        "url": "https://pokeapi.co/api/v2/pokemon-species/83/"
      }
    },
    {
      "entry_number": 84,
      "pokemon_species": {
        "name": "doduo",
        "url": "https://pokeapi.co/api/v2/pokemon-species/84/"
      }
    },
    {
      "entry_number": 85,
      "pokemon_species": {
        "name": "dodrio",
        "url": "https://pokeapi.co/api/v2/pokemon-species/85/"
      }
    },
    {
      "entry_number": 86,
      "pokemon_species": {
        "name": "seel",
        "url": "https://pokeapi.co/api/v2/pokemon-species/86/"
      }
    },
    {
      "entry_number": 87,
      "pokemon_species": {
        "name": "dewgong",
        "url": "https://pokeapi.co/api/v2/pokemon-species/87/"
      }
    },
    {
      "entry_number": 88,
      "pokemon_species": {
        "name": "grimer",
        "url": "https://pokeapi.co/api/v2/pokemon-species/88/"
      }
    },
    {
      "entry_number": 89,
      "pokemon_species": {
        "name": "muk",
        "url": "https://pokeapi.co/api/v2/pokemon-species/89/"
      }
    },
    {
      "entry_number": 90,
      "pokemon_species": {
        "name": "shellder",
        "url": "https://pokeapi.co/api/v2/pokemon-species/90/"
      }
    },
    {
      "entry_number": 91,
      "pokemon_species": {
        "name": "cloyster",
        "url": "https://pokeapi.co/api/v2/pokemon-species/91/"
      }
    },
    {
      "entry_number": 92,
      "pokemon_species": {
        "name": "gastly",
        "url": "https://pokeapi.co/api/v2/pokemon-species/92/"
      }
    },
    {
      "entry_number": 93,
      "pokemon_species": {
        "name": "haunter",
        "url": "https://pokeapi.co/api/v2/pokemon-species/93/"
      }
    },
    {
      "entry_number": 94,
      "pokemon_species": {
        "name": "gengar",
        "url": "https://pokeapi.co/api/v2/pokemon-species/94/"
      }
    },
    {
      "entry_number": 95,
      "pokemon_species": {
        "name": "onix",
        "url": "https://pokeapi.co/api/v2/pokemon-species/95/"
      }
    },
    {
      "entry_number": 96,
      "pokemon_species": {
        "name": "drowzee",
        "url": "https://pokeapi.co/api/v2/pokemon-species/96/"
      }
    },
    {
      "entry_number": 97,
      "pokemon_species": {
        "name": "hypno",
        "url": "https://pokeapi.co/api/v2/pokemon-species/97/"
      }
    },
    {
      "entry_number": 98,
      "pokemon_species": {
        "name": "krabby",
        "url": "https://pokeapi.co/api/v2/pokemon-species/98/"
      }
    },
    {
      "entry_number": 99,
      "pokemon_species": {
        "name": "kingler",
        "url": "https://pokeapi.co/api/v2/pokemon-species/99/"
      }
    },
    {
      "entry_number": 100,
      "pokemon_species": {
        "name": "voltorb",
        "url": "https://pokeapi.co/api/v2/pokemon-species/100/"
      }
    },
    {
      "entry_number": 101,
      "pokemon_species": {
        "name": "electrode",
        "url": "https://pokeapi.co/api/v2/pokemon-species/101/"
      }
    },
    {
      "entry_number": 102,
      "pokemon_species": {
        "name": "exeggcute",
        "url": "https://pokeapi.co/api/v2/pokemon-species/102/"
      }
    },
    {
      "entry_number": 103,
      "pokemon_species": {
        "name": "exeggutor",
        "url": "https://pokeapi.co/api/v2/pokemon-species/103/"
      }
    },
    {
      "entry_number": 104,
      "pokemon_species": {
        "name": "cubone",
        "url": "https://pokeapi.co/api/v2/pokemon-species/104/"
      }
    },
    {
      "entry_number": 105,
      "pokemon_species": {
        "name": "marowak",
        "url": "https://pokeapi.co/api/v2/pokemon-species/105/"
      }
    },
    {
      "entry_number": 106,
      "pokemon_species": {
        "name": "hitmonlee",
        "url": "https://pokeapi.co/api/v2/pokemon-species/106/"
      }
    },
    {
      "entry_number": 107,
      "pokemon_species": {
        "name": "hitmonchan",
        "url": "https://pokeapi.co/api/v2/pokemon-species/107/"
      }
    },
    {
      "entry_number": 108,
      "pokemon_species": {
        "name": "lickitung",
        "url": "https://pokeapi.co/api/v2/pokemon-species/108/"
      }
    },
    {
      "entry_number": 109,
      "pokemon_species": {
        "name": "koffing",
        "url": "https://pokeapi.co/api/v2/pokemon-species/109/"
      }
    },
    {
      "entry_number": 110,
      "pokemon_species": {
        "name": "weezing",
        "url": "https://pokeapi.co/api/v2/pokemon-species/110/"
      }
    },
    {
      "entry_number": 111,
      "pokemon_species": {
        "name": "rhyhorn",
        "url": "https://pokeapi.co/api/v2/pokemon-species/111/"
      }
    },
    {
      "entry_number": 112,
      "pokemon_species": {
        "name": "rhydon",
        "url": "https://pokeapi.co/api/v2/pokemon-species/112/"
      }
    },
    {
      "entry_number": 113,
      "pokemon_species": {
        "name": "chansey",
        "url": "https://pokeapi.co/api/v2/pokemon-species/113/"
      }
    },
    {
      "entry_number": 114,
      "pokemon_species": {
        "name": "tangela",
        "url": "https://pokeapi.co/api/v2/pokemon-species/114/"
      }
    },
    {
      "entry_number": 115,
      "pokemon_species": {
        "name": "kangaskhan",
        "url": "https://pokeapi.co/api/v2/pokemon-species/115/"
      }
    },
    {
      "entry_number": 116,
      "pokemon_species": {
        "name": "horsea",
        "url": "https://pokeapi.co/api/v2/pokemon-species/116/"
      }
    },
    {
      "entry_number": 117,
      "pokemon_species": {
        "name": "seadra",
        "url": "https://pokeapi.co/api/v2/pokemon-species/117/"
      }
    },
    {
      "entry_number": 118,
      "pokemon_species": {
        "name": "goldeen",
        "url": "https://pokeapi.co/api/v2/pokemon-species/118/"
      }
    },
    {
      "entry_number": 119,
      "pokemon_species": {
        "name": "seaking",
        "url": "https://pokeapi.co/api/v2/pokemon-species/119/"
      }
    },
    {
      "entry_number": 120,
      "pokemon_species": {
        "name": "staryu",
        "url": "https://pokeapi.co/api/v2/pokemon-species/120/"
      }
    },
    {
      "entry_number": 121,
      "pokemon_species": {
        "name": "starmie",
        "url": "https://pokeapi.co/api/v2/pokemon-species/121/"
      }
    },
    {
      "entry_number": 122,
      "pokemon_species": {
        "name": "mr-mime",
        "url": "https://pokeapi.co/api/v2/pokemon-species/122/"
      }
    },
    {
      "entry_number": 123,
      "pokemon_species": {
        "name": "scyther",
        "url": "https://pokeapi.co/api/v2/pokemon-species/123/"
      }
    },
    {
      "entry_number": 124,
      "pokemon_species": {
        "name": "jynx",
        "url": "https://pokeapi.co/api/v2/pokemon-species/124/"
      }
    },
    {
      "entry_number": 125,
      "pokemon_species": {
        "name": "electabuzz",
        "url": "https://pokeapi.co/api/v2/pokemon-species/125/"
      }
    },
    {
      "entry_number": 126,
      "pokemon_species": {
        "name": "magmar",
        "url": "https://pokeapi.co/api/v2/pokemon-species/126/"
      }
    },
    {
      "entry_number": 127,
      "pokemon_species": {
        "name": "pinsir",
        "url": "https://pokeapi.co/api/v2/pokemon-species/127/"
      }
    },
    {
      "entry_number": 128,
      "pokemon_species": {
        "name": "tauros",
        "url": "https://pokeapi.co/api/v2/pokemon-species/128/"
      }
    },
    {
      "entry_number": 129,
      "pokemon_species": {
        "name": "magikarp",
        "url": "https://pokeapi.co/api/v2/pokemon-species/129/"
      }
    },
    {
      "entry_number": 130,
      "pokemon_species": {
        "name": "gyarados",
        "url": "https://pokeapi.co/api/v2/pokemon-species/130/"
      }
    },
    {
      "entry_number": 131,
      "pokemon_species": {
        "name": "lapras",
        "url": "https://pokeapi.co/api/v2/pokemon-species/131/"
      }
    },
    {
      "entry_number": 132,
      "pokemon_species": {
        "name": "ditto",
        "url": "https://pokeapi.co/api/v2/pokemon-species/132/"
      }
    },
    {
      "entry_number": 133,
      "pokemon_species": {
        "name": "eevee",
        "url": "https://pokeapi.co/api/v2/pokemon-species/133/"
      }
    },
    {
      "entry_number": 134,
      "pokemon_species": {
        "name": "vaporeon",
        "url": "https://pokeapi.co/api/v2/pokemon-species/134/"
      }
    },
    {
      "entry_number": 135,
      "pokemon_species": {
        "name": "jolteon",
        "url": "https://pokeapi.co/api/v2/pokemon-species/135/"
      }
    },
    {
      "entry_number": 136,
      "pokemon_species": {
        "name": "flareon",
        "url": "https://pokeapi.co/api/v2/pokemon-species/136/"
      }
    },
    {
      "entry_number": 137,
      "pokemon_species": {
        "name": "porygon",
        "url": "https://pokeapi.co/api/v2/pokemon-species/137/"
      }
    },
    {
      "entry_number": 138,
      "pokemon_species": {
        "name": "omanyte",
        "url": "https://pokeapi.co/api/v2/pokemon-species/138/"
      }
    },
    {
      "entry_number": 139,
      "pokemon_species": {
        "name": "omastar",
        "url": "https://pokeapi.co/api/v2/pokemon-species/139/"
      }
    },
    {
      "entry_number": 140,
      "pokemon_species": {
        "name": "kabuto",
        "url": "https://pokeapi.co/api/v2/pokemon-species/140/"
      }
    },
    {
      "entry_number": 141,
      "pokemon_species": {
        "name": "kabutops",
        "url": "https://pokeapi.co/api/v2/pokemon-species/141/"
      }
    },
    {
      "entry_number": 142,
      "pokemon_species": {
        "name": "aerodactyl",
        "url": "https://pokeapi.co/api/v2/pokemon-species/142/"
      }
    },
    {
      "entry_number": 143,
      "pokemon_species": {
        "name": "snorlax",
        "url": "https://pokeapi.co/api/v2/pokemon-species/143/"
      }
    },
    {
      "entry_number": 144,
      "pokemon_species": {
        "name": "articuno",
        "url": "https://pokeapi.co/api/v2/pokemon-species/144/"
      }
    },
    {
      "entry_number": 145,
      "pokemon_species": {
        "name": "zapdos",
        "url": "https://pokeapi.co/api/v2/pokemon-species/145/"
      }
    },
    {
      "entry_number": 146,
      "pokemon_species": {
        "name": "moltres",
        "url": "https://pokeapi.co/api/v2/pokemon-species/146/"
      }
    },
    {
      "entry_number": 147,
      "pokemon_species": {
        "name": "dratini",
        "url": "https://pokeapi.co/api/v2/pokemon-species/147/"
      }
    },
    {
      "entry_number": 148,
      "pokemon_species": {
        "name": "dragonair",
        "url": "https://pokeapi.co/api/v2/pokemon-species/148/"
      }
    },
    {
      "entry_number": 149,
      "pokemon_species": {
        "name": "dragonite",
        "url": "https://pokeapi.co/api/v2/pokemon-species/149/"
      }
    },
    {
      "entry_number": 150,
      "pokemon_species": {
        "name": "mewtwo",
        "url": "https://pokeapi.co/api/v2/pokemon-species/150/"
      }
    },
    {
      "entry_number": 151,
      "pokemon_species": {
        "name": "mew",
        "url": "https://pokeapi.co/api/v2/pokemon-species/151/"
      }
    }
  ],
  "region": {
    "name": "kanto",
    "url": "https://pokeapi.co/api/v2/region/1/"
  },
  "version_groups": [
    {
      "name": "red-blue",
      "url": "https://pokeapi.co/api/v2/version-group/1/"
    },
    {
      "name": "yellow",
      "url": "https://pokeapi.co/api/v2/version-group/2/"
    },
    {
      "name": "firered-leafgreen",
      "url": "https://pokeapi.co/api/v2/version-group/7/"
    }
  ]
}
*/