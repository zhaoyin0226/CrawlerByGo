package persist

import (
	"crawlerByGo/src/CmMovies/Moudle"
	"crawlerByGo/src/Persist"
	"github.com/elastic/go-elasticsearch"
	"log"
)

type ItemSaverService struct {
	Client    *elasticsearch.Client
	IndexName string
}

func (s *ItemSaverService) Save(item Moudle.Movie, result *string) error {
	err := Persist.Save(item, s.Client, s.IndexName)
	log.Printf("item %#v saved", item)
	if err == nil {
		*result = "OK"
	} else {
		log.Printf("item saved error")
	}
	return err
}
