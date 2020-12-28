package service

import (
	"github.com/mmagr/planets/internal/model"
	"github.com/mmagr/planets/internal/model/conditions"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockLine struct{ mock.Mock }

func (ml *MockLine) Includes(p model.Point) bool {
	args := ml.Called(p)
	return args.Bool(0)
}

type MockLineFactory struct{ mock.Mock }

func (mlf *MockLineFactory) FromPoints(p1, p2 model.Point) Line {
	args := mlf.Called(p1, p2)
	return args.Get(0).(Line)
}

type MockTriangle struct{ mock.Mock }

func (mt *MockTriangle) Perimeter() float64 {
	args := mt.Called()
	return args.Get(0).(float64)
}

func (mt *MockTriangle) Includes(p model.Point) bool {
	args := mt.Called(p)
	return args.Bool(0)
}

func (mt *MockTriangle) Valid() bool {
	args := mt.Called()
	return args.Bool(0)
}

type MockPolygonFactory struct{ mock.Mock }

func (mpf *MockPolygonFactory) FromPoints(points ...model.Point) Polygon {
	args := mpf.Called(points)
	return args.Get(0).(Polygon)
}

func TestClimatempoPerfect(t *testing.T) {

	ml := new(MockLine)
	ml.On("Includes", mock.Anything).Return(false)

	mlf := new(MockLineFactory)
	mlf.On("FromPoints", mock.Anything, mock.Anything).Return(ml)

	mt := new(MockTriangle)
	mt.On("Valid", mock.Anything).Return(false)
	mt.On("Includes", mock.Anything).Return(true)

	mpf := new(MockPolygonFactory)
	mpf.On("FromPoints", mock.Anything, mock.Anything, mock.Anything).Return(mt)

	tested := climatempo{model.Planet{}, model.Planet{}, model.Planet{}, mlf, mpf}
	condition, _ := tested.ConditionsOn(0)
	assert.Equal(t, conditions.Perfect, condition)
}

func TestClimatempoDraught(t *testing.T) {

	ml := new(MockLine)
	ml.On("Includes", mock.Anything).Return(true)
	mlf := new(MockLineFactory)
	mlf.On("FromPoints", mock.Anything, mock.Anything).Return(ml)

	mt := new(MockTriangle)
	mt.On("Valid", mock.Anything).Return(false)
	mt.On("Includes", mock.Anything).Return(false)

	mpf := new(MockPolygonFactory)
	mpf.On("FromPoints", mock.Anything, mock.Anything, mock.Anything).Return(mt)

	tested := climatempo{model.Planet{}, model.Planet{}, model.Planet{}, mlf, mpf}
	condition, _ := tested.ConditionsOn(0)
	assert.Equal(t, conditions.Draught, condition)
}

func TestClimatempoRain(t *testing.T) {

	ml := new(MockLine)
	ml.On("Includes", mock.Anything).Return(false)
	mlf := new(MockLineFactory)
	mlf.On("FromPoints", mock.Anything, mock.Anything).Return(ml)

	mt := new(MockTriangle)
	mt.On("Valid", mock.Anything).Return(true)
	mt.On("Includes", mock.Anything).Return(true)
	mt.On("Perimeter", mock.Anything).Return(12.3)

	mpf := new(MockPolygonFactory)
	mpf.On("FromPoints", mock.Anything, mock.Anything, mock.Anything).Return(mt)

	tested := climatempo{model.Planet{}, model.Planet{}, model.Planet{}, mlf, mpf}
	condition, metric := tested.ConditionsOn(0)
	assert.Equal(t, conditions.Rain, condition)
	assert.Equal(t, 12.3, metric)
}

func TestClimatempoCloudy(t *testing.T) {

	ml := new(MockLine)
	ml.On("Includes", mock.Anything).Return(false)
	mlf := new(MockLineFactory)
	mlf.On("FromPoints", mock.Anything, mock.Anything).Return(ml)

	mt := new(MockTriangle)
	mt.On("Valid", mock.Anything).Return(true)
	mt.On("Includes", mock.Anything).Return(false)
	mt.On("Perimeter", mock.Anything).Return(32.1)

	mpf := new(MockPolygonFactory)
	mpf.On("FromPoints", mock.Anything, mock.Anything, mock.Anything).Return(mt)

	tested := climatempo{model.Planet{}, model.Planet{}, model.Planet{}, mlf, mpf}
	condition, metric := tested.ConditionsOn(0)
	assert.Equal(t, conditions.Cloudy, condition)
	assert.Equal(t, 32.1, metric)
}
