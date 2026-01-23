# changes to import gpx categories properly

## master data

The stored master data table categories gets two companions:

- categories_i18n for the names in a specific language
- categories_aliases for the altzernative names in different gpx variants

The translated names of such masterdata has to be part of the master data.
Otherwise it would be a hard job to provide proper translations when some instance add categories or tries to rename them. We store this in categories_i18n.

Since there are no well defined values for the "type" of a gps track a lot of different places and names exist in different gpx variants.

In the master data we cannot tell between the different places but we can provide a list of alternative names to  detect the proper category for wanderer.
We do not tell between different languages here since the they can come from very different sources even in foreign languages.

To populate these tables we provide a json file with the data. 
So every admin can provide a set of categories, its translations and aliases to fits his/her needs.

Ideally it would be possible to reread such a config file to enhance the given config.

An example of such a file:

```json
{
    "categories": [
        {
            "defaultName": "hiking",
            "icon": "hm, lets try",
            "i18n": 
            {
                "de": "Wandern",
                "en": "hiking"
            },
            "aliases": [
                "Wandern",
                "hiking",
                "Bergwandern",
                "..."
            ]
        }
    ]
}
```

## components

This master data is stored and maintained in pocketbase.
It can easily be maintained through the pocketbase web interface.

The web component can access the data through API.

In general I prefer to parse the gpx file in pocketbase since it is a business function and not view operation.
Let's see how far I can go here.

## changes in db

### migrations

- create a migration to the new structure

### import

- on startup: if no categories given in db, read the config file and fill the tables.

- optional: define a command to reimport the config file

## changes in web

- when parsing a gpx file try to find the proper place for the type and use DB-API to get the category-id