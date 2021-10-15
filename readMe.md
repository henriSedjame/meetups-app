
#### Init an EdgeDB Project
`edgedb project init`

#### Createmigrations

After define database model in file <i>dbschema/schema.esdl</i> run:

`edgedb -I meetups_app create-migration`

`edgedb -I meetups_app migrate`
