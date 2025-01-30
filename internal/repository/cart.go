package repository

import (
    "context"
    "database/sql"
    "time"

    "github.com/ingarondel/GO-APIDevelopment/internal/model"
)

type CartRepository struct {
    db *sql.DB
}

func NewCartRepository(db *sql.DB) *CartRepository {
    return &CartRepository{db}
}
// TODO context.WithTimeout(ctx, 5*time.Second) - это делается на уровне handler тк именно там ты задаешь время на выполнение запроса в целом
// здесь лишь вызовы к бд
// TODO тебе нужно возвращать (model.Cart , error)
func (r *CartRepository) CreateCart(ctx context.Context, cart *model.Cart) error {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()
    query := "INSERT INTO carts DEFAULT VALUES RETURNING id"
    // TODO нужно обязатель оборачивать ошибки 
    // if err!=nil{
    //     ...
    // }

    return r.db.QueryRowContext(ctx, query).Scan(&cart.ID)
}
// TODO context.WithTimeout(ctx, 5*time.Second) - это делается на уровне handler тк именно там ты задаешь время на выполнение запроса в целом
// здесь лишь вызовы к бд
func (r *CartRepository) GetCart(ctx context.Context, id int64) (model.Cart, error) {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()
    query := "SELECT id FROM carts WHERE id = $1"
    var cart model.Cart
        // TODO а что если такой cart не существует в базе?
        // в этом случае тебе нужно вернуть ошибку о том, что cart не существует 
        // рекомендую создать свой тип ошибки

    // TODO нужно обязатель оборачивать ошибки 
    // if err!=nil{
    //     ...
    // }
    err := r.db.QueryRowContext(ctx, query, id).Scan(&cart.ID)
    return cart, err
}

//во все методы репозитория нужно передавать контекстс
//внутри каждого метода нужно создавать новый контекст с таймаутом допустим на 5 сек
//использовать sqldb
