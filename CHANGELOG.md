# Changelog
All notable changes to this project will be documented in this file.

## 0.3.0 - 2021-12-23

### 🚀 Features

- Clusters for scenes ([#75](https://github.com/reearth/reearth-backend/pull/75)) [`3512c0`](https://github.com/reearth/reearth-backend/commit/3512c0)
- Add fields of scene property for terrain [`8693b4`](https://github.com/reearth/reearth-backend/commit/8693b4)
- Camera limiter  ([#87](https://github.com/reearth/reearth-backend/pull/87)) [`63c582`](https://github.com/reearth/reearth-backend/commit/63c582)

### 🔧 Bug Fixes

- Terrain fields of scene property [`5e3d25`](https://github.com/reearth/reearth-backend/commit/5e3d25)
- Numbers are not decoded from gql to value [`2ddbc8`](https://github.com/reearth/reearth-backend/commit/2ddbc8)
- Layers have their own tags separate from the scene ([#90](https://github.com/reearth/reearth-backend/pull/90)) [`c4fb9a`](https://github.com/reearth/reearth-backend/commit/c4fb9a)
- Return property with clusters data ([#89](https://github.com/reearth/reearth-backend/pull/89)) [`1b99c6`](https://github.com/reearth/reearth-backend/commit/1b99c6)
- Cast values, rename value.OptionalValue ([#93](https://github.com/reearth/reearth-backend/pull/93)) [`ba4b18`](https://github.com/reearth/reearth-backend/commit/ba4b18)
- Synchronize mongo migration ([#94](https://github.com/reearth/reearth-backend/pull/94)) [`db4cea`](https://github.com/reearth/reearth-backend/commit/db4cea)

### 📖 Documentation

- Add pkg.go.dev badge to readme [`91f9b3`](https://github.com/reearth/reearth-backend/commit/91f9b3)

### ✨ Refactor

- Make property.Value and dataset.Value independent in pkg/value ([#77](https://github.com/reearth/reearth-backend/pull/77)) [`73143b`](https://github.com/reearth/reearth-backend/commit/73143b)

### Miscellaneous Tasks

- Fix plugin manifest JSON schema [`2b57b1`](https://github.com/reearth/reearth-backend/commit/2b57b1)

## 0.2.0 - 2021-11-18

### 🚀 Features

- Support opentelemetry ([#68](https://github.com/reearth/reearth-backend/pull/68)) [`25c581`](https://github.com/reearth/reearth-backend/commit/25c581)

### 🔧 Bug Fixes

- Add an index to mongo project collection to prevent creating projects whose alias is duplicated [`10f745`](https://github.com/reearth/reearth-backend/commit/10f745)
- Check project alias duplication on project update [`443f2c`](https://github.com/reearth/reearth-backend/commit/443f2c)

### ✨ Refactor

- Add PropertySchemaGroupID to pkg/id ([#70](https://github.com/reearth/reearth-backend/pull/70)) [`9ece9e`](https://github.com/reearth/reearth-backend/commit/9ece9e)

### Miscellaneous Tasks

- Fix typo in github actions [`4a9dc5`](https://github.com/reearth/reearth-backend/commit/4a9dc5)
- Clean up unused code [`b5b01b`](https://github.com/reearth/reearth-backend/commit/b5b01b)
- Update codecov.yml to add ignored files [`d54309`](https://github.com/reearth/reearth-backend/commit/d54309)
- Ignore generated files in codecov [`9d3822`](https://github.com/reearth/reearth-backend/commit/9d3822)
- Upgrade dependencies [`215947`](https://github.com/reearth/reearth-backend/commit/215947)

## 0.1.0 - 2021-11-01

### 🚀 Features

- Support Auth0 audience ([#3](https://github.com/reearth/reearth-backend/pull/3)) [`c3758e`](https://github.com/reearth/reearth-backend/commit/c3758e)
- Basic auth for projects ([#6](https://github.com/reearth/reearth-backend/pull/6)) [`5db065`](https://github.com/reearth/reearth-backend/commit/5db065)
- Google analytics for scene ([#10](https://github.com/reearth/reearth-backend/pull/10)) [`b44249`](https://github.com/reearth/reearth-backend/commit/b44249)
- Create installable plugins ([#1](https://github.com/reearth/reearth-backend/pull/1)) [`5b7a5f`](https://github.com/reearth/reearth-backend/commit/5b7a5f)
- Add thumbnail, author fields on plugin metadata query ([#15](https://github.com/reearth/reearth-backend/pull/15)) [`888fe0`](https://github.com/reearth/reearth-backend/commit/888fe0)
- Published page api ([#11](https://github.com/reearth/reearth-backend/pull/11)) [`aebac3`](https://github.com/reearth/reearth-backend/commit/aebac3)
- Import dataset from google sheets ([#16](https://github.com/reearth/reearth-backend/pull/16)) [`2ef7ef`](https://github.com/reearth/reearth-backend/commit/2ef7ef)
- Add scenePlugin resolver to layers ([#20](https://github.com/reearth/reearth-backend/pull/20)) [`5213f3`](https://github.com/reearth/reearth-backend/commit/5213f3)
- Marker label position [`bb9e4c`](https://github.com/reearth/reearth-backend/commit/bb9e4c)
- Refine dataset import ([#26](https://github.com/reearth/reearth-backend/pull/26)) [`5dd3db`](https://github.com/reearth/reearth-backend/commit/5dd3db)
- Plugin upload and deletion ([#33](https://github.com/reearth/reearth-backend/pull/33)) [`8742db`](https://github.com/reearth/reearth-backend/commit/8742db)
- New primitives, new properties on primitives [`108711`](https://github.com/reearth/reearth-backend/commit/108711)
- Set scene theme ([#35](https://github.com/reearth/reearth-backend/pull/35)) [`2e4f52`](https://github.com/reearth/reearth-backend/commit/2e4f52)
- Widget align system ([#19](https://github.com/reearth/reearth-backend/pull/19)) [`94611f`](https://github.com/reearth/reearth-backend/commit/94611f)
- Tag system ([#67](https://github.com/reearth/reearth-backend/pull/67)) [`163fcf`](https://github.com/reearth/reearth-backend/commit/163fcf)

### 🔧 Bug Fixes

- Add mutex for each memory repo ([#2](https://github.com/reearth/reearth-backend/pull/2)) [`f4c3b0`](https://github.com/reearth/reearth-backend/commit/f4c3b0)
- Auth0 audience in reearth_config.json [`72e3ed`](https://github.com/reearth/reearth-backend/commit/72e3ed)
- Auth0 domain and multiple auds [`835a02`](https://github.com/reearth/reearth-backend/commit/835a02)
- Signing up and deleting user [`f17b9d`](https://github.com/reearth/reearth-backend/commit/f17b9d)
- Deleting user [`e9b8c9`](https://github.com/reearth/reearth-backend/commit/e9b8c9)
- Sign up and update user [`e5ab87`](https://github.com/reearth/reearth-backend/commit/e5ab87)
- Make gql mutation payloads optional [`9b1c4a`](https://github.com/reearth/reearth-backend/commit/9b1c4a)
- Auth0 [`6a27c6`](https://github.com/reearth/reearth-backend/commit/6a27c6)
- Errors are be overwriten by tx [`2d08c5`](https://github.com/reearth/reearth-backend/commit/2d08c5)
- Deleting user [`f531bd`](https://github.com/reearth/reearth-backend/commit/f531bd)
- Always enable dev mode in debug [`0815d3`](https://github.com/reearth/reearth-backend/commit/0815d3)
- User deletion [`a5eeae`](https://github.com/reearth/reearth-backend/commit/a5eeae)
- Invisible layer issue in published project ([#7](https://github.com/reearth/reearth-backend/pull/7)) [`06cd44`](https://github.com/reearth/reearth-backend/commit/06cd44)
- Dataset link merge bug #378 ([#18](https://github.com/reearth/reearth-backend/pull/18)) [`25da0d`](https://github.com/reearth/reearth-backend/commit/25da0d)
- Ogp image for published page ([#17](https://github.com/reearth/reearth-backend/pull/17)) [`dcb4b0`](https://github.com/reearth/reearth-backend/commit/dcb4b0)
- Change default value of marker label position [`a2059e`](https://github.com/reearth/reearth-backend/commit/a2059e)
- Import dataset from google sheet bug ([#23](https://github.com/reearth/reearth-backend/pull/23)) [`077558`](https://github.com/reearth/reearth-backend/commit/077558)
- Public api param [`846957`](https://github.com/reearth/reearth-backend/commit/846957)
- Replace strings.Split() with strings.field() ([#25](https://github.com/reearth/reearth-backend/pull/25)) [`ba7d16`](https://github.com/reearth/reearth-backend/commit/ba7d16)
- Project public image type [`e82b54`](https://github.com/reearth/reearth-backend/commit/e82b54)
- Published API ([#27](https://github.com/reearth/reearth-backend/pull/27)) [`8ad1f8`](https://github.com/reearth/reearth-backend/commit/8ad1f8)
- Plugin manifest parser bugs ([#32](https://github.com/reearth/reearth-backend/pull/32)) [`78ac13`](https://github.com/reearth/reearth-backend/commit/78ac13)
- Dataset layers are not exported correctly ([#36](https://github.com/reearth/reearth-backend/pull/36)) [`0b8c00`](https://github.com/reearth/reearth-backend/commit/0b8c00)
- Hide parent infobox fields when child infobox is not nil ([#37](https://github.com/reearth/reearth-backend/pull/37)) [`d8c8cd`](https://github.com/reearth/reearth-backend/commit/d8c8cd)
- Mongo.PropertySchema.FindByIDs, propertySchemaID.Equal [`be00da`](https://github.com/reearth/reearth-backend/commit/be00da)
- Gql propertySchemaGroup.translatedTitle resolver [`a4770e`](https://github.com/reearth/reearth-backend/commit/a4770e)
- Use PropertySchemaID.Equal [`8a6459`](https://github.com/reearth/reearth-backend/commit/8a6459)
- Use PropertySchemaID.Equal [`1c3cf1`](https://github.com/reearth/reearth-backend/commit/1c3cf1)
- Tweak field names of model primitive [`080ab9`](https://github.com/reearth/reearth-backend/commit/080ab9)
- Layer importing bug ([#41](https://github.com/reearth/reearth-backend/pull/41)) [`02b17f`](https://github.com/reearth/reearth-backend/commit/02b17f)
- Skip nil geometries ([#42](https://github.com/reearth/reearth-backend/pull/42)) [`90c327`](https://github.com/reearth/reearth-backend/commit/90c327)
- Validate widget extended when moved [`a7daf7`](https://github.com/reearth/reearth-backend/commit/a7daf7)
- Widget extended validation [`98db7e`](https://github.com/reearth/reearth-backend/commit/98db7e)
- Nil error in mongodoc plugin [`d236be`](https://github.com/reearth/reearth-backend/commit/d236be)
- Add widget to default location [`eb1db4`](https://github.com/reearth/reearth-backend/commit/eb1db4)
- Invalid extension data from GraphQL, plugin manifest schema improvement, more friendly error from manifest parser ([#56](https://github.com/reearth/reearth-backend/pull/56)) [`92d137`](https://github.com/reearth/reearth-backend/commit/92d137)
- Translated fields in plugin gql [`0a658a`](https://github.com/reearth/reearth-backend/commit/0a658a)
- Fallback widgetLocation [`579b7a`](https://github.com/reearth/reearth-backend/commit/579b7a)

### 📖 Documentation

- Refine readme ([#28](https://github.com/reearth/reearth-backend/pull/28)) [`a9d209`](https://github.com/reearth/reearth-backend/commit/a9d209)
- Add badges to readme [skip ci] [`cc63cd`](https://github.com/reearth/reearth-backend/commit/cc63cd)

### ✨ Refactor

- Remove unused code [`37b2c2`](https://github.com/reearth/reearth-backend/commit/37b2c2)
- Pkg/error ([#31](https://github.com/reearth/reearth-backend/pull/31)) [`a3f8b6`](https://github.com/reearth/reearth-backend/commit/a3f8b6)
- Graphql adapter ([#40](https://github.com/reearth/reearth-backend/pull/40)) [`2a1d4f`](https://github.com/reearth/reearth-backend/commit/2a1d4f)
- Reorganize graphql schema ([#43](https://github.com/reearth/reearth-backend/pull/43)) [`d3360b`](https://github.com/reearth/reearth-backend/commit/d3360b)

### 🧪 Testing

- Pkg/shp ([#5](https://github.com/reearth/reearth-backend/pull/5)) [`72ed8e`](https://github.com/reearth/reearth-backend/commit/72ed8e)
- Pkg/id ([#4](https://github.com/reearth/reearth-backend/pull/4)) [`c31bdb`](https://github.com/reearth/reearth-backend/commit/c31bdb)

### Miscellaneous Tasks

- Enable nightly release workflow [`16c037`](https://github.com/reearth/reearth-backend/commit/16c037)
- Set up workflows [`819639`](https://github.com/reearth/reearth-backend/commit/819639)
- Fix workflows [`c022a4`](https://github.com/reearth/reearth-backend/commit/c022a4)
- Print config [`0125aa`](https://github.com/reearth/reearth-backend/commit/0125aa)
- Load .env instead of .env.local [`487a73`](https://github.com/reearth/reearth-backend/commit/487a73)
- Add godoc workflow [`9629dd`](https://github.com/reearth/reearth-backend/commit/9629dd)
- Fix godoc workflow [`cc45b5`](https://github.com/reearth/reearth-backend/commit/cc45b5)
- Fix godoc workflow [`0db163`](https://github.com/reearth/reearth-backend/commit/0db163)
- Fix godoc workflow [`9b78fc`](https://github.com/reearth/reearth-backend/commit/9b78fc)
- Fix godoc workflow [`f1e5a7`](https://github.com/reearth/reearth-backend/commit/f1e5a7)
- Fix godoc workflow [`f7866c`](https://github.com/reearth/reearth-backend/commit/f7866c)
- Fix godoc workflow [`5bc089`](https://github.com/reearth/reearth-backend/commit/5bc089)
- Fix godoc workflow [`5f808b`](https://github.com/reearth/reearth-backend/commit/5f808b)
- Fix godoc workflow [`9f8e11`](https://github.com/reearth/reearth-backend/commit/9f8e11)
- Fix godoc workflow [`150550`](https://github.com/reearth/reearth-backend/commit/150550)
- Use go:embed ([#24](https://github.com/reearth/reearth-backend/pull/24)) [`f7866e`](https://github.com/reearth/reearth-backend/commit/f7866e)
- Add internal error log [`41c377`](https://github.com/reearth/reearth-backend/commit/41c377)
- Support multiple platform docker image [`3651e2`](https://github.com/reearth/reearth-backend/commit/3651e2)
- Stop using upx as it doesn't work on arm64 [`3b5f93`](https://github.com/reearth/reearth-backend/commit/3b5f93)
- Update golang version and modules ([#51](https://github.com/reearth/reearth-backend/pull/51)) [`33f4c7`](https://github.com/reearth/reearth-backend/commit/33f4c7)
- Updating modules ([#62](https://github.com/reearth/reearth-backend/pull/62)) [`65ae32`](https://github.com/reearth/reearth-backend/commit/65ae32)
- Add github workflows to release [`fbcdef`](https://github.com/reearth/reearth-backend/commit/fbcdef)
- Fix release workflow, fix build comment [skip ci] [`cfc79a`](https://github.com/reearth/reearth-backend/commit/cfc79a)
- Fix renaming file names in release workflow [`96f0b3`](https://github.com/reearth/reearth-backend/commit/96f0b3)
- Fix and refactor release workflow [skip ci] [`d5466b`](https://github.com/reearth/reearth-backend/commit/d5466b)

<!-- generated by git-cliff -->
