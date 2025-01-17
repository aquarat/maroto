// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import maroto "github.com/johnfercher/maroto"
import mock "github.com/stretchr/testify/mock"

// text is an autogenerated mock type for the text type
type Text struct {
	mock.Mock
}

// Add provides a mock function with given fields: text, fontFamily, fontStyle, fontSize, marginTop, align, actualCol, qtdCols
func (_m *Text) Add(text string, fontFamily maroto.Family, fontStyle maroto.Style, fontSize float64, marginTop float64, align maroto.Align, actualCol float64, qtdCols float64) {
	_m.Called(text, fontFamily, fontStyle, fontSize, marginTop, align, actualCol, qtdCols)
}
