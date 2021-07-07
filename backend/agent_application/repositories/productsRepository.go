package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/david-drvar/xws2021-nistagram/agent_application/model/domain"
	"github.com/david-drvar/xws2021-nistagram/agent_application/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/agent_application/util"
	"github.com/david-drvar/xws2021-nistagram/agent_application/util/images"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type ProductRepository interface {
	CreateProduct(context.Context, persistence.Product) error
	SaveProductPhoto(ctx context.Context, product persistence.Product) error
	GetAllProductsByAgentId(ctx context.Context, id string) ([]persistence.Product, error)
	GetAllProducts(ctx context.Context) ([]persistence.Product, error)
	GetProductById(ctx context.Context, id string) (persistence.Product, error)
	GetProductByAgentId(ctx context.Context, id string) ([]persistence.Product, error)
	DeleteProduct(ctx context.Context, id string) error
	UpdateProduct(ctx context.Context, product *persistence.Product) error
	OrderProduct(ctx context.Context, order *persistence.Order) error
	GetOrdersByUser(ctx context.Context, userId string) ([]domain.Order, error)
	GetOrdersByAgent(ctx context.Context, agentId string) ([]domain.Order, error)
	GetOrdersByProductId(ctx context.Context, productId string) ([]persistence.Order, error)
}

type productRepository struct {
	DB             *gorm.DB
	userRepository UserRepository
}

func NewProductRepo(db *gorm.DB) (*productRepository, error) {
	if db == nil {
		panic("UserRepository not created, gorm.DB is nil")
	}
	userRepository, _ := NewUserRepo(db)

	return &productRepository{DB: db, userRepository: userRepository}, nil
}

func (repository *productRepository) CreateProduct(ctx context.Context, product persistence.Product) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "CreateUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	product.Id = uuid.New().String()
	resultUser := repository.DB.Create(&product)
	if resultUser.Error != nil {
		return resultUser.Error
	}

	if product.Photo != "" {
		err := repository.SaveProductPhoto(ctx, product)
		if err != nil {
			return err
		}
	}

	return nil
}

func (repository *productRepository) GetAllProductsByAgentId(ctx context.Context, id string) ([]persistence.Product, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllProductsByAgentId")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var products []persistence.Product
	resultUser := repository.DB.Where("agent_id = ? AND is_active = true", id).Find(&products)
	if resultUser.Error != nil {
		return nil, resultUser.Error
	}

	return products, nil
}

func (repository *productRepository) GetProductById(ctx context.Context, id string) (persistence.Product, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetProductById")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var product persistence.Product
	resultUser := repository.DB.Where("id = ? AND is_active = true", id).Find(&product)
	if resultUser.Error != nil {
		return persistence.Product{}, resultUser.Error
	}

	return product, nil
}

func (repository *productRepository) GetProductByAgentId(ctx context.Context, id string) ([]persistence.Product, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetProductById")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var products []persistence.Product
	resultUser := repository.DB.Where("agent_id = ? AND is_active = true", id).Find(&products)
	if resultUser.Error != nil {
		return nil, resultUser.Error
	}

	return products, nil
}

func (repository *productRepository) DeleteProduct(ctx context.Context, id string) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "DeleteProduct")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var product persistence.Product
	product, err := repository.GetProductById(ctx, id)
	if err != nil {
		return err
	}
	product.IsActive = false
	return repository.DB.Model(&product).Update("is_active", false).Error
}

func (repository *productRepository) UpdateProduct(ctx context.Context, product *persistence.Product) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateProduct")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	productDB, err := repository.GetProductById(ctx, product.Id)
	if err != nil {
		return err
	}
	productDB.Name = product.Name
	productDB.Quantity = product.Quantity
	productDB.Price = product.Price

	err = repository.DB.Model(&productDB).Updates(persistence.Product{Price: productDB.Price, Name: productDB.Name, Quantity: productDB.Quantity}).Error
	if err != nil {
		return err
	}

	if product.Photo != "" {
		productDB.Photo = product.Photo
		err := repository.SaveProductPhoto(ctx, productDB)
		if err != nil {
			return err
		}
	}
	return nil
}

