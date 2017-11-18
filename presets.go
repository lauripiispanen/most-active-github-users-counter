package main

type PresetLocations []string

var PRESETS = map[string]PresetLocations{
  "finland":PresetLocations{"finland", "suomi", "helsinki", "tampere", "oulu", "espoo", "vantaa", "turku"},
  "sweden":PresetLocations{"sweden", "sverige", "stockholm", "malm%C3%B6", "uppsala", "g%C3%B6teborg", "gothenburg"},
  "denmark":PresetLocations{"denmark", "danmark", "copenhagen", "københavn"},
  "norway":PresetLocations{"norway", "norge", "oslo", "bergen"},
  "germany":PresetLocations{"germany", "deutschland", "berlin", "frankfurt", "munich", "m%C3%BCnchen", "hamburg", "cologne", "k%C3%B6ln"},
  "netherlands":PresetLocations{"netherlands", "nederland", "amsterdam", "rotterdam", "hague", "utrecht", "holland"},
  "ukraine":PresetLocations{"ukraine", "kiev", "kharkiv", "dnipro", "odesa", "donetsk", "zaporizhia"},
  "japan":PresetLocations{"japan", "tokyo", "yokohama", "osaka", "nagoya", "sapporo", "kobe", "kyoto", "fukuoka", "kawasaki", "saitama", "hiroshima", "sendai"},
  "russia":PresetLocations{"russia", "moscow", "saint%2Bpetersburg", "novosibirsk", "yekaterinburg", "nizhny%2Bnovgorod", "samara", "omsk", "kazan", "chelyabinsk", "rostov-on-don", "ufa", "volgograd"},
  "estonia":PresetLocations{"estonia", "eesti", "tallinn", "tartu", "narva", "p%C3%A4rnu"}}

func Preset(name string) []string {
  return PRESETS[name]
}
