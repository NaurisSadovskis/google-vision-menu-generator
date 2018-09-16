# Google Vision API Menu Generator

This calls Google Vision API using local image file and obtains a list of labels. It then uses top label to construct a menu.

**Note**: First 1'000 API calls are free, you are charged $1.50 per 1'000 calls afterwards. 

## Build & Run

Initialise dependencies:
```
$ make dep
```

Set your credentials file:
```
$ export GOOGLE_APPLICATION_CREDENTIALS=/path/to/credentials.json
```

## Run
```
$ make run
go run generate-menu.go
***Google Vision API results***
map[151_BenR_170130.jpg:[olive ingredient bead jewelry making] Chicken-Pad-Thai_0.jpg:[pad thai noodle taglierini thai food chow mein asian food lo mein] banana.jpg:[banana banana family yellow cooking plantain peel] sabor-aroma-central.jpg:[prosciutto jamón serrano bayonne ham bresaola cold cut salt cured meat smoked salmon turkey ham]]

***Today's Menu Consists Of***
Todays menu consists of:
* pad thai
* banana
* prosciutto
* olive
```

## TODO: 
* Add a web UI
* Add more generic words to exclusions (`food`, `cuisine`)
* Add data persistance to limit API calls
* Add Cloud Storage integration