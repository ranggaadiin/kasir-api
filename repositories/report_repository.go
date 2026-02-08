package repositories

import (
	"database/sql"
	"kasir-api/models"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (r *ReportRepository) GetTodaySummary() (int, int, *models.BestSeller, error) {
	var revenue int
	var totalTx int

	err := r.db.QueryRow(`
		SELECT COALESCE(SUM(total_amount), 0), COUNT(*)
		FROM transactions
		WHERE DATE(created_at) = CURRENT_DATE
	`).Scan(&revenue, &totalTx)
	if err != nil {
		return 0, 0, nil, err
	}

	var bs models.BestSeller
	err = r.db.QueryRow(`
		SELECT p.name, SUM(td.quantity)
		FROM transaction_details td
		JOIN products p ON p.id = td.product_id
		JOIN transactions t ON t.id = td.transaction_id
		WHERE DATE(t.created_at) = CURRENT_DATE
		GROUP BY p.name
		ORDER BY SUM(td.quantity) DESC
		LIMIT 1
	`).Scan(&bs.Nama, &bs.QtyTerjual)

	if err == sql.ErrNoRows {
		return revenue, totalTx, nil, nil
	}
	if err != nil {
		return 0, 0, nil, err
	}

	return revenue, totalTx, &bs, nil
}
