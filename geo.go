package filespot

// Geo represents platformcraft `geo` type
// It's a map of `map[continent]map[country]true`
// Keys `continent` and `country` are ISO codes of continents and counties (e.g, EU и RU).
// All keys MUST be capitalized.
// When `country` was set to "ALL" it grants access to whole world.
// For example:
//      grants acccess to Europe and North America
//        &Geo{
//            "EU": {
//                "ALL": true,
//            },
//            "NA": {
//                "ALL": true,
//            },
//        }
//      grants access only to Russia
//        &Geo{
//            "EU": {
//                "RU": true,
//            },
//        }
type Geo map[string]map[string]bool
