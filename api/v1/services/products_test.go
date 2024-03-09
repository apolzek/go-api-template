package services

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetProducts(t *testing.T) {
	// Cria um novo mock de banco de dados e um contexto falso
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	ctx := context.Background()

	// Define o comportamento esperado do mock
	rows := sqlmock.NewRows([]string{"id", "name", "description", "price", "created", "modified"}).
		AddRow(1, "Product 1", "Description 1", 100, time.Now(), time.Now()).
		AddRow(2, "Product 2", "Description 2", 200, time.Now(), time.Now())
	// mock.ExpectQuery(`SELECT \* FROM "products"`).WillReturnRows(rows)
	mock.ExpectQuery(`SELECT "products"\.\* FROM "products"`).WillReturnRows(rows)

	// Chama a função GetProducts
	products, serviceErr := GetProducts(db, ctx)

	// Verifica se o mock foi chamado conforme esperado
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// Verifica se o serviço retornou o resultado correto
	assert.Nil(t, serviceErr)
	assert.NotNil(t, products)
	assert.Len(t, products, 2)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 2", products[1].Name)
}
