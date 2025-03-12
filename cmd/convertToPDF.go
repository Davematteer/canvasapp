package cmd

import (
	"fmt"
	"os"

	"github.com/jung-kurt/gofpdf"
)

func ConvertToPDF() {
	// Read the text file
	text, err := os.ReadFile("studyPlan.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Create a new PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	// Add the text to the PDF
	pdf.MultiCell(190, 10, string(text), "", "L", false)

	// Save as PDF
	err = pdf.OutputFileAndClose("study_plan.pdf")
	if err != nil {
		fmt.Println("Error saving PDF:", err)
		return
	}

	fmt.Println("PDF generated successfully!")
}
