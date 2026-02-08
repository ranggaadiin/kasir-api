package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
)

type ReportService struct {
	repo *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetTodayReport() (*models.ReportResponse, error) {
	revenue, totalTx, bestSeller, err := s.repo.GetTodaySummary()
	if err != nil {
		return nil, err
	}

	return &models.ReportResponse{
		TotalRevenue:   revenue,
		TotalTransaksi: totalTx,
		ProdukTerlaris: bestSeller,
	}, nil
}
