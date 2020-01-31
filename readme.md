<p align=center>
	<img src="https://raw.githubusercontent.com/rl404/go-malscraper/master/assets/logo.png"><br>
	<a href="https://travis-ci.org/rl404/go-malscraper"><img src="https://api.travis-ci.org/rl404/go-malscraper.svg?branch=master" alt="Build Status"></a>
	<a href="https://coveralls.io/github/rl404/go-malscraper"><img src="https://coveralls.io/repos/github/rl404/go-malscraper/badge.svg" alt="Coverage Status"></a>
	<a href="https://goreportcard.com/report/github.com/rl404/go-malscraper"><img src="https://goreportcard.com/badge/github.com/rl404/go-malscraper" alt="Go Report Card"></a>
	<a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/github/license/rl404/go-malscraper.svg" alt="License: MIT"></a>
	<a href="https://github.com/rl404/go-malscraper/wiki"><img src="https://img.shields.io/badge/docs-wiki-blue" alt="Documentation"></a>
</p>

_go-malscraper_ is just another project scraping/parsing [MyAnimeList](https://myanimelist.net/) website to a useful and easy-to-use data by using [Go](https://golang.org/). It is a simple REST API that you can host yourself. It also provides the API library that you can use for your other Go projects.

Well, it is created to help people get MyAnimeList data without relying on MyAnimeList since they limited/disabled/closed their API. It's working as long as the web is up and the UI design stays the same so it can get the page sources and parse them.

_go-malscraper_ is using [PuerkitoBio's](https://github.com/PuerkitoBio/goquery) HTML DOM parser and inspired by [Jikan's](https://github.com/jikan-me/jikan) API library and my PHP [Mal-Scraper](https://github.com/rl404/MAL-Scraper) library.

## Features
* Get anime information (details, characters, episodes, pictures, etc)
* Get manga information (details, characters, pictures, recommendations, etc)
* Get character information (details, pictures)
* Get people information (details, pictures)
* Get list of all anime/manga's genres
* Get list of all anime/manga's producers/studios/licensors/magazines/serializations
* Get anime/manga's recommendations
* Get anime/manga's reviews
* Search anime, manga, character and people
* Get seasonal anime list
* Get anime, manga, character and people top list
* Get user information (profile, friends, histories, recommendations, reviews)

_More will be coming soon..._

## Quick Start
##### With [docker](https://www.docker.com/) + [docker-compose](https://docs.docker.com/compose/)
```
git clone https://github.com/rl404/go-malscraper.git
cd go-malscraper
docker-compose -f deployment/docker-compose.yml up
```
##### With [Go](https://golang.org/)
```
go get github.com/rl404/go-malscraper
cd $GOPATH/src/github/rl404/go-malscraper/cmd/malscraper
go run main.go
```
[http://localhost:8005](http://localhost:8005) is ready to use.

<table>
	<tr>
		<th>Feature</th>
		<th>Endpoint</th>
		<th>Example</th>
	</tr>
	<tr>
		<td>Get anime detail information</td>
		<td>/v1/anime/:id</td>
		<td>/v1/anime/1</td>
	</tr>
    <tr>
		<td>Get manga detail information</td>
		<td>/v1/manga/:id</td>
		<td>/v1/manga/1</td>
	</tr>
    <tr>
		<td>Get character detail information</td>
		<td>/v1/character/:id</td>
		<td>/v1/character/1</td>
	</tr>
    <tr>
		<td>Get manga detail information</td>
		<td>/v1/manga/:id</td>
		<td>/v1/manga/1</td>
	</tr>
    <tr>
		<td>Get people detail information</td>
		<td>/v1/people/:id</td>
		<td>/v1/people/1</td>
	</tr>
    <tr>
		<td>Get genres detail information</td>
		<td>/v1/genres/:type</td>
		<td>/v1/genres/anime</td>
	</tr>
    <tr>
		<td>Get list of all anime producers</td>
		<td>/v1/producers</td>
		<td>/v1/producers</td>
	</tr>
    <tr>
		<td>Get list of all anime recommendations</td>
		<td>/v1/recommendations/:type</td>
		<td>/v1/recommendations/anime</td>
	</tr>
    <tr>
		<td>Get list of all anime reviews</td>
		<td>/v1/reviews/:type</td>
		<td>/v1/reviews/anime</td>
	</tr>
    <tr>
		<td>Search anime</td>
		<td>/v1/search/anime</td>
		<td>/v1/search/anime?query=naruto</td>
	</tr>
    <tr>
		<td>Get list of seasonal anime</td>
		<td>/v1/season</td>
		<td>/v1/season?year=2019&season=fall</td>
	</tr>
    <tr>
		<td>Get list of top anime</td>
		<td>/v1/top/anime</td>
		<td>/v1/top/anime</td>
	</tr>
    <tr>
		<td>Get user profile</td>
		<td>/v1/user/:user</td>
		<td>/v1/user/:rl404</td>
	</tr>
</table>

*For more detail information, please go to the [wiki](https://github.com/rl404/go-malscraper/wiki).*

## Contributing
1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request.

## Disclamer
_go-malscraper_ is meant for educational purpose and personal usage only. Although there is no limit in using the API, do remember that every scraper method is accessing MyAnimeList page so use it responsibly according to MyAnimeList's [Terms Of Service](https://myanimelist.net/about/terms_of_use).

All data (including anime, manga, people, etc) and MyAnimeList logos belong to their respective copyrights owners. go-malscraper does not have any affiliation with content providers.

## Additional Badges
Since I like badges, I try to collect as many badges as I can. :)

<a href="https://circleci.com/gh/rl404/go-malscraper"><img src="https://circleci.com/gh/rl404/go-malscraper.svg?style=svg" alt="Circle CI"></a>
<a href="https://codecov.io/gh/rl404/go-malscraper"><img src="https://codecov.io/gh/rl404/go-malscraper/branch/master/graph/badge.svg" alt="codecov"></a>
<a href="https://golangci.com/r/github.com/rl404/go-malscraper"><img src="https://golangci.com/badges/github.com/rl404/go-malscraper.svg" alt="GolangCI"></a>
<a href="https://codeclimate.com/github/rl404/go-malscraper/maintainability"><img src="https://api.codeclimate.com/v1/badges/ceb1dee23a8c08aacb5a/maintainability" /></a>
<a href="https://codeclimate.com/github/rl404/go-malscraper/test_coverage"><img src="https://api.codeclimate.com/v1/badges/ceb1dee23a8c08aacb5a/test_coverage" /></a>
<a href="https://www.codacy.com/manual/rl404/go-malscraper?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=rl404/go-malscraper&amp;utm_campaign=Badge_Grade"><img src="https://api.codacy.com/project/badge/Grade/ccef4351408c40e38af9a0f47d6d9195"/></a>

<a href="https://github.com/rl404"><img src="https://forthebadge.com/images/badges/built-with-love.svg" alt="ForTheBadge built-with-love"></a>

## License
MIT License

Copyright (c) 2020 Axel