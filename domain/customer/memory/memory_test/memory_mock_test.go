package memory_test

import (
	"ddd_tavern/aggregate"
	"ddd_tavern/domain/customer"
	repoMock "ddd_tavern/domain/customer/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"testing"
)

func TestGet(t *testing.T) {

	//+ Имитируем репозиторий
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := repoMock.NewMockCustomerRepository(ctl)

	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}
	//- Имитируем репозиторий

	//+ Создаем нового покупателя
	cust, err := aggregate.NewCustomer("Percy")
	if err != nil {
		t.Fatal(err)
	}
	id := cust.GetID()
	//- Создаем нового покупателя

	//+ Тестируем "взятие" из репозитория
	testCases := []testCase{
		{
			name:        "No Customer By ID",
			id:          uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			expectedErr: customer.ErrCustomerNotFound,
		}, {
			name:        "Customer By ID",
			id:          id,
			expectedErr: nil,
		},
	}

	//+ Закладываем работу репозитория (что ожидать?)
	repo.EXPECT().Get(uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479")).Return(aggregate.Customer{}, customer.ErrCustomerNotFound)
	repo.EXPECT().Get(id).Return(cust, nil)
	//- Закладываем работу репозитория (что ожидать?)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			_, err := repo.Get(tc.id)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
	//- Тестируем "взятие" из репозитория

}
