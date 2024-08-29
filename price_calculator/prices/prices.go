package prices

import (
	"errors"
	"fmt"

	"example.com/price_calculator/conversion"
	"example.com/price_calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	InputPrices       []float64         `json:"input_prices"`
	TaxRate           float64           `json:"tax_rate"`
	TaxIncludedPrices map[string]string `json:"tax_included_prices"`
	// json:"-" : means ignore from output
	IOManager iomanager.IOManager `json:"-"`
}

func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()
	prices, err := conversion.StringsToFloats(lines) // make([]float64, len(lines))
	if err != nil {

		return errors.New("Failed to convert string to float")
	}
	job.InputPrices = prices
	return nil
}
func (job *TaxIncludedPriceJob) Process() error {
	err := job.LoadData()
	if err != nil {
		return err
	}
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("tax-%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	job.TaxIncludedPrices = result
	return job.IOManager.WriteResult(job)

}
func NewTaxIncludedPriceJob(iom iomanager.IOManager, prices []float64, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: prices,
		TaxRate:     taxRate,
		IOManager:   iom,
		// TaxIncludedPrices: make(map[string]float64),
	}
}
