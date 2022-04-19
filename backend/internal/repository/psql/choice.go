package psql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"bimbo/internal/model"
)

type ChoiceRepository struct {
	db *sql.DB
}

func ChoiceRepoInit(db *sql.DB) *ChoiceRepository {
	return &ChoiceRepository{
		db: db,
	}
}

func BulkInsert(rows []model.Choice, id int) (string, []interface{}, error) {
	valueStrings := make([]string, 0, len(rows))
	valueArgs := make([]interface{}, 0, len(rows)*5)
	i := 0

	for _, choice := range rows {
		// insert header, etc

		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)", i*5+1, i*5+2, i*5+3, i*5+4, i*5+5))
		valueArgs = append(valueArgs, choice.Key)
		valueArgs = append(valueArgs, choice.Value)
		valueArgs = append(valueArgs, choice.Tag)
		valueArgs = append(valueArgs, time.Now())
		valueArgs = append(valueArgs, id)
		i++
	}
	log.Println(valueArgs, valueStrings)
	sqlQuery := fmt.Sprintf("INSERT INTO bimbo_choice(jsonKey, value, htmlTag, created_at, template_id) VALUES %s", strings.Join(valueStrings, ","))

	return sqlQuery, valueArgs, nil
}

func (ur ChoiceRepository) Create(ctx context.Context, listChoices []model.Choice, tmplId int) error {
	query, args, err := BulkInsert(listChoices, tmplId)
	if err != nil {
		return err
	}
	_, err = ur.db.ExecContext(ctx, query, args...)

	if err != nil {
		return err
	}
	return nil
}

func (cr ChoiceRepository) GetList(ctx context.Context, id int) ([]model.Choice, error) {
	query := `SELECT id, jsonKey, value, htmlTag  FROM bimbo_choice WHERE template_id = $1`
	result := []model.Choice{}
	rows, err := cr.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		temp := model.Choice{}
		if err = rows.Scan(
			&temp.ID,
			&temp.Key,
			&temp.Value,
			&temp.Tag,
		); err != nil {
			return nil, err
		}
		result = append(result, temp)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return result, nil
}
