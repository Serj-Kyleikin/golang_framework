package subscriptions

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	coredb "subscriptions/Backend/core/db"
	"subscriptions/Backend/db/repositories"

	subscriptionsContracts "subscriptions/Backend/services/subscriptions/contracts"
)

type SubscriptionsRepository struct {
	*repositories.BaseRepository[SubscriptionModel]
	pool *pgxpool.Pool
}

func Construct() *SubscriptionsRepository {
	pool := coredb.MustPool()
	return &SubscriptionsRepository{
		BaseRepository: repositories.Construct[SubscriptionModel](pool, "subscriptions"),
		pool:           pool,
	}
}

func (subscriptionsRepository *SubscriptionsRepository) SumTotalCost(c *gin.Context, totalInputDTO subscriptionsContracts.TotalInputDTO) (int64, error) {

	start := totalInputDTO.StartDate
	end := totalInputDTO.EndDate
	userID := totalInputDTO.UserID
	serviceName := totalInputDTO.ServiceName

	var stringsBuilder strings.Builder
	stringsBuilder.WriteString(`SELECT COALESCE(SUM(price), 0) FROM subscriptions WHERE start_date >= $1 AND start_date <= $2`)

	args := make([]any, 0, 4)
	args = append(args, start, end)
	argPos := 3

	if userID != nil {
		stringsBuilder.WriteString(fmt.Sprintf(` AND user_id = $%d`, argPos))
		args = append(args, *userID)
		argPos++
	}

	if serviceName != nil && *serviceName != "" {
		stringsBuilder.WriteString(fmt.Sprintf(` AND service_name = $%d`, argPos))
		args = append(args, *serviceName)
		argPos++
	}

	var total int64
	if err := subscriptionsRepository.pool.QueryRow(c, stringsBuilder.String(), args...).Scan(&total); err != nil {
		return 0, err
	}

	return total, nil
}
