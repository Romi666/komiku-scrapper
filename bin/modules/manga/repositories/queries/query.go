package queries

import "komiku-srapper/bin/pkg/utils"

type MangaQuery interface {
	GetAllComic() utils.Result
	GetComicInfo(endpoint string) utils.Result
	SearchManga(query string) utils.Result
	GetAllGenre() utils.Result
	GetPopularManga() utils.Result
	GetRecommendedManga() utils.Result
	GetNewestManga() utils.Result
}