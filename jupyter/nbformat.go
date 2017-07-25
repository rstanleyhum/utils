package jupyter

import "encoding/json"

// CellTypeCode is the value of the CellType that denotes a code block
const CellTypeCode = "code"

// Notebook is the top-level data based on Jupyter nbformat version 4
type Notebook struct {
	TopLevelData
	Cells []Cell `json:"cells"`
}

// TopLevelData contains all the top level fields except the Cell list
type TopLevelData struct {
	Metadata    interface{} `json:"metadata"`
	FormatMajor interface{} `json:"nbformat"`
	FormatMinor interface{} `json:"nbformat_minor"`
}

// Cell is the combined cell format
type Cell struct {
	CommonCell
	OutputCell
}

// CommonCell contains all of the common fields for a cell except for outputs
type CommonCell struct {
	CellType    interface{} `json:"cell_type"`
	Metadata    interface{} `json:"metadata"`
	Source      interface{} `json:"source"`
	Attachments interface{} `json:"attachments,omitempty"`
}

// OutputCell contains all of the output fields for a cell
type OutputCell struct {
	ExecutionCount interface{} `json:"execution_count,omitempty"`
	Outputs        interface{} `json:"outputs,omitempty"`
}

// SourceOnlyNotebook is the notebook which reads only the SourceOnlyCells
type SourceOnlyNotebook struct {
	TopLevelData
	Cells []SourceOnlyCell `json:"cells"`
}

// SourceOnlyCell is the struct for the source only and does not have the output fields
type SourceOnlyCell struct {
	CommonCell
}

// MarshalJSON on SourceOnlyCell adds ExecutionCount and Outputs to code cells before encoding to JSON object
func (c *SourceOnlyCell) MarshalJSON() ([]byte, error) {
	cell := Cell{
		CommonCell{
			CellType:    c.CellType,
			Metadata:    c.Metadata,
			Source:      c.Source,
			Attachments: c.Attachments,
		},
		OutputCell{
			ExecutionCount: nil,
			Outputs:        nil,
		},
	}
	if c.CellType == CellTypeCode {
		cell.ExecutionCount = 0
		cell.Outputs = make([]interface{}, 0)
	}
	return json.Marshal(cell)
}
