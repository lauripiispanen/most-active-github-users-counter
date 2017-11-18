package main

type PresetLocations []string

var PRESETS = map[string]PresetLocations{
  "finland":PresetLocations{"finland", "suomi", "helsinki", "tampere", "oulu", "espoo", "vantaa", "turku"},
  "sweden":PresetLocations{"sweden", "sverige", "stockholm", "malm%C3%B6", "uppsala", "g%C3%B6teborg", "gothenburg"},
  "norway":PresetLocations{"norway", "norge", "oslo", "bergen"},
  "germany":PresetLocations{"germany", "deutschland", "berlin", "frankfurt", "munich", "m%C3%BCnchen", "hamburg", "cologne", "k%C3%B6ln"},
  "netherlands":PresetLocations{"netherlands", "nederland", "amsterdam", "rotterdam", "hague", "utrecht", "holland"},
  "ukraine":PresetLocations{"ukraine", "kiev", "kharkiv", "dnipro", "odesa", "donetsk", "zaporizhia"},
  "japan":PresetLocations{"japan", "tokyo", "yokohama", "osaka", "nagoya", "sapporo", "kobe", "kyoto", "fukuoka", "kawasaki", "saitama", "hiroshima", "sendai"},
  "russia":PresetLocations{"russia", "moscow", "saint%2Bpetersburg", "novosibirsk", "yekaterinburg", "nizhny%2Bnovgorod", "samara", "omsk", "kazan", "chelyabinsk", "rostov-on-don", "ufa", "volgograd"},
  "estonia":PresetLocations{"estonia", "eesti", "tallinn", "tartu", "narva", "p%C3%A4rnu"},
  "denmark":PresetLocations{"denmark", "danmark", "copenhagen","aarhus","odense","aalborg"},
  "france":PresetLocations{"france","paris","marseille","lyon","toulouse","nice","nantes","strasbourg","montpellier","bordeaux","lille","rennes","reims"},
  "spain":PresetLocations{"spain","espa%C3%B1a","madrid","barcelona","valencia","seville","sevilla","zaragoza","malaga","murcia","palma","bilbao","alicante","cordoba"},
  "italy":PresetLocations{"italy","italia","rome","roma","milan","naples","napoli","turin","torino","palermo","genoa","genova","bologna","florence","firenze","bari","catania","venice","verona"},
  "uk": PresetLocations{"uk","london","birmingham","leeds","glasgow","sheffield","bradford","manchester","edinburgh","liverpool","bristol","cardiff","belfast","leicester","wakefield","coventry","nottingham","newcastle"}}

func Preset(name string) []string {
  return PRESETS[name]
}