func (repository *productRepository) GetAllProducts(ctx context.Context) ([]persistence.Product, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAllProducts")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var products []persistence.Product
	resultUser := repository.DB.Where("is_active = true").Find(&products)
	if resultUser.Error != nil {
		return nil, resultUser.Error
	}

	return products, nil
}

func (repository *productRepository) SaveProductPhoto(ctx context.Context, product persistence.Product) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "SaveUserProfilePhoto")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	mimeType, err := images.GetImageType(product.Photo)
	if err != nil {
		return err
	}

	t := time.Now()
	formatted := fmt.Sprintf("%s%d%02d%02d%02d%02d%02d%02d", product.Id, t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond())
	name := formatted + "." + mimeType

	err = images.SaveImage(name, product.Photo)
	if err != nil {
		return err
	}

	product.Photo = util.GetContentLocation(name)
	db := repository.DB.Model(&product).Where("id = ?", product.Id).Updates(product)
	if db.Error != nil {
		return db.Error
	}

	return nil
}

func (repository *productRepository) OrderProduct(ctx context.Context, order *persistence.Order) error {
	span := tracer.StartSpanFromContextMetadata(ctx, "OrderProduct")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var product persistence.Product
	product, err := repository.GetProductById(ctx, order.ProductId)
	if err != nil {
		return err
	}
	if order.Quantity > product.Quantity {
		return errors.New("Order quantity cannot exceed product quantity")
	}

	order.Id = uuid.New().String()
	order.DateCreated = time.Now()
	order.TotalPrice = float32(order.Quantity) * product.Price

	return repository.DB.Create(&order).Error
}

func (repository *productRepository) GetOrdersByUser(ctx context.Context, userId string) ([]domain.Order, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetOrdersByUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var orders []persistence.Order
	resultUser := repository.DB.Where("user_id = ?", userId).Find(&orders)
	if resultUser.Error != nil {
		return nil, resultUser.Error
	}

	var finalOrders []domain.Order
	for _, order := range orders {
		product, err := repository.GetProductById(ctx, order.ProductId)
		if err != nil {
			return nil, err
		}
		user, err := repository.userRepository.GetUserById(ctx, product.AgentId)
		if err != nil {
			return nil, err
		}
		finalOrders = append(finalOrders, domain.Order{
			Id:          order.Id,
			UserId:      order.UserId,
			ProductId:   order.ProductId,
			Quantity:    order.Quantity,
			DateCreated: order.DateCreated,
			TotalPrice:  order.TotalPrice,
			Referral:    order.Referral,
			Username:    user.Username,
			ProductName: product.Name,
		})

	}

	return finalOrders, nil
}

func (repository *productRepository) GetOrdersByProductId(ctx context.Context, productId string) ([]persistence.Order, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetOrdersByUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	var orders []persistence.Order
	resultUser := repository.DB.Where("product_id = ?", productId).Find(&orders)
	if resultUser.Error != nil {
		return nil, resultUser.Error
	}

	return orders, nil
}

func (repository *productRepository) GetOrdersByAgent(ctx context.Context, agentId string) ([]domain.Order, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetOrdersByAgent")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	products, err := repository.GetProductByAgentId(ctx, agentId)
	if err != nil {
		return nil, err
	}

	var orders []domain.Order
	for _, product := range products {
		tempOrders, err := repository.GetOrdersByProductId(ctx, product.Id)
		if err != nil {
			return nil, err
		}
		for _, order := range tempOrders {
			orders = append(orders, domain.Order{
				Id:          order.Id,
				UserId:      order.UserId,
				ProductId:   order.ProductId,
				Quantity:    order.Quantity,
				DateCreated: order.DateCreated,
				TotalPrice:  order.TotalPrice,
				Referral:    order.Referral,
				Username:    "",
				ProductName: product.Name,
			})
		}
	}

	var finalOrders []domain.Order
	for _, order := range orders {
		user, err := repository.userRepository.GetUserById(ctx, order.UserId)
		if err != nil {
			return nil, err
		}
		finalOrders = append(finalOrders, domain.Order{
			Id:          order.Id,
			UserId:      order.UserId,
			ProductId:   order.ProductId,
			Quantity:    order.Quantity,
			DateCreated: order.DateCreated,
			TotalPrice:  order.TotalPrice,
			Referral:    order.Referral,
			Username:    user.Username,
			ProductName: order.ProductName,
		})
	}

	return finalOrders, nil
}
