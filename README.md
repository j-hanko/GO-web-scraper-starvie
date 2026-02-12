# Starvie Racket Scraper (Go)

A simple Go scraper that collects **Starvie** padel racket data from **starvie.com** and outputs it to JSON files.

## What it does
- Visits the Starvie main page (`/en`) and collects available racket **series/lines** (menu items containing `"line"`)
- For each series page (collection), it extracts:
    - `brand`, `model`, `price`, `imageUrl`, `racketPage`, `series`
- Opens each racket detail page and parses (from the product features table):
    - `weight`, `shape`, `material` (`Surface`)
- Writes results to:
    - `StarvieRacket<Series>.json` (e.g. `StarvieRacketSuperProLine.json`)

## Output format
Each item looks like:
```json
{
  "brand": "Starvie",
  "model": "BlackTitan",
  "price": "€320,00",
  "imageUrl": "starvie.com/cdn/shop/files/Black_Titan_pala_padel_STARVIE.MAIN.png",
  "racketPage": "https://starvie.com/en/products/black-titan",
  "weight": "358 (+/- 8 g)",
  "shape": "Hybrid",
  "material": "24K Carbon",
  "series": "SuperProLine"
}