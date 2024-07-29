package persistent

import (
	"errors"
	"github.com/stretchr/testify/require"
	"gopl.io/hw/mocks/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockStore(ctrl)

	// возвращаемая ошибка
	errEmptyKey := errors.New("Указан пустой ключ")

	m.EXPECT().Get("").Return(nil, errEmptyKey)

	_, err := Lookup(m, "")
	require.Error(t, errEmptyKey, err)
}
