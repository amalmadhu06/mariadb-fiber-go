package memory

// package memory

import (
	"context"
	"github.com/amalmadhu06/mariadb-fiber-go/internal/entities"
	m "github.com/amalmadhu06/mariadb-fiber-go/internal/memory/interfaces"
	repo "github.com/amalmadhu06/mariadb-fiber-go/internal/repository/interfaces"
	"sync"
	"time"
)

// OfferMemory represents an in-memory mechanism to store the latest data
type OfferMemory struct {
	offerRepo repo.OfferRepo                     // offerRepo
	cacheMap  map[string][]entities.OfferCompany // cacheMap holds the latest data based on country
	mutex     *sync.Mutex                        // mutex
}

// NewOfferMemory creates a new instance of the OfferMemory
func NewOfferMemory(offerRepo repo.OfferRepo) m.OfferMemory {
	newOfferMem := &OfferMemory{
		offerRepo: offerRepo,
		cacheMap:  make(map[string][]entities.OfferCompany),
		mutex:     &sync.Mutex{},
	}
	newOfferMem.StartCacheUpdater()
	return newOfferMem
}

// GetOffer returns the latest offers data based on country
func (o *OfferMemory) GetOffer(ctx context.Context, country string) ([]entities.OfferCompany, error) {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	offers, found := o.cacheMap[country]
	if found {
		return offers, nil
	}

	// If the data for the country is not in the cache, fetch it from the repository.
	latestData, err := o.offerRepo.GetOffer(ctx, country)
	if err != nil {
		return nil, err
	}

	// Update the cache with the latest data.
	o.cacheMap[country] = latestData

	return latestData, nil
}

// StartCacheUpdater is a worker for retrieving data from the database and
// storing in memory
func (o *OfferMemory) StartCacheUpdater() {
	go func() {
		for {
			// Todo : rather than hard coding here, fetch list of countries from database and create this array
			countries := []string{"US", "FR", "BR", "CA"}
			for _, country := range countries {
				o.updateCacheForCountry(country)
			}

			// Sleep for 10 seconds before the next update.
			time.Sleep(10 * time.Second)
		}
	}()
}

// updateCacheForCountry is a helper which fetches the latest data from db
// and writes to memory
func (o *OfferMemory) updateCacheForCountry(country string) {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	// Fetch the latest data from the repository.
	latestData, err := o.offerRepo.GetOffer(context.Background(), country)
	if err == nil {
		// Update the cache with the latest data.
		o.cacheMap[country] = latestData
	}
	// Todo : handle error
}
