package service

import (
	"encoding/json"
	"fmt"
)

type IProcessorService interface {
	ProcessData(jsonData []byte)
}

type ProcessorService[T any] struct {
	JSONService IJSONService[T]
}

func NewProcessorService[T any](jsonService IJSONService[T]) *ProcessorService[T] {
	return &ProcessorService[T]{
		JSONService: jsonService,
	}
}

func (p *ProcessorService[T]) ProcessData(jsonData []byte) error {
	data, err := p.JSONService.ValidateJSON(jsonData)
	if err != nil {
		return err
	}

	jsonOutput, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}
	fmt.Println(string(jsonOutput))

	return nil
}
